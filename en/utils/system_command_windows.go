//go:build windows

// Windows 系统命令模块
// 提供系统关机/重启/注销、剪辑版操作、消息框、命令执行、屏幕保护控制等功能。
// 基于 Windows user32.dll/kernel32.dll/advapi32.dll API 实现。
package utils

import (
	"os"
	"os/exec"
	"reflect"
	"syscall"
	"unsafe"
)

var (
	user32SysCmd   = syscall.NewLazyDLL("user32.dll")
	kernel32SysCmd = syscall.NewLazyDLL("kernel32.dll")
	advapi32       = syscall.NewLazyDLL("advapi32.dll")

	procExitWindowsEx               = user32SysCmd.NewProc("ExitWindowsEx")
	procOpenClipboard               = user32SysCmd.NewProc("OpenClipboard")
	procCloseClipboard              = user32SysCmd.NewProc("CloseClipboard")
	procEmptyClipboard              = user32SysCmd.NewProc("EmptyClipboard")
	procGetClipboardData            = user32SysCmd.NewProc("GetClipboardData")
	procSetClipboardData            = user32SysCmd.NewProc("SetClipboardData")
	procIsClipboardFormatAvailable  = user32SysCmd.NewProc("IsClipboardFormatAvailable")
	procMessageBoxW                 = user32SysCmd.NewProc("MessageBoxW")
	procGetSystemMetricsX           = user32SysCmd.NewProc("GetSystemMetrics")
	procSystemParametersInfoW       = user32SysCmd.NewProc("SystemParametersInfoW")
	procLockWorkStation             = user32SysCmd.NewProc("LockWorkStation")

	procGlobalAlloc     = kernel32SysCmd.NewProc("GlobalAlloc")
	procGlobalLock      = kernel32SysCmd.NewProc("GlobalLock")
	procGlobalUnlock    = kernel32SysCmd.NewProc("GlobalUnlock")
	procGlobalFree      = kernel32SysCmd.NewProc("GlobalFree")
	procGetComputerNameW = kernel32SysCmd.NewProc("GetComputerNameW")
	procGetUserNameW    = kernel32SysCmd.NewProc("GetUserNameW")

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

	CF_TEXT        = 1
	CF_UNICODETEXT = 13
	GHND           = 0x0042

	MB_OK              = 0x00000000
	MB_OKCANCEL        = 0x00000001
	MB_YESNO           = 0x00000004
	MB_YESNOCANCEL     = 0x00000003
	MB_ICONINFORMATION = 0x00000040
	MB_ICONWARNING     = 0x00000030
	MB_ICONERROR       = 0x00000010
	MB_ICONQUESTION    = 0x00000020

	IDOK     = 1
	IDCANCEL = 2
	IDYES    = 6
	IDNO     = 7

	SPI_SETSCREENSAVEACTIVE = 17
	SPI_GETSCREENSAVETIMEOUT = 14
	SPI_SETSCREENSAVETIMEOUT = 15

	TOKEN_ADJUST_PRIVILEGES = 0x0020
	TOKEN_QUERY             = 0x0008
	SE_PRIVILEGE_ENABLED    = 0x00000002
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

// SystemCmd_Shutdown 关机。
func SystemCmd_Shutdown(force bool) error {
	var flags uint32 = EWX_SHUTDOWN | EWX_POWEROFF
	if force {
		flags |= EWX_FORCE
	}
	return systemCmdExitWindows(flags)
}

// SystemCmd_Reboot 重启。
func SystemCmd_Reboot(force bool) error {
	var flags uint32 = EWX_REBOOT
	if force {
		flags |= EWX_FORCE
	}
	return systemCmdExitWindows(flags)
}

// SystemCmd_Logoff 注销。
func SystemCmd_Logoff(force bool) error {
	var flags uint32 = EWX_LOGOFF
	if force {
		flags |= EWX_FORCE
	}
	return systemCmdExitWindows(flags)
}

func systemCmdExitWindows(flags uint32) error {
	ret, _, err := procExitWindowsEx.Call(uintptr(flags), 0)
	if ret == 0 {
		return err
	}
	return nil
}

// SystemCmd_LockWorkstation 锁定工作站。
func SystemCmd_LockWorkstation() error {
	ret, _, err := procLockWorkStation.Call()
	if ret == 0 {
		return err
	}
	return nil
}

// SystemCmd_EnableShutdownPrivilege 启用关机权限。
func SystemCmd_EnableShutdownPrivilege() error {
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
	defer Thread_CloseHandle(hToken)

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

// SystemCmd_RemoteShutdown 远程关机。
func SystemCmd_RemoteShutdown(computerName string, message string, timeoutSec uint32, force bool) error {
	cName, _ := syscall.UTF16PtrFromString(computerName)
	msg, _ := syscall.UTF16PtrFromString(message)
	var flags uint32
	if force {
		flags = 1
	}
	ret, _, err := procInitiateSystemShutdownExW.Call(
		uintptr(unsafe.Pointer(cName)),
		uintptr(unsafe.Pointer(msg)),
		uintptr(timeoutSec),
		uintptr(flags),
		0,
	)
	if ret == 0 {
		return err
	}
	return nil
}

// SystemCmd_SetClipboard 设置剪辑版文本。
func SystemCmd_SetClipboard(text string) error {
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

	var buf []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	sh.Data = ptr
	sh.Len = size
	sh.Cap = size
	for i, c := range utf16Text {
		buf[i*2] = byte(c)
		buf[i*2+1] = byte(c >> 8)
	}

	procGlobalUnlock.Call(hMem)

	ret, _, err = procSetClipboardData.Call(uintptr(CF_UNICODETEXT), hMem)
	if ret == 0 {
		procGlobalFree.Call(hMem)
		return err
	}
	return nil
}

// SystemCmd_GetClipboard 获取剪辑版文本。
func SystemCmd_GetClipboard() (string, error) {
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

	var src []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	sh.Data = ptr
	sh.Len = 4096
	sh.Cap = 4096

	runes := make([]uint16, 0, 1024)
	for i := 0; i < len(src)-1; i += 2 {
		c := uint16(src[i]) | uint16(src[i+1])<<8
		if c == 0 {
			break
		}
		runes = append(runes, c)
	}

	return syscall.UTF16ToString(runes), nil
}

// SystemCmd_ClearClipboard 清空剪辑版。
func SystemCmd_ClearClipboard() error {
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

// SystemCmd_MessageBox 显示信息框。
func SystemCmd_MessageBox(title, text string) {
	t, _ := syscall.UTF16PtrFromString(title)
	x, _ := syscall.UTF16PtrFromString(text)
	procMessageBoxW.Call(0, uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(t)), uintptr(MB_OK|MB_ICONINFORMATION))
}

// SystemCmd_MessageBoxConfirm 显示确认对话框。
func SystemCmd_MessageBoxConfirm(title, text string) int {
	t, _ := syscall.UTF16PtrFromString(title)
	x, _ := syscall.UTF16PtrFromString(text)
	ret, _, _ := procMessageBoxW.Call(0, uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(t)), uintptr(MB_OKCANCEL|MB_ICONQUESTION))
	return int(ret)
}

// SystemCmd_MessageBoxYesNo 显示是/否对话框。
func SystemCmd_MessageBoxYesNo(title, text string) int {
	t, _ := syscall.UTF16PtrFromString(title)
	x, _ := syscall.UTF16PtrFromString(text)
	ret, _, _ := procMessageBoxW.Call(0, uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(t)), uintptr(MB_YESNO|MB_ICONQUESTION))
	return int(ret)
}

