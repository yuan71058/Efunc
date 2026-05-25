//go:build windows

// 内存操作模块
// 提供进程内存读写、内存搜索（AOB 特征码搜索）、虚拟内存管理、进程枚举与管理等功能。
// 基于 Windows kernel32.dll API 实现，支持 ReadProcessMemory、WriteProcessMemory、
// VirtualAllocEx、VirtualProtectEx 等底层调用。
// 包含进程快照（Toolhelp32Snapshot）枚举和多级指针读取能力。
package utils

import (
	"errors"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	kernel32Mem = syscall.NewLazyDLL("kernel32.dll")

	procOpenProcessM              = kernel32Mem.NewProc("OpenProcess")
	procCloseHandleM              = kernel32Mem.NewProc("CloseHandle")
	procReadProcessMemory         = kernel32Mem.NewProc("ReadProcessMemory")
	procWriteProcessMemory        = kernel32Mem.NewProc("WriteProcessMemory")
	procVirtualAllocEx            = kernel32Mem.NewProc("VirtualAllocEx")
	procVirtualFreeEx             = kernel32Mem.NewProc("VirtualFreeEx")
	procVirtualProtectEx          = kernel32Mem.NewProc("VirtualProtectEx")
	procVirtualQueryEx            = kernel32Mem.NewProc("VirtualQueryEx")

	procCreateToolhelp32Snapshot = kernel32Mem.NewProc("CreateToolhelp32Snapshot")
	procProcess32FirstW          = kernel32Mem.NewProc("Process32FirstW")
	procProcess32NextW           = kernel32Mem.NewProc("Process32NextW")
	procModule32FirstW           = kernel32Mem.NewProc("Module32FirstW")
	procModule32NextW            = kernel32Mem.NewProc("Module32NextW")
	procGetModuleBaseNameW       = kernel32Mem.NewProc("GetModuleBaseNameW")

	procGetExitCodeProcessM  = kernel32Mem.NewProc("GetExitCodeProcess")
	procTerminateProcessM    = kernel32Mem.NewProc("TerminateProcess")
)

const (
	PROCESS_VM_READ           = 0x0010
	PROCESS_VM_WRITE          = 0x0020
	PROCESS_VM_OPERATION      = 0x0008
	PROCESS_QUERY_INFORMATION = 0x0400
	PROCESS_TERMINATE         = 0x0001
	PROCESS_CREATE_THREAD     = 0x0002
	PROCESS_ALL_ACCESS        = 0x1F0FFF

	PAGE_NOACCESS          = 0x01
	PAGE_READONLY          = 0x02
	PAGE_READWRITE         = 0x04
	PAGE_WRITECOPY         = 0x08
	PAGE_EXECUTE           = 0x10
	PAGE_EXECUTE_READ      = 0x20
	PAGE_EXECUTE_READWRITE = 0x40

	TH32CS_SNAPPROCESS  = 0x00000002
	TH32CS_SNAPMODULE   = 0x00000008
	TH32CS_SNAPMODULE32 = 0x00000010

	MAX_MODULE_NAME32 = 255
	MAX_PATH          = 260
)

type MEMORY_BASIC_INFORMATION struct {
	BaseAddress       uintptr
	AllocationBase    uintptr
	AllocationProtect uint32
	RegionSize        uintptr
	State             uint32
	Protect           uint32
	Type              uint32
}

type PROCESSENTRY32W struct {
	DwSize              uint32
	CntUsage            uint32
	Th32ProcessID       uint32
	Th32DefaultHeapID   uintptr
	Th32ModuleID        uint32
	CntThreads          uint32
	Th32ParentProcessID uint32
	PcPriClassBase      int32
	DwFlags             uint32
	SzExeFile           [MAX_PATH]uint16
}

