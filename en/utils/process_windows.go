//go:build windows

// Windows 进程操作模块
// 提供进程创建、终止、枚举、优先级设置等功能。
// 基于 Windows kernel32.dll API 实现。
package utils

import (
	"syscall"
	"unsafe"
)

var (
	kernel32                   = syscall.NewLazyDLL("kernel32.dll")
	procCreateProcessW         = kernel32.NewProc("CreateProcessW")
	procOpenProcess            = kernel32.NewProc("OpenProcess")
	procTerminateProcess       = kernel32.NewProc("TerminateProcess")
	procGetExitCodeProcess     = kernel32.NewProc("GetExitCodeProcess")
	procWaitForSingleObject    = kernel32.NewProc("WaitForSingleObject")
	procCloseHandle            = kernel32.NewProc("CloseHandle")
	procGetProcessId           = kernel32.NewProc("GetProcessId")
	procGetCurrentProcessId    = kernel32.NewProc("GetCurrentProcessId")
	procGetCurrentProcess      = kernel32.NewProc("GetCurrentProcess")
	procSetPriorityClass       = kernel32.NewProc("SetPriorityClass")
	procGetPriorityClass       = kernel32.NewProc("GetPriorityClass")
	procGetModuleFileNameW     = kernel32.NewProc("GetModuleFileNameW")
	procCreateToolhelp32Snapshot = kernel32.NewProc("CreateToolhelp32Snapshot")
	procProcess32FirstW        = kernel32.NewProc("Process32FirstW")
	procProcess32NextW         = kernel32.NewProc("Process32NextW")
)

const (
	PROCESS_TERMINATE                  = 0x0001
	PROCESS_QUERY_INFORMATION          = 0x0400
	PROCESS_QUERY_LIMITED_INFORMATION = 0x1000
	STILL_ACTIVE                       = 259
	INFINITE                           = 0xFFFFFFFF
	TH32CS_SNAPPROCESS                 = 0x00000002

	IDLE_PRIORITY_CLASS         = 0x0040
	BELOW_NORMAL_PRIORITY_CLASS = 0x4000
	NORMAL_PRIORITY_CLASS       = 0x0020
	ABOVE_NORMAL_PRIORITY_CLASS = 0x8000
	HIGH_PRIORITY_CLASS         = 0x0080
	REALTIME_PRIORITY_CLASS     = 0x0100
)

type PROCESSENTRY32W struct {
	Size              uint32
	Usage             uint32
	ProcessID         uint32
	DefaultHeapID     uintptr
	ModuleID          uint32
	Threads           uint32
	ParentProcessID   uint32
	PriorityClassBase int32
	Flags             uint32
	ExeFile           [260]uint16
}

type STARTUPINFOW struct {
	Cb            uint32
	LPReserved    *uint16
	LPDesktop     *uint16
	LPTitle       *uint16
	DwX           uint32
	DwY           uint32
	DwXSize       uint32
	DwYSize       uint32
	DwXCountChars uint32
	DwYCountChars uint32
	DwFillAttribute uint32
	DwFlags       uint32
	WShowWindow   uint16
	CbReserved2   uint16
	LPReserved2   *byte
	HStdInput     syscall.Handle
	HStdOutput    syscall.Handle
	HStdError     syscall.Handle
}

type PROCESS_INFORMATION struct {
	Process   syscall.Handle
	Thread    syscall.Handle
	ProcessID uint32
	ThreadID  uint32
}

// Process_Create 创建一个新的进程。
// 启动指定程序，返回进程信息和句柄。
// 使用完毕后应调用 Process_CloseHandle 释放资源。
func Process_Create(programPath string, commandLine string, workDir string) (*PROCESS_INFORMATION, error) {
	pathPtr, _ := syscall.UTF16PtrFromString(programPath)
	var cmdPtr *uint16
	if commandLine != "" {
		cmdPtr, _ = syscall.UTF16PtrFromString(commandLine)
	}
	var dirPtr *uint16
	if workDir != "" {
		dirPtr, _ = syscall.UTF16PtrFromString(workDir)
	}

	var si STARTUPINFOW
	si.Cb = uint32(unsafe.Sizeof(si))
	var pi PROCESS_INFORMATION

	ret, _, err := procCreateProcessW.Call(
		uintptr(unsafe.Pointer(pathPtr)),
		uintptr(unsafe.Pointer(cmdPtr)),
		0, 0, 0, 0, 0,
		uintptr(unsafe.Pointer(dirPtr)),
		uintptr(unsafe.Pointer(&si)),
		uintptr(unsafe.Pointer(&pi)),
	)
	if ret == 0 {
		return nil, err
	}
	return &pi, nil
}

// Process_Open 打开一个已存在的进程，获取其句柄。
func Process_Open(processID uint32, access uint32) (syscall.Handle, error) {
	ret, _, err := procOpenProcess.Call(uintptr(access), 0, uintptr(processID))
	if ret == 0 {
		return 0, err
	}
	return syscall.Handle(ret), nil
}

