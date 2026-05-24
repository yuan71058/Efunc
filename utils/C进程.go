//go:build windows

package utils

import (
	"syscall"
	"unsafe"
)

var (
	kernel32                  = syscall.NewLazyDLL("kernel32.dll")
	procCreateProcessW        = kernel32.NewProc("CreateProcessW")
	procOpenProcess           = kernel32.NewProc("OpenProcess")
	procTerminateProcess      = kernel32.NewProc("TerminateProcess")
	procGetExitCodeProcess    = kernel32.NewProc("GetExitCodeProcess")
	procWaitForSingleObject   = kernel32.NewProc("WaitForSingleObject")
	procCloseHandle           = kernel32.NewProc("CloseHandle")
	procGetProcessId          = kernel32.NewProc("GetProcessId")
	procGetCurrentProcessId   = kernel32.NewProc("GetCurrentProcessId")
	procGetCurrentProcess     = kernel32.NewProc("GetCurrentProcess")
	procSetPriorityClass      = kernel32.NewProc("SetPriorityClass")
	procGetPriorityClass      = kernel32.NewProc("GetPriorityClass")
	procGetModuleFileNameW    = kernel32.NewProc("GetModuleFileNameW")
	procCreateToolhelp32Snapshot = kernel32.NewProc("CreateToolhelp32Snapshot")
	procProcess32FirstW       = kernel32.NewProc("Process32FirstW")
	procProcess32NextW        = kernel32.NewProc("Process32NextW")
)