type MODULEENTRY32W struct {
	DwSize        uint32
	Th32ModuleID  uint32
	Th32ProcessID uint32
	GlblcntUsage  uint32
	ProccntUsage  uint32
	ModBaseAddr   uintptr
	ModBaseSize   uint32
	HModule       syscall.Handle
	SzModule      [MAX_MODULE_NAME32 + 1]uint16
	SzExePath     [MAX_PATH]uint16
}

// Memory_OpenProcess 打开目标进程并获取句柄。
// 参数 processID：目标进程 PID
// 参数 desiredAccess：进程访问权限（PROCESS_VM_READ | PROCESS_VM_WRITE 等）
// 返回 syscall.Handle：进程句柄（使用完毕后需调用 Memory_CloseHandle）
// 返回 error：失败时返回错误信息
func Memory_OpenProcess(processID uint32, desiredAccess uint32) (syscall.Handle, error) {
	hProcess, _, err := procOpenProcessM.Call(uintptr(desiredAccess), 0, uintptr(processID))
	if hProcess == 0 {
		return 0, err
	}
	return syscall.Handle(hProcess), nil
}

// Memory_ReadInt32 从目标进程读取 32 位有符号整数。
// 参数 hProcess：进程句柄
// 参数 address：目标内存地址
// 返回 int32：读取的整数值
// 返回 error：失败时返回错误信息
func Memory_ReadInt32(hProcess syscall.Handle, address uintptr) (int32, error) {
	var buffer [4]byte
	var numberOfBytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess), address,
		uintptr(unsafe.Pointer(&buffer[0])), 4,
		uintptr(unsafe.Pointer(&numberOfBytesRead)),
	)
	if ret == 0 {
		return 0, err
	}
	return int32(buffer[0]) | int32(buffer[1])<<8 | int32(buffer[2])<<16 | int32(buffer[3])<<24, nil
}

// Memory_ReadInt64 从目标进程读取 64 位有符号整数。
// 参数 hProcess：进程句柄
// 参数 address：目标内存地址
// 返回 int64：读取的整数值
// 返回 error：失败时返回错误信息
func Memory_ReadInt64(hProcess syscall.Handle, address uintptr) (int64, error) {
	var buffer [8]byte
	var numberOfBytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess), address,
		uintptr(unsafe.Pointer(&buffer[0])), 8,
		uintptr(unsafe.Pointer(&numberOfBytesRead)),
	)
	if ret == 0 {
		return 0, err
	}
	return int64(buffer[0]) | int64(buffer[1])<<8 | int64(buffer[2])<<16 | int64(buffer[3])<<24 |
		int64(buffer[4])<<32 | int64(buffer[5])<<40 | int64(buffer[6])<<48 | int64(buffer[7])<<56, nil
}

// Memory_ReadBytes 从目标进程读取指定长度的原始字节数据。
// 参数 hProcess：进程句柄
// 参数 address：目标内存地址
// 参数 length：要读取的字节数
// 返回 []byte：读取的字节数据
// 返回 error：失败时返回错误信息
func Memory_ReadBytes(hProcess syscall.Handle, address uintptr, length uint32) ([]byte, error) {
	buffer := make([]byte, length)
	var numberOfBytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess), address,
		uintptr(unsafe.Pointer(&buffer[0])), uintptr(length),
		uintptr(unsafe.Pointer(&numberOfBytesRead)),
	)
	if ret == 0 {
		return nil, err
	}
	return buffer[:numberOfBytesRead], nil
}

// Memory_ReadFloat32 从目标进程读取 32 位浮点数（float32）。
// 参数 hProcess：进程句柄
// 参数 address：目标内存地址
// 返回 float32：读取的浮点数值
// 返回 error：失败时返回错误信息
func Memory_ReadFloat32(hProcess syscall.Handle, address uintptr) (float32, error) {
	var buffer [4]byte
	var numberOfBytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess), address,
		uintptr(unsafe.Pointer(&buffer[0])), 4,
		uintptr(unsafe.Pointer(&numberOfBytesRead)),
	)
	if ret == 0 {
		return 0, err
	}
	bits := uint32(buffer[0]) | uint32(buffer[1])<<8 | uint32(buffer[2])<<16 | uint32(buffer[3])<<24
	return *(*float32)(unsafe.Pointer(&bits)), nil
}

