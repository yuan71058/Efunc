package utils

import (
	"encoding/binary"
	"sync"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	h汇编代码缓冲区   []byte
	h汇编代码缓冲锁 sync.Mutex
)

func H汇编_置代码(代码 []byte) {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	if 代码 == nil || len(代码) == 0 {
		h汇编代码缓冲区 = nil
	} else {
		h汇编代码缓冲区 = make([]byte, len(代码))
		copy(h汇编代码缓冲区, 代码)
	}
}

func H汇编_取代码() []byte {
	h汇编代码缓冲锁.Lock()
	defer h汇编代码缓冲锁.Unlock()
	result := make([]byte, len(h汇编代码缓冲区))
	copy(result, h汇编代码缓冲区)
	return result
}

func h汇编追加字节(b ...byte) {
	h汇编代码缓冲区 = append(h汇编代码缓冲区, b...)
}

func h汇编追加密文(hex string) {
	bytes := h汇编解析十六进制(hex)
	h汇编代码缓冲区 = append(h汇编代码缓冲区, bytes...)
}

func h汇编追加Int32(v int32) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(v))
	h汇编追加字节(buf...)
}

func h汇编追加Uint32(v uint32) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, v)
	h汇编追加字节(buf...)
}

func h汇编追加Int8(v int8) {
	h汇编追加字节(byte(v))
}

func h汇编解析十六进制(hex string) []byte {
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

// H汇编_运行汇编代码 在当前进程中执行已构建的机器码。
// 使用 VirtualAlloc 分配可执行内存，执行后自动释放。
// 返回执行结果（EAX 寄存器值）。
func H汇编_运行汇编代码() (uintptr, error) {
	h汇编代码缓冲锁.Lock()
	代码 := make([]byte, len(h汇编代码缓冲区))
	copy(代码, h汇编代码缓冲区)
	h汇编代码缓冲锁.Unlock()

	if len(代码) == 0 {
		return 0, nil
	}

	addr, err := windows.VirtualAlloc(0, uintptr(len(代码)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_EXECUTE_READWRITE)
	if err != nil {
		return 0, err
	}

	dst := unsafe.Slice((*byte)(unsafe.Pointer(addr)), len(代码))
	copy(dst, 代码)

	fn := *(*func() uintptr)(unsafe.Pointer(&addr))
	result := fn()

	windows.VirtualFree(addr, 0, windows.MEM_RELEASE)
	return result, nil
}

// H汇编_远程执行汇编代码 在指定进程中远程执行已构建的机器码。
// 使用 VirtualAllocEx + WriteProcessMemory + CreateRemoteThread。
func H汇编_远程执行汇编代码(进程ID uint32, 代码 []byte) error {
	h汇编代码缓冲锁.Lock()
	if 代码 == nil {
		代码 = make([]byte, len(h汇编代码缓冲区))
		copy(代码, h汇编代码缓冲区)
	}
	h汇编代码缓冲锁.Unlock()

	if len(代码) == 0 {
		return nil
	}

	hProcess, err := windows.OpenProcess(
		windows.PROCESS_CREATE_THREAD|windows.PROCESS_QUERY_INFORMATION|
			windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_WRITE|windows.PROCESS_VM_READ,
		false,
		进程ID,
	)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(hProcess)

	remoteAddr, err := windows.VirtualAllocEx(hProcess, 0, uintptr(len(代码)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_EXECUTE_READWRITE)
	if err != nil {
		return err
	}

	var written uintptr
	err = windows.WriteProcessMemory(hProcess, remoteAddr, &代码[0], uintptr(len(代码)), &written)
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