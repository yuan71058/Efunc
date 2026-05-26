//go:build windows

package utils

import (
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

var (
	user32X系统 = syscall.NewLazyDLL("user32.dll")
	kernel32X系统 = syscall.NewLazyDLL("kernel32.dll")
	advapi32  = syscall.NewLazyDLL("advapi32.dll")

	procExitWindowsEx     = user32X系统.NewProc("ExitWindowsEx")
	procOpenClipboard     = user32X系统.NewProc("OpenClipboard")
	procCloseClipboard    = user32X系统.NewProc("CloseClipboard")
	procEmptyClipboard    = user32X系统.NewProc("EmptyClipboard")
	procGetClipboardData  = user32X系统.NewProc("GetClipboardData")
	procSetClipboardData  = user32X系统.NewProc("SetClipboardData")
	procIsClipboardFormatAvailable = user32X系统.NewProc("IsClipboardFormatAvailable")
	procMessageBoxW       = user32X系统.NewProc("MessageBoxW")
	procGetSystemMetricsX = user32X系统.NewProc("GetSystemMetrics")
	procSystemParametersInfoW = user32X系统.NewProc("SystemParametersInfoW")
	procLockWorkStation   = user32X系统.NewProc("LockWorkStation")

	procGlobalAlloc     = kernel32X系统.NewProc("GlobalAlloc")
	procGlobalLock      = kernel32X系统.NewProc("GlobalLock")
	procGlobalUnlock    = kernel32X系统.NewProc("GlobalUnlock")
	procGlobalFree      = kernel32X系统.NewProc("GlobalFree")
	procGetComputerNameW = kernel32X系统.NewProc("GetComputerNameW")
	procGetUserNameW    = kernel32X系统.NewProc("GetUserNameW")

	procInitiateSystemShutdownExW = advapi32.NewProc("InitiateSystemShutdownExW")
	procLookupPrivilegeValueW     = advapi32.NewProc("LookupPrivilegeValueW")
	procAdjustTokenPrivileges     = advapi32.NewProc("AdjustTokenPrivileges")
)

const (
	EWX_LOGOFF   = 0x00000000
	EWX_SHUTDOWN = 0x00000001
	EWX_REBOOT   = 0x00000002
	EWX_FORCE    = 0x00000004
	EWX_POWEROFF = 0x00000008
	EWX_HYBRID_SHUTDOWN = 0x00400000

	CF_TEXT        = 1
	CF_UNICODETEXT = 13

	GHND = 0x0042

	MB_OK               = 0x00000000
	MB_OKCANCEL         = 0x00000001
	MB_YESNO            = 0x00000004
	MB_YESNOCANCEL      = 0x00000003
	MB_ICONINFORMATION  = 0x00000040
	MB_ICONWARNING      = 0x00000030
	MB_ICONERROR        = 0x00000010
	MB_ICONQUESTION     = 0x00000020

	IDOK     = 1
	IDCANCEL = 2
	IDYES    = 6
	IDNO     = 7

	SPI_SETSCREENSAVEACTIVE = 17
	SPI_GETSCREENSAVETIMEOUT = 14
	SPI_SETSCREENSAVETIMEOUT = 15

	TOKEN_ADJUST_PRIVILEGES = 0x0020
	TOKEN_QUERY             = 0x0008

	SE_PRIVILEGE_ENABLED = 0x00000002
)

type LUID struct {
	LowPart  uint32
	HighPart int32
}

type LUID_AND_ATTRIBUTES struct {
	Luid       LUID
	Attributes uint32
}

type TOKEN_PRIVILEGES struct {
	PrivilegeCount uint32
	Privileges     [1]LUID_AND_ATTRIBUTES
}

// ===================== 系统关机/重启/注销 =====================

func X系统_关机(强制 bool) error {
	var flags uint32 = EWX_SHUTDOWN | EWX_POWEROFF
	if 强制 {
		flags |= EWX_FORCE
	}
	return X系统_execExitWindows(flags)
}

func X系统_重启(强制 bool) error {
	var flags uint32 = EWX_REBOOT
	if 强制 {
		flags |= EWX_FORCE
	}
	return X系统_execExitWindows(flags)
}

func X系统_注销(强制 bool) error {
	var flags uint32 = EWX_LOGOFF
	if 强制 {
		flags |= EWX_FORCE
	}
	return X系统_execExitWindows(flags)
}

func X系统_execExitWindows(flags uint32) error {
	ret, _, err := procExitWindowsEx.Call(uintptr(flags), 0)
	if ret == 0 {
		return err
	}
	return nil
}

func X系统_锁定工作站() error {
	ret, _, err := procLockWorkStation.Call()
	if ret == 0 {
		return err
	}
	return nil
}

func X系统_启用关机权限() error {
	var hToken syscall.Handle
	hProc, _ := syscall.GetCurrentProcess()
	ret, _, err := advapi32.NewProc("OpenProcessToken").Call(
		uintptr(hProc),
		uintptr(TOKEN_ADJUST_PRIVILEGES|TOKEN_QUERY),
		uintptr(unsafe.Pointer(&hToken)),
	)
	if ret == 0 {
		return err
	}
	defer X线程_关闭句柄(hToken)

	var luid LUID
	name, _ := syscall.UTF16PtrFromString("SeShutdownPrivilege")
	ret, _, err = procLookupPrivilegeValueW.Call(0, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(&luid)))
	if ret == 0 {
		return err
	}

	tp := TOKEN_PRIVILEGES{
		PrivilegeCount: 1,
		Privileges: [1]LUID_AND_ATTRIBUTES{
			{Luid: luid, Attributes: SE_PRIVILEGE_ENABLED},
		},
	}
	ret, _, err = procAdjustTokenPrivileges.Call(
		uintptr(hToken), 0,
		uintptr(unsafe.Pointer(&tp)),
		uintptr(unsafe.Sizeof(tp)), 0, 0,
	)
	if ret == 0 {
		return err
	}
	return nil
}