// Memory_ReadFloat64 从目标进程读取 64 位浮点数（float64）。
// 参数 hProcess：进程句柄
// 参数 address：目标内存地址
// 返回 float64：读取的浮点数值
// 返回 error：失败时返回错误信息
func Memory_ReadFloat64(hProcess syscall.Handle, address uintptr) (float64, error) {
	var buffer [8]byte
	var numberOfBytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess), address,
		uintptr(unsafe.Pointer(&buffer[0])), 8,
		uintptr(unsafe.Pointer(&numberOfBytesRead)),
	)
	if ret == 0 {
		return 0, err
	}
	bits := uint64(buffer[0]) | uint64(buffer[1])<<8 | uint64(buffer[2])<<16 | uint64(buffer[3])<<24 |
		uint64(buffer[4])<<32 | uint64(buffer[5])<<40 | uint64(buffer[6])<<48 | uint64(buffer[7])<<56
	return *(*float64)(unsafe.Pointer(&bits)), nil
}

// Memory_WriteInt32 向目标进程写入 32 位有符号整数。
// 参数 hProcess：进程句柄
// 参数 address：目标内存地址
// 参数 value：要写入的整数值
// 返回 error：失败时返回错误信息
func Memory_WriteInt32(hProcess syscall.Handle, address uintptr, value int32) error {
	buffer := [4]byte{
		byte(value),
		byte(value >> 8),
		byte(value >> 16),
		byte(value >> 24),
	}
	var numberOfBytesWritten uintptr
	ret, _, err := procWriteProcessMemory.Call(
		uintptr(hProcess), address,
		uintptr(unsafe.Pointer(&buffer[0])), 4,
		uintptr(unsafe.Pointer(&numberOfBytesWritten)),
	)
	if ret == 0 {
		return err
	}
	return nil
}

// Memory_WriteInt64 向目标进程写入 64 位有符号整数。
// 参数 hProcess：进程句柄
// 参数 address：目标内存地址
// 参数 value：要写入的 64 位整数值
// 返回 error：失败时返回错误信息
func Memory_WriteInt64(hProcess syscall.Handle, address uintptr, value int64) error {
	buffer := [8]byte{
		byte(value), byte(value >> 8), byte(value >> 16), byte(value >> 24),
		byte(value >> 32), byte(value >> 40), byte(value >> 48), byte(value >> 56),
	}
	var numberOfBytesWritten uintptr
	ret, _, err := procWriteProcessMemory.Call(
		uintptr(hProcess), address,
		uintptr(unsafe.Pointer(&buffer[0])), 8,
		uintptr(unsafe.Pointer(&numberOfBytesWritten)),
	)
	if ret == 0 {
		return err
	}
	return nil
}

// Memory_WriteBytes 向目标进程写入原始字节数据。
// 参数 hProcess：进程句柄
// 参数 address：目标内存地址
// 参数 data：要写入的字节数据
// 返回 error：失败时返回错误信息
func Memory_WriteBytes(hProcess syscall.Handle, address uintptr, data []byte) error {
	var numberOfBytesWritten uintptr
	ret, _, err := procWriteProcessMemory.Call(
		uintptr(hProcess), address,
		uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)),
		uintptr(unsafe.Pointer(&numberOfBytesWritten)),
	)
	if ret == 0 {
		return err
	}
	return nil
}

// Memory_CloseHandle 关闭进程句柄，释放系统资源。
// 参数 hProcess：进程句柄
// 返回 error：失败时返回错误信息
func Memory_CloseHandle(hProcess syscall.Handle) error {
	ret, _, err := procCloseHandleM.Call(uintptr(hProcess))
	if ret == 0 {
		return err
	}
	return nil
}