const (
	PROCESS_TERMINATE         = 0x0001
	PROCESS_QUERY_INFORMATION = 0x0400
	PROCESS_QUERY_LIMITED_INFORMATION = 0x1000
	STILL_ACTIVE              = 259
	INFINITE                  = 0xFFFFFFFF
	TH32CS_SNAPPROCESS        = 0x00000002

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

// C进程_创建 创建一个新的进程。
// 启动指定程序，返回进程信息和句柄。
// 使用完毕后应调用 C进程_关闭句柄 释放资源。
//
// 参数:
//   - 程序路径: 可执行文件的完整路径
//   - 命令行: 命令行参数，可为空
//   - 工作目录: 进程工作目录，可为空（使用父进程目录）
//
// 返回:
//   - *PROCESS_INFORMATION: 进程信息（含句柄和 ID）
//   - error: 创建失败时返回错误
func C进程_创建(程序路径 string, 命令行 string, 工作目录 string) (*PROCESS_INFORMATION, error) {
	程序路径Ptr, _ := syscall.UTF16PtrFromString(程序路径)
	var 命令行Ptr *uint16
	if 命令行 != "" {
		命令行Ptr, _ = syscall.UTF16PtrFromString(命令行)
	}
	var 工作目录Ptr *uint16
	if 工作目录 != "" {
		工作目录Ptr, _ = syscall.UTF16PtrFromString(工作目录)
	}

	var 启动信息 STARTUPINFOW
	启动信息.Cb = uint32(unsafe.Sizeof(启动信息))
	var 进程信息 PROCESS_INFORMATION

	ret, _, err := procCreateProcessW.Call(
		uintptr(unsafe.Pointer(程序路径Ptr)),
		uintptr(unsafe.Pointer(命令行Ptr)),
		0, 0,
		0,
		0,
		0,
		uintptr(unsafe.Pointer(工作目录Ptr)),
		uintptr(unsafe.Pointer(&启动信息)),
		uintptr(unsafe.Pointer(&进程信息)),
	)
	if ret == 0 {
		return nil, err
	}
	return &进程信息, nil
}

// C进程_打开 打开一个已存在的进程，获取其句柄。
// 需要指定访问权限，如 PROCESS_TERMINATE、PROCESS_QUERY_INFORMATION。
//
// 参数:
//   - 进程ID: 目标进程的 ID
//   - 访问权限: 请求的访问权限标志
//
// 返回:
//   - syscall.Handle: 进程句柄
//   - error: 打开失败时返回错误
func C进程_打开(进程ID uint32, 访问权限 uint32) (syscall.Handle, error) {
	ret, _, err := procOpenProcess.Call(uintptr(访问权限), 0, uintptr(进程ID))
	if ret == 0 {
		return 0, err
	}
	return syscall.Handle(ret), nil
}

// C进程_终止 终止指定进程。
// 先打开进程获取句柄，再终止，最后关闭句柄。
//
// 参数:
//   - 进程ID: 目标进程的 ID
//   - 退出码: 进程的退出代码，通常为 1
//
// 返回:
//   - error: 终止失败时返回错误
func C进程_终止(进程ID uint32, 退出码 uint32) error {
	句柄, err := C进程_打开(进程ID, PROCESS_TERMINATE)
	if err != nil {
		return err
	}
	defer C进程_关闭句柄(句柄)

	ret, _, _ := procTerminateProcess.Call(uintptr(句柄), uintptr(退出码))
	if ret == 0 {
		return syscall.GetLastError()
	}
	return nil
}

// C进程_是否存活 检查指定进程是否仍在运行。
//
// 参数:
//   - 进程ID: 目标进程的 ID
//
// 返回:
//   - bool: 进程存活返回 true，已退出返回 false
func C进程_是否存活(进程ID uint32) bool {
	句柄, err := C进程_打开(进程ID, PROCESS_QUERY_LIMITED_INFORMATION)
	if err != nil {
		return false
	}
	defer C进程_关闭句柄(句柄)

	var 退出码 uint32
	ret, _, _ := procGetExitCodeProcess.Call(uintptr(句柄), uintptr(unsafe.Pointer(&退出码)))
	if ret == 0 {
		return false
	}
	return 退出码 == STILL_ACTIVE
}

// C进程_等待 等待进程退出。
// 阻塞当前 goroutine 直到目标进程结束或超时。
//
// 参数:
//   - 进程句柄: 进程句柄（从 C进程_创建 或 C进程_打开 获取）
//   - 超时毫秒: 等待超时时间（毫秒），INFINITE 表示无限等待
//
// 返回:
//   - uint32: 等待结果（0 表示成功，258 表示超时）
func C进程_等待(进程句柄 syscall.Handle, 超时毫秒 uint32) uint32 {
	ret, _, _ := procWaitForSingleObject.Call(uintptr(进程句柄), uintptr(超时毫秒))
	return uint32(ret)
}

// C进程_关闭句柄 关闭进程或线程句柄，释放系统资源。
// 使用 C进程_创建 后必须关闭进程句柄和线程句柄。
//
// 参数:
//   - 句柄: 要关闭的句柄
//
// 返回:
//   - bool: 关闭成功返回 true
func C进程_关闭句柄(句柄 syscall.Handle) bool {
	ret, _, _ := procCloseHandle.Call(uintptr(句柄))
	return ret != 0
}

// C进程_取当前ID 获取当前进程的 ID。
//
// 返回:
//   - uint32: 当前进程 ID
func C进程_取当前ID() uint32 {
	ret, _, _ := procGetCurrentProcessId.Call()
	return uint32(ret)
}

// C进程_取ID 通过进程句柄获取进程 ID。
//
// 参数:
//   - 进程句柄: 进程句柄
//
// 返回:
//   - uint32: 进程 ID
func C进程_取ID(进程句柄 syscall.Handle) uint32 {
	ret, _, _ := procGetProcessId.Call(uintptr(进程句柄))
	return uint32(ret)
}

// C进程_取退出码 获取进程的退出代码。
// 退出代码为 STILL_ACTIVE(259) 表示进程仍在运行。
//
// 参数:
//   - 进程句柄: 进程句柄
//
// 返回:
//   - uint32: 退出代码
//   - bool: 获取成功返回 true
func C进程_取退出码(进程句柄 syscall.Handle) (uint32, bool) {
	var 退出码 uint32
	ret, _, _ := procGetExitCodeProcess.Call(uintptr(进程句柄), uintptr(unsafe.Pointer(&退出码)))
	return 退出码, ret != 0
}

// C进程_设置优先级 设置进程的优先级。
// 常用优先级：IDLE_PRIORITY_CLASS(空闲)、NORMAL_PRIORITY_CLASS(正常)、
// HIGH_PRIORITY_CLASS(高)、REALTIME_PRIORITY_CLASS(实时)。
//
// 参数:
//   - 进程句柄: 进程句柄
//   - 优先级: 优先级标志
//
// 返回:
//   - bool: 设置成功返回 true
func C进程_设置优先级(进程句柄 syscall.Handle, 优先级 uint32) bool {
	ret, _, _ := procSetPriorityClass.Call(uintptr(进程句柄), uintptr(优先级))
	return ret != 0
}

// C进程_取优先级 获取进程的优先级。
//
// 参数:
//   - 进程句柄: 进程句柄
//
// 返回:
//   - uint32: 优先级标志
//   - bool: 获取成功返回 true
func C进程_取优先级(进程句柄 syscall.Handle) (uint32, bool) {
	ret, _, _ := procGetPriorityClass.Call(uintptr(进程句柄))
	return uint32(ret), ret != 0
}

// C进程_枚举 枚举系统中所有正在运行的进程。
// 返回进程 ID 和可执行文件名的列表。
//
// 返回:
//   - []PROCESSENTRY32W: 进程信息列表
//   - error: 枚举失败时返回错误
func C进程_枚举() ([]PROCESSENTRY32W, error) {
	快照句柄, _, err := procCreateToolhelp32Snapshot.Call(TH32CS_SNAPPROCESS, 0)
	if 快照句柄 == ^uintptr(0) {
		return nil, err
	}
	defer C进程_关闭句柄(syscall.Handle(快照句柄))

	var 进程信息 PROCESSENTRY32W
	进程信息.Size = uint32(unsafe.Sizeof(进程信息))

	ret, _, _ := procProcess32FirstW.Call(快照句柄, uintptr(unsafe.Pointer(&进程信息)))
	if ret == 0 {
		return nil, syscall.GetLastError()
	}

	var 列表 []PROCESSENTRY32W
	for {
		列表 = append(列表, 进程信息)
		进程信息.Size = uint32(unsafe.Sizeof(进程信息))
		ret, _, _ = procProcess32NextW.Call(快照句柄, uintptr(unsafe.Pointer(&进程信息)))
		if ret == 0 {
			break
		}
	}
	return 列表, nil
}

// C进程_按名查找 按进程名查找所有匹配的进程 ID。
// 进程名为可执行文件名（如 "notepad.exe"），不区分大小写。
//
// 参数:
//   - 进程名: 可执行文件名
//
// 返回:
//   - []uint32: 匹配的进程 ID 列表
//   - error: 枚举失败时返回错误
func C进程_按名查找(进程名 string) ([]uint32, error) {
	列表, err := C进程_枚举()
	if err != nil {
		return nil, err
	}

	进程名小写 := W文本_到小写(进程名)
	var 结果 []uint32
	for _, 进程 := range 列表 {
		当前名称 := W文本_到小写(syscall.UTF16ToString(进程.ExeFile[:]))
		if 当前名称 == 进程名小写 {
			结果 = append(结果, 进程.ProcessID)
		}
	}
	return 结果, nil
}

// C进程_取模块路径 获取指定进程的可执行文件完整路径。
// 需要进程具有 PROCESS_QUERY_LIMITED_INFORMATION 权限。
//
// 参数:
//   - 进程ID: 目标进程的 ID
//
// 返回:
//   - string: 可执行文件的完整路径
//   - error: 获取失败时返回错误
func C进程_取模块路径(进程ID uint32) (string, error) {
	句柄, err := C进程_打开(进程ID, PROCESS_QUERY_LIMITED_INFORMATION)
	if err != nil {
		return "", err
	}
	defer C进程_关闭句柄(句柄)

	buf := make([]uint16, 260)
	ret, _, _ := procGetModuleFileNameW.Call(uintptr(句柄), uintptr(unsafe.Pointer(&buf[0])), 260)
	if ret == 0 {
		return "", syscall.GetLastError()
	}
	return syscall.UTF16ToString(buf[:ret]), nil
}

// C进程_取父进程ID 获取指定进程的父进程 ID。
//
// 参数:
//   - 进程ID: 目标进程的 ID
//
// 返回:
//   - uint32: 父进程 ID
//   - error: 获取失败时返回错误
func C进程_取父进程ID(进程ID uint32) (uint32, error) {
	列表, err := C进程_枚举()
	if err != nil {
		return 0, err
	}
	for _, 进程 := range 列表 {
		if 进程.ProcessID == 进程ID {
			return 进程.ParentProcessID, nil
		}
	}
	return 0, syscall.ERROR_NOT_FOUND
}