func X系统_远程关机(计算机名 string, 消息 string, 超时秒 uint32, 强制 bool) error {
	compName, _ := syscall.UTF16PtrFromString(计算机名)
	msg, _ := syscall.UTF16PtrFromString(消息)
	var flags uint32
	if 强制 {
		flags = 1
	}
	ret, _, err := procInitiateSystemShutdownExW.Call(
		uintptr(unsafe.Pointer(compName)),
		uintptr(unsafe.Pointer(msg)),
		uintptr(超时秒),
		uintptr(flags),
		0,
	)
	if ret == 0 {
		return err
	}
	return nil
}

// ===================== 剪辑版操作 =====================

func X系统_置剪辑版文本(text string) error {
	ret, _, err := procOpenClipboard.Call(0)
	if ret == 0 {
		return err
	}
	defer procCloseClipboard.Call()

	procEmptyClipboard.Call()

	utf16Text, _ := syscall.UTF16FromString(text)
	size := len(utf16Text) * 2

	hMem, _, err := procGlobalAlloc.Call(uintptr(GHND), uintptr(size))
	if hMem == 0 {
		return err
	}

	ptr, _, err := procGlobalLock.Call(hMem)
	if ptr == 0 {
		procGlobalFree.Call(hMem)
		return err
	}

	dst := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), size)
	for i, c := range utf16Text {
		dst[i*2] = byte(c)
		dst[i*2+1] = byte(c >> 8)
	}

	procGlobalUnlock.Call(hMem)

	ret, _, err = procSetClipboardData.Call(uintptr(CF_UNICODETEXT), hMem)
	if ret == 0 {
		procGlobalFree.Call(hMem)
		return err
	}
	return nil
}