// Memory_SearchBytes 在目标进程内存中搜索二进制特征码（AOB 搜索）。
// 特征码中 0xFF 表示通配符（匹配任意字节）。
// 参数 hProcess：进程句柄
// 参数 pattern：要搜索的字节模式（0xFF 为通配符）
// 参数 startAddr：搜索起始内存地址
// 参数 length：搜索范围大小（字节）
// 返回 []uintptr：所有匹配地址的切片
// 返回 error：失败时返回错误信息
func Memory_SearchBytes(hProcess syscall.Handle, pattern []byte, startAddr uintptr, length uintptr) ([]uintptr, error) {
	buf := make([]byte, length)
	var read uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess), startAddr,
		uintptr(unsafe.Pointer(&buf[0])), length,
		uintptr(unsafe.Pointer(&read)),
	)
	if ret == 0 {
		return nil, err
	}

	var results []uintptr
	patLen := len(pattern)
	for i := uintptr(0); i < read-uintptr(patLen)+1; i++ {
		match := true
		for j := 0; j < patLen; j++ {
			if pattern[j] != 0xFF && buf[i+uintptr(j)] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			results = append(results, startAddr+i)
		}
	}
	return results, nil
}

// Memory_SearchText 在目标进程内存中搜索文本字符串。
// 参数 hProcess：进程句柄
// 参数 text：要搜索的文本字符串
// 参数 startAddr：搜索起始内存地址
// 参数 length：搜索范围大小（字节）
// 返回 []uintptr：所有匹配地址的切片
// 返回 error：失败时返回错误信息
func Memory_SearchText(hProcess syscall.Handle, text string, startAddr uintptr, length uintptr) ([]uintptr, error) {
	return Memory_SearchBytes(hProcess, []byte(text), startAddr, length)
}

// Memory_Alloc 在目标进程中分配可执行读写内存。
// 参数 hProcess：进程句柄
// 参数 size：分配大小（字节）
// 返回 uintptr：分配的内存地址
// 返回 error：失败时返回错误信息
func Memory_Alloc(hProcess syscall.Handle, size uintptr) (uintptr, error) {
	addr, _, err := procVirtualAllocEx.Call(
		uintptr(hProcess), 0, size,
		uintptr(uint32(windows.MEM_COMMIT|windows.MEM_RESERVE)),
		uintptr(PAGE_EXECUTE_READWRITE),
	)
	if addr == 0 {
		return 0, err
	}
	return addr, nil
}

// Memory_Free 释放目标进程中已分配的内存区域。
// 参数 hProcess：进程句柄
// 参数 address：要释放的内存地址
// 返回 error：失败时返回错误信息
func Memory_Free(hProcess syscall.Handle, address uintptr) error {
	ret, _, err := procVirtualFreeEx.Call(uintptr(hProcess), address, 0, uintptr(windows.MEM_RELEASE))
	if ret == 0 {
		return err
	}
	return nil
}

// Memory_Protect 修改目标进程内存区域的保护属性。
// 参数 hProcess：进程句柄
// 参数 address：目标内存地址
// 参数 size：区域大小
// 参数 newProtect：新保护属性（PAGE_READWRITE / PAGE_EXECUTE_READWRITE 等）
// 返回 uint32：修改前的旧保护属性
// 返回 error：失败时返回错误信息
func Memory_Protect(hProcess syscall.Handle, address uintptr, size uintptr, newProtect uint32) (uint32, error) {
	var oldProtect uint32
	ret, _, err := procVirtualProtectEx.Call(
		uintptr(hProcess), address, size,
		uintptr(newProtect),
		uintptr(unsafe.Pointer(&oldProtect)),
	)
	if ret == 0 {
		return 0, err
	}
	return oldProtect, nil
}

