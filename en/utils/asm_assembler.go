// x86 汇编代码汇编器模块
// 提供机器码构建、代码执行、远程注入等功能。
// 内部使用指令集模块生成的机器码，支持在当前进程或远程进程中执行。
package utils

import (
	"encoding/binary"
	"sync"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	asmCodeBuf   []byte
	asmCodeMu    sync.Mutex
)

func ASM_SetCode(code []byte) {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	if code == nil || len(code) == 0 {
		asmCodeBuf = nil
	} else {
		asmCodeBuf = make([]byte, len(code))
		copy(asmCodeBuf, code)
	}
}

func ASM_GetCode() []byte {
	asmCodeMu.Lock()
	defer asmCodeMu.Unlock()
	result := make([]byte, len(asmCodeBuf))
	copy(result, asmCodeBuf)
	return result
}

func asmAppendBytes(b ...byte) {
	asmCodeBuf = append(asmCodeBuf, b...)
}

func asmAppendHex(hex string) {
	bytes := asmParseHex(hex)
	asmCodeBuf = append(asmCodeBuf, bytes...)
}

func asmAppendInt32(v int32) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(v))
	asmAppendBytes(buf...)
}

func asmAppendUint32(v uint32) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, v)
	asmAppendBytes(buf...)
}

func asmAppendInt8(v int8) {
	asmAppendBytes(byte(v))
}

func asmParseHex(hex string) []byte {
	result := make([]byte, 0, len(hex)/2)
	for i := 0; i < len(hex); i += 2 {
		b := byte(0)
		for j := 0; j < 2; j++ {
			c := hex[i+j]
			b <<= 4
			if c >= '0' && c <= '9' {
				b |= c - '0'
			} else if c >= 'A' && c <= 'F' {
				b |= c - 'A' + 10
			} else if c >= 'a' && c <= 'f' {
				b |= c - 'a' + 10
			}
		}
		result = append(result, b)
	}
	return result
}

// ASM_Execute 在当前进程中执行已构建的机器码。
// 使用 VirtualAlloc 分配可执行内存，执行后自动释放。
// 返回执行结果（EAX 寄存器值）。
func ASM_Execute() (uintptr, error) {
	asmCodeMu.Lock()
	code := make([]byte, len(asmCodeBuf))
	copy(code, asmCodeBuf)
	asmCodeMu.Unlock()

	if len(code) == 0 {
		return 0, nil
	}

	addr, err := windows.VirtualAlloc(0, uintptr(len(code)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_EXECUTE_READWRITE)
	if err != nil {
		return 0, err
	}

	dst := unsafe.Slice((*byte)(unsafe.Pointer(addr)), len(code))
	copy(dst, code)

	fn := *(*func() uintptr)(unsafe.Pointer(&addr))
	result := fn()

	windows.VirtualFree(addr, 0, windows.MEM_RELEASE)
	return result, nil
}

// ASM_ExecuteRemote 在指定进程中远程执行已构建的机器码。
// 使用 VirtualAllocEx + WriteProcessMemory + CreateRemoteThread。
func ASM_ExecuteRemote(processID uint32, code []byte) error {
	asmCodeMu.Lock()
	if code == nil {
		code = make([]byte, len(asmCodeBuf))
		copy(code, asmCodeBuf)
	}
	asmCodeMu.Unlock()

	if len(code) == 0 {
		return nil
	}

	hProcess, err := windows.OpenProcess(
		windows.PROCESS_CREATE_THREAD|windows.PROCESS_QUERY_INFORMATION|
			windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_WRITE|windows.PROCESS_VM_READ,
		false,
		processID,
	)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(hProcess)

	remoteAddr, err := windows.VirtualAllocEx(hProcess, 0, uintptr(len(code)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_EXECUTE_READWRITE)
	if err != nil {
		return err
	}

	var written uintptr
	err = windows.WriteProcessMemory(hProcess, remoteAddr, &code[0], uintptr(len(code)), &written)
	if err != nil {
		windows.VirtualFreeEx(hProcess, remoteAddr, 0, windows.MEM_RELEASE)
		return err
	}

	var threadID uint32
	hThread, err := windows.CreateRemoteThread(hProcess, nil, 0, remoteAddr, 0, 0, &threadID)
	if err != nil {
		windows.VirtualFreeEx(hProcess, remoteAddr, 0, windows.MEM_RELEASE)
		return err
	}
	defer windows.CloseHandle(hThread)

	windows.WaitForSingleObject(hThread, windows.INFINITE)
	windows.VirtualFreeEx(hProcess, remoteAddr, 0, windows.MEM_RELEASE)
	return nil
}