// Process_Terminate 终止指定进程。
func Process_Terminate(processID uint32, exitCode uint32) error {
	h, err := Process_Open(processID, PROCESS_TERMINATE)
	if err != nil {
		return err
	}
	defer Process_CloseHandle(h)

	ret, _, _ := procTerminateProcess.Call(uintptr(h), uintptr(exitCode))
	if ret == 0 {
		return syscall.GetLastError()
	}
	return nil
}

// Process_IsAlive 检查指定进程是否仍在运行。
func Process_IsAlive(processID uint32) bool {
	h, err := Process_Open(processID, PROCESS_QUERY_LIMITED_INFORMATION)
	if err != nil {
		return false
	}
	defer Process_CloseHandle(h)

	var exitCode uint32
	ret, _, _ := procGetExitCodeProcess.Call(uintptr(h), uintptr(unsafe.Pointer(&exitCode)))
	if ret == 0 {
		return false
	}
	return exitCode == STILL_ACTIVE
}

// Process_Wait 等待进程退出（阻塞）。
func Process_Wait(hProcess syscall.Handle, timeoutMs uint32) uint32 {
	ret, _, _ := procWaitForSingleObject.Call(uintptr(hProcess), uintptr(timeoutMs))
	return uint32(ret)
}

// Process_CloseHandle 关闭进程或线程句柄。
func Process_CloseHandle(handle syscall.Handle) bool {
	ret, _, _ := procCloseHandle.Call(uintptr(handle))
	return ret != 0
}

// Process_CurrentID 获取当前进程 ID。
func Process_CurrentID() uint32 {
	ret, _, _ := procGetCurrentProcessId.Call()
	return uint32(ret)
}

// Process_GetID 通过进程句柄获取进程 ID。
func Process_GetID(hProcess syscall.Handle) uint32 {
	ret, _, _ := procGetProcessId.Call(uintptr(hProcess))
	return uint32(ret)
}

// Process_GetExitCode 获取进程的退出代码。
func Process_GetExitCode(hProcess syscall.Handle) (uint32, bool) {
	var exitCode uint32
	ret, _, _ := procGetExitCodeProcess.Call(uintptr(hProcess), uintptr(unsafe.Pointer(&exitCode)))
	return exitCode, ret != 0
}

// Process_SetPriority 设置进程优先级。
func Process_SetPriority(hProcess syscall.Handle, priorityClass uint32) bool {
	ret, _, _ := procSetPriorityClass.Call(uintptr(hProcess), uintptr(priorityClass))
	return ret != 0
}

// Process_GetPriority 获取进程优先级。
func Process_GetPriority(hProcess syscall.Handle) (uint32, bool) {
	ret, _, _ := procGetPriorityClass.Call(uintptr(hProcess))
	return uint32(ret), ret != 0
}

// Process_Enum 枚举系统中所有正在运行的进程。
func Process_Enum() ([]PROCESSENTRY32W, error) {
	hSnap, _, err := procCreateToolhelp32Snapshot.Call(TH32CS_SNAPPROCESS, 0)
	if hSnap == ^uintptr(0) {
		return nil, err
	}
	defer Process_CloseHandle(syscall.Handle(hSnap))

	var pe PROCESSENTRY32W
	pe.Size = uint32(unsafe.Sizeof(pe))

	ret, _, _ := procProcess32FirstW.Call(hSnap, uintptr(unsafe.Pointer(&pe)))
	if ret == 0 {
		return nil, syscall.GetLastError()
	}

	var list []PROCESSENTRY32W
	for {
		list = append(list, pe)
		pe.Size = uint32(unsafe.Sizeof(pe))
		ret, _, _ = procProcess32NextW.Call(hSnap, uintptr(unsafe.Pointer(&pe)))
		if ret == 0 {
			break
		}
	}
	return list, nil
}

// Process_FindByName 按进程名查找所有匹配的进程 ID。
func Process_FindByName(name string) ([]uint32, error) {
	list, err := Process_Enum()
	if err != nil {
		return nil, err
	}
	lowerName := Text_ToLower(name)
	var result []uint32
	for _, p := range list {
		curName := Text_ToLower(syscall.UTF16ToString(p.ExeFile[:]))
		if curName == lowerName {
			result = append(result, p.ProcessID)
		}
	}
	return result, nil
}

// Process_GetModulePath 获取指定进程的可执行文件完整路径。
func Process_GetModulePath(processID uint32) (string, error) {
	h, err := Process_Open(processID, PROCESS_QUERY_LIMITED_INFORMATION)
	if err != nil {
		return "", err
	}
	defer Process_CloseHandle(h)

	buf := make([]uint16, 260)
	ret, _, _ := procGetModuleFileNameW.Call(uintptr(h), uintptr(unsafe.Pointer(&buf[0])), 260)
	if ret == 0 {
		return "", syscall.GetLastError()
	}
	return syscall.UTF16ToString(buf[:ret]), nil
}

// Process_GetParentID 获取指定进程的父进程 ID。
func Process_GetParentID(processID uint32) (uint32, error) {
	list, err := Process_Enum()
	if err != nil {
		return 0, err
	}
	for _, p := range list {
		if p.ProcessID == processID {
			return p.ParentProcessID, nil
		}
	}
	return 0, syscall.ERROR_NOT_FOUND
}