// Memory_Query 查询目标进程指定地址的内存信息。
// 参数 hProcess：进程句柄
// 参数 address：要查询的内存地址
// 返回 *MEMORY_BASIC_INFORMATION：内存区域信息（基址、大小、保护属性等）
// 返回 error：失败时返回错误信息
func Memory_Query(hProcess syscall.Handle, address uintptr) (*MEMORY_BASIC_INFORMATION, error) {
	var mbi MEMORY_BASIC_INFORMATION
	ret, _, err := procVirtualQueryEx.Call(
		uintptr(hProcess), address,
		uintptr(unsafe.Pointer(&mbi)),
		unsafe.Sizeof(mbi),
	)
	if ret == 0 {
		return nil, err
	}
	return &mbi, nil
}

// Memory_GetProcessID 通过进程名称获取进程 ID。
// 参数 processName：进程可执行文件名（如 "notepad.exe"）
// 返回 uint32：进程 ID
// 返回 error：未找到时返回错误
func Memory_GetProcessID(processName string) (uint32, error) {
	snapshot, _, err := procCreateToolhelp32Snapshot.Call(uintptr(TH32CS_SNAPPROCESS), 0)
	if snapshot == uintptr(syscall.InvalidHandle) {
		return 0, err
	}
	defer procCloseHandleM.Call(snapshot)

	var pe32 PROCESSENTRY32W
	pe32.DwSize = uint32(unsafe.Sizeof(pe32))

	ret, _, _ := procProcess32FirstW.Call(snapshot, uintptr(unsafe.Pointer(&pe32)))
	if ret == 0 {
		return 0, errors.New("unable to enumerate processes")
	}

	for {
		name := syscall.UTF16ToString(pe32.SzExeFile[:])
		if name == processName {
			return pe32.Th32ProcessID, nil
		}
		ret, _, _ = procProcess32NextW.Call(snapshot, uintptr(unsafe.Pointer(&pe32)))
		if ret == 0 {
			break
		}
	}
	return 0, errors.New("process not found: " + processName)
}

// Memory_EnumProcesses 枚举系统中所有正在运行的进程。
// 返回 []PROCESSENTRY32W：进程信息列表
// 返回 error：失败时返回错误
func Memory_EnumProcesses() ([]PROCESSENTRY32W, error) {
	snapshot, _, err := procCreateToolhelp32Snapshot.Call(uintptr(TH32CS_SNAPPROCESS), 0)
	if snapshot == uintptr(syscall.InvalidHandle) {
		return nil, err
	}
	defer procCloseHandleM.Call(snapshot)

	var pe32 PROCESSENTRY32W
	pe32.DwSize = uint32(unsafe.Sizeof(pe32))

	var result []PROCESSENTRY32W
	ret, _, _ := procProcess32FirstW.Call(snapshot, uintptr(unsafe.Pointer(&pe32)))
	if ret == 0 {
		return nil, errors.New("unable to enumerate processes")
	}

	for {
		result = append(result, pe32)
		ret, _, _ = procProcess32NextW.Call(snapshot, uintptr(unsafe.Pointer(&pe32)))
		if ret == 0 {
			break
		}
	}
	return result, nil
}

// Memory_GetModuleBase 获取目标进程中指定模块的加载基址。
// 参数 processID：目标进程 ID
// 参数 moduleName：模块文件名（如 "kernel32.dll"）
// 返回 uintptr：模块基址
// 返回 error：未找到时返回错误
func Memory_GetModuleBase(processID uint32, moduleName string) (uintptr, error) {
	snapshot, _, err := procCreateToolhelp32Snapshot.Call(
		uintptr(TH32CS_SNAPMODULE|TH32CS_SNAPMODULE32),
		uintptr(processID),
	)
	if snapshot == uintptr(syscall.InvalidHandle) {
		return 0, err
	}
	defer procCloseHandleM.Call(snapshot)

	var me32 MODULEENTRY32W
	me32.DwSize = uint32(unsafe.Sizeof(me32))

	ret, _, _ := procModule32FirstW.Call(snapshot, uintptr(unsafe.Pointer(&me32)))
	if ret == 0 {
		return 0, errors.New("unable to enumerate modules")
	}

	want, _ := syscall.UTF16FromString(moduleName)
	for {
		name := me32.SzModule[:]
		match := true
		for i := 0; i < len(want) && want[i] != 0; i++ {
			c1, c2 := want[i], name[i]
			if c1 >= 'A' && c1 <= 'Z' {
				c1 += 32
			}
			if c2 >= 'A' && c2 <= 'Z' {
				c2 += 32
			}
			if c1 != c2 {
				match = false
				break
			}
		}
		if match {
			return me32.ModBaseAddr, nil
		}
		ret, _, _ = procModule32NextW.Call(snapshot, uintptr(unsafe.Pointer(&me32)))
		if ret == 0 {
			break
		}
	}
	return 0, errors.New("module not found: " + moduleName)
}