func X系统_取剪辑版文本() (string, error) {
	ret, _, err := procOpenClipboard.Call(0)
	if ret == 0 {
		return "", err
	}
	defer procCloseClipboard.Call()

	ret, _, _ = procIsClipboardFormatAvailable.Call(uintptr(CF_UNICODETEXT))
	if ret == 0 {
		return "", nil
	}

	hData, _, err := procGetClipboardData.Call(uintptr(CF_UNICODETEXT))
	if hData == 0 {
		return "", err
	}

	ptr, _, err := procGlobalLock.Call(hData)
	if ptr == 0 {
		return "", err
	}
	defer procGlobalUnlock.Call(hData)

	runes := make([]uint16, 0, 1024)
	for i := 0; ; i += 2 {
		b0 := *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
		b1 := *(*byte)(unsafe.Pointer(ptr + uintptr(i+1)))
		c := uint16(b0) | uint16(b1)<<8
		if c == 0 {
			break
		}
		runes = append(runes, c)
	}

	return syscall.UTF16ToString(runes), nil
}

func X系统_清空剪辑版() error {
	ret, _, err := procOpenClipboard.Call(0)
	if ret == 0 {
		return err
	}
	defer procCloseClipboard.Call()

	ret, _, err = procEmptyClipboard.Call()
	if ret == 0 {
		return err
	}
	return nil
}

// ===================== 消息框 =====================

func X系统_信息框(标题, 内容 string) {
	title, _ := syscall.UTF16PtrFromString(标题)
	text, _ := syscall.UTF16PtrFromString(内容)
	procMessageBoxW.Call(0, uintptr(unsafe.Pointer(text)), uintptr(unsafe.Pointer(title)), uintptr(MB_OK|MB_ICONINFORMATION))
}

func X系统_信息框_确认(标题, 内容 string) int {
	title, _ := syscall.UTF16PtrFromString(标题)
	text, _ := syscall.UTF16PtrFromString(内容)
	ret, _, _ := procMessageBoxW.Call(0, uintptr(unsafe.Pointer(text)), uintptr(unsafe.Pointer(title)), uintptr(MB_OKCANCEL|MB_ICONQUESTION))
	return int(ret)
}

func X系统_信息框_是否(标题, 内容 string) int {
	title, _ := syscall.UTF16PtrFromString(标题)
	text, _ := syscall.UTF16PtrFromString(内容)
	ret, _, _ := procMessageBoxW.Call(0, uintptr(unsafe.Pointer(text)), uintptr(unsafe.Pointer(title)), uintptr(MB_YESNO|MB_ICONQUESTION))
	return int(ret)
}

// ===================== 系统命令 =====================

func X系统_执行命令(命令 string, 参数 ...string) (string, error) {
	cmd := exec.Command(命令, 参数...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func X系统_执行命令隐藏(命令 string, 参数 ...string) error {
	cmd := exec.Command(命令, 参数...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.Start()
}

func X系统_取计算机名() (string, error) {
	var size uint32 = 256
	buf := make([]uint16, size)
	ret, _, err := procGetComputerNameW.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&size)))
	if ret == 0 {
		return "", err
	}
	return syscall.UTF16ToString(buf[:size]), nil
}

func X系统_取用户名() (string, error) {
	var size uint32 = 256
	buf := make([]uint16, size)
	ret, _, err := procGetUserNameW.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&size)))
	if ret == 0 {
		return "", err
	}
	return syscall.UTF16ToString(buf[:size-1]), nil
}

// ===================== 屏幕保护 =====================

func X系统_禁止屏幕保护(禁止 bool) error {
	var val uintptr
	if !禁止 {
		val = 1
	}
	ret, _, err := procSystemParametersInfoW.Call(uintptr(SPI_SETSCREENSAVEACTIVE), val, 0, 0)
	if ret == 0 {
		return err
	}
	return nil
}

func X系统_置屏保超时(秒 int) error {
	ret, _, err := procSystemParametersInfoW.Call(uintptr(SPI_SETSCREENSAVETIMEOUT), uintptr(秒), 0, 0)
	if ret == 0 {
		return err
	}
	return nil
}

// ===================== 环境变量 =====================

func X系统_置环境变量(名称, 值 string) error {
	return os.Setenv(名称, 值)
}

func X系统_取环境变量(名称 string) string {
	return os.Getenv(名称)
}

func X系统_删除环境变量(名称 string) error {
	return os.Unsetenv(名称)
}