// SystemCmd_Exec 执行系统命令并返回输出。
func SystemCmd_Exec(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// SystemCmd_ExecHidden 隐藏窗口执行系统命令。
func SystemCmd_ExecHidden(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.Start()
}

// SystemCmd_GetComputerName 获取计算机名。
func SystemCmd_GetComputerName() (string, error) {
	var size uint32 = 256
	buf := make([]uint16, size)
	ret, _, err := procGetComputerNameW.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&size)))
	if ret == 0 {
		return "", err
	}
	return syscall.UTF16ToString(buf[:size]), nil
}

// SystemCmd_GetUserName 获取当前用户名。
func SystemCmd_GetUserName() (string, error) {
	var size uint32 = 256
	buf := make([]uint16, size)
	ret, _, err := procGetUserNameW.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&size)))
	if ret == 0 {
		return "", err
	}
	return syscall.UTF16ToString(buf[:size-1]), nil
}

// SystemCmd_DisableScreenSaver 禁止/启用屏幕保护。
func SystemCmd_DisableScreenSaver(disable bool) error {
	var val uintptr
	if !disable {
		val = 1
	}
	ret, _, err := procSystemParametersInfoW.Call(uintptr(SPI_SETSCREENSAVEACTIVE), val, 0, 0)
	if ret == 0 {
		return err
	}
	return nil
}

// SystemCmd_SetScreenSaverTimeout 设置屏保超时时间（秒）。
func SystemCmd_SetScreenSaverTimeout(seconds int) error {
	ret, _, err := procSystemParametersInfoW.Call(uintptr(SPI_SETSCREENSAVETIMEOUT), uintptr(seconds), 0, 0)
	if ret == 0 {
		return err
	}
	return nil
}

// SystemCmd_SetEnv 设置环境变量。
func SystemCmd_SetEnv(name, value string) error {
	return os.Setenv(name, value)
}

// SystemCmd_GetEnv 获取环境变量。
func SystemCmd_GetEnv(name string) string {
	return os.Getenv(name)
}

// SystemCmd_UnsetEnv 删除环境变量。
func SystemCmd_UnsetEnv(name string) error {
	return os.Unsetenv(name)
}