// Memory_TerminateProcess 强制终止指定进程。
// 参数 processID：目标进程 ID
// 返回 error：失败时返回错误信息
func Memory_TerminateProcess(processID uint32) error {
	hProcess, err := Memory_OpenProcess(processID, PROCESS_TERMINATE)
	if err != nil {
		return err
	}
	defer Memory_CloseHandle(hProcess)

	ret, _, err := procTerminateProcessM.Call(uintptr(hProcess), 0)
	if ret == 0 {
		return err
	}
	return nil
}

// Memory_GetProcessName 获取指定进程 ID 对应的可执行文件名称。
// 参数 processID：目标进程 ID
// 返回 string：进程名称
// 返回 []uint16：原始 UTF-16 名称数组
// 返回 error：失败时返回错误
func Memory_GetProcessName(processID uint32) (string, []uint16, error) {
	snapshot, _, err := procCreateToolhelp32Snapshot.Call(uintptr(TH32CS_SNAPPROCESS), 0)
	if snapshot == uintptr(syscall.InvalidHandle) {
		return "", nil, err
	}
	defer procCloseHandleM.Call(snapshot)

	var pe32 PROCESSENTRY32W
	pe32.DwSize = uint32(unsafe.Sizeof(pe32))

	ret, _, _ := procProcess32FirstW.Call(snapshot, uintptr(unsafe.Pointer(&pe32)))
	if ret == 0 {
		return "", nil, errors.New("unable to enumerate processes")
	}

	for {
		if pe32.Th32ProcessID == processID {
			return syscall.UTF16ToString(pe32.SzExeFile[:]), pe32.SzExeFile[:], nil
		}
		ret, _, _ = procProcess32NextW.Call(snapshot, uintptr(unsafe.Pointer(&pe32)))
		if ret == 0 {
			break
		}
	}
	return "", nil, errors.New("process not found")
}

// Memory_IsProcessRunning 检查指定进程是否正在运行。
// 参数 processID：目标进程 ID
// 返回 bool：true 表示进程存在且正在运行（退出码为 STILL_ACTIVE=259）
func Memory_IsProcessRunning(processID uint32) bool {
	hProcess, err := Memory_OpenProcess(processID, PROCESS_QUERY_INFORMATION)
	if err != nil {
		return false
	}
	defer Memory_CloseHandle(hProcess)

	var exitCode uint32
	ret, _, _ := procGetExitCodeProcessM.Call(uintptr(hProcess), uintptr(unsafe.Pointer(&exitCode)))
	if ret == 0 {
		return false
	}
	return exitCode == 259
}

// Memory_GetProcessBits 检测目标进程是 32 位还是 64 位。
// 参数 processID：目标进程 ID
// 返回 int：32 或 64，失败时返回 0
func Memory_GetProcessBits(processID uint32) int {
	handle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION, false, processID)
	if err != nil {
		return 0
	}
	defer windows.CloseHandle(handle)

	var isWow64 bool
	err = windows.IsWow64Process(handle, &isWow64)
	if err != nil {
		return 0
	}
	if isWow64 {
		return 32
	}
	return 64
}