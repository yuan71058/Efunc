//go:build windows

// Windows 窗口管理模块
// 提供窗口查找、子窗口枚举、标题/类名获取与设置、窗口位置/大小控制、
// 消息发送/投递、前台窗口管理等功能。
// 基于 Windows user32.dll API 实现。
package utils

import (
	"syscall"
	"unsafe"
)

var (
	user32Win             = syscall.NewLazyDLL("user32.dll")
	procFindWindowW       = user32Win.NewProc("FindWindowW")
	procFindWindowExW     = user32Win.NewProc("FindWindowExW")
	procGetWindowTextW    = user32Win.NewProc("GetWindowTextW")
	procGetWindowTextLen  = user32Win.NewProc("GetWindowTextLengthW")
	procSetWindowTextW    = user32Win.NewProc("SetWindowTextW")
	procGetClassNameW     = user32Win.NewProc("GetClassNameW")
	procGetWindowRect     = user32Win.NewProc("GetWindowRect")
	procMoveWindow        = user32Win.NewProc("MoveWindow")
	procShowWindow        = user32Win.NewProc("ShowWindow")
	procSendMessageW      = user32Win.NewProc("SendMessageW")
	procPostMessageW      = user32Win.NewProc("PostMessageW")
	procGetForegroundWnd  = user32Win.NewProc("GetForegroundWindow")
	procSetForegroundWnd  = user32Win.NewProc("SetForegroundWindow")
	procEnumWindows       = user32Win.NewProc("EnumWindows")
	procIsWindowVisible   = user32Win.NewProc("IsWindowVisible")
	procIsWindow          = user32Win.NewProc("IsWindow")
	procGetWindowTID      = user32Win.NewProc("GetWindowThreadProcessId")
	procGetParent         = user32Win.NewProc("GetParent")
	procGetDesktopWindow  = user32Win.NewProc("GetDesktopWindow")
	procCloseWindow       = user32Win.NewProc("CloseWindow")
	procDestroyWindow     = user32Win.NewProc("DestroyWindow")
	procGetWindow         = user32Win.NewProc("GetWindow")
	procEnableWindow      = user32Win.NewProc("EnableWindow")
)

type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

const (
	GW_HWNDNEXT = 2
	GW_HWNDPREV = 3
	GW_OWNER    = 4
	SW_HIDE     = 0
	SW_SHOW     = 5
	SW_MINIMIZE = 6
	SW_MAXIMIZE = 3
	SW_RESTORE  = 9
	WM_CLOSE    = 0x0010
	WM_SETTEXT  = 0x000C
	WM_GETTEXT  = 0x000D
	WM_CLICK    = 0x00F5
	BM_CLICK    = 0x00F5
)

// Window_Find 按类名和窗口标题查找顶层窗口。
// 类名或标题可为空字符串（表示不限制该条件）。
//
// 参数:
//   - className: 窗口类名，如 "Notepad"；空字符串表示忽略类名
//   - title: 窗口标题；空字符串表示忽略标题
//
// 返回:
//   - syscall.Handle: 窗口句柄；未找到返回 0
func Window_Find(className string, title string) syscall.Handle {
	var classPtr, titlePtr *uint16
	if className != "" {
		classPtr, _ = syscall.UTF16PtrFromString(className)
	}
	if title != "" {
		titlePtr, _ = syscall.UTF16PtrFromString(title)
	}
	ret, _, _ := procFindWindowW.Call(uintptr(unsafe.Pointer(classPtr)), uintptr(unsafe.Pointer(titlePtr)))
	return syscall.Handle(ret)
}

// Window_FindChild 在父窗口中查找子窗口。
// 类名或标题可为空字符串（表示不限制该条件）。
//
// 参数:
//   - parent: 父窗口句柄
//   - afterChild: 从该子窗口之后开始查找，0 表示从第一个开始
//   - className: 子窗口类名
//   - title: 子窗口标题
//
// 返回:
//   - syscall.Handle: 子窗口句柄；未找到返回 0
func Window_FindChild(parent syscall.Handle, afterChild syscall.Handle, className string, title string) syscall.Handle {
	var classPtr, titlePtr *uint16
	if className != "" {
		classPtr, _ = syscall.UTF16PtrFromString(className)
	}
	if title != "" {
		titlePtr, _ = syscall.UTF16PtrFromString(title)
	}
	ret, _, _ := procFindWindowExW.Call(
		uintptr(parent), uintptr(afterChild),
		uintptr(unsafe.Pointer(classPtr)), uintptr(unsafe.Pointer(titlePtr)))
	return syscall.Handle(ret)
}

// Window_GetTitle 获取窗口的标题文本。
//
// 参数:
//   - hWnd: 目标窗口句柄
//
// 返回:
//   - string: 窗口标题文本
func Window_GetTitle(hWnd syscall.Handle) string {
	length, _, _ := procGetWindowTextLen.Call(uintptr(hWnd))
	if length == 0 {
		return ""
	}
	buf := make([]uint16, length+1)
	procGetWindowTextW.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(length+1))
	return syscall.UTF16ToString(buf)
}

// Window_SetTitle 设置窗口的标题文本。
//
// 参数:
//   - hWnd: 目标窗口句柄
//   - title: 要设置的新标题
//
// 返回:
//   - bool: 设置成功返回 true
func Window_SetTitle(hWnd syscall.Handle, title string) bool {
	titlePtr, _ := syscall.UTF16PtrFromString(title)
	ret, _, _ := procSetWindowTextW.Call(uintptr(hWnd), uintptr(unsafe.Pointer(titlePtr)))
	return ret != 0
}

// Window_GetClassName 获取窗口的类名。
//
// 参数:
//   - hWnd: 目标窗口句柄
//
// 返回:
//   - string: 窗口类名
func Window_GetClassName(hWnd syscall.Handle) string {
	buf := make([]uint16, 256)
	ret, _, _ := procGetClassNameW.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&buf[0])), 256)
	if ret == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// Window_GetRect 获取窗口在屏幕上的位置和大小。
//
// 参数:
//   - hWnd: 目标窗口句柄
//
// 返回:
//   - RECT: 窗口矩形区域（Left, Top, Right, Bottom）
//   - bool: 获取成功返回 true
func Window_GetRect(hWnd syscall.Handle) (RECT, bool) {
	var rect RECT
	ret, _, _ := procGetWindowRect.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&rect)))
	return rect, ret != 0
}

// Window_Move 移动并调整窗口的大小和位置。
//
// 参数:
//   - hWnd: 目标窗口句柄
//   - left: 窗口左上角 X 坐标
//   - top: 窗口左上角 Y 坐标
//   - width: 窗口宽度
//   - height: 窗口高度
//   - repaint: 是否重绘窗口
//
// 返回:
//   - bool: 移动成功返回 true
func Window_Move(hWnd syscall.Handle, left int32, top int32, width int32, height int32, repaint bool) bool {
	var repaintVal uintptr
	if repaint {
		repaintVal = 1
	}
	ret, _, _ := procMoveWindow.Call(uintptr(hWnd), uintptr(left), uintptr(top), uintptr(width), uintptr(height), repaintVal)
	return ret != 0
}

// Window_Show 控制窗口的显示状态。
// 常用命令：SW_HIDE(隐藏)、SW_SHOW(显示)、SW_MINIMIZE(最小化)、SW_MAXIMIZE(最大化)、SW_RESTORE(还原)。
//
// 参数:
//   - hWnd: 目标窗口句柄
//   - cmdShow: 显示命令，如 SW_SHOW、SW_HIDE
//
// 返回:
//   - bool: 操作成功返回 true
func Window_Show(hWnd syscall.Handle, cmdShow int) bool {
	ret, _, _ := procShowWindow.Call(uintptr(hWnd), uintptr(cmdShow))
	return ret != 0
}

// Window_SendMessage 向窗口发送同步消息，等待消息处理完毕后返回。
//
// 参数:
//   - hWnd: 目标窗口句柄
//   - msg: 消息标识符，如 WM_CLOSE、WM_SETTEXT
//   - wParam: 消息的 wParam 参数
//   - lParam: 消息的 lParam 参数
//
// 返回:
//   - uintptr: 消息处理结果
func Window_SendMessage(hWnd syscall.Handle, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	ret, _, _ := procSendMessageW.Call(uintptr(hWnd), uintptr(msg), wParam, lParam)
	return ret
}

// Window_PostMessage 向窗口投递异步消息，不等待处理结果立即返回。
// 适用于不需要返回值的消息，如 WM_CLOSE。
//
// 参数:
//   - hWnd: 目标窗口句柄
//   - msg: 消息标识符
//   - wParam: 消息的 wParam 参数
//   - lParam: 消息的 lParam 参数
//
// 返回:
//   - bool: 投递成功返回 true
func Window_PostMessage(hWnd syscall.Handle, msg uint32, wParam uintptr, lParam uintptr) bool {
	ret, _, _ := procPostMessageW.Call(uintptr(hWnd), uintptr(msg), wParam, lParam)
	return ret != 0
}

// Window_Close 发送 WM_CLOSE 消息关闭窗口。
// 窗口可以拦截此消息拒绝关闭。
//
// 参数:
//   - hWnd: 目标窗口句柄
//
// 返回:
//   - bool: 发送成功返回 true
func Window_Close(hWnd syscall.Handle) bool {
	return Window_PostMessage(hWnd, WM_CLOSE, 0, 0)
}

// Window_ClickButton 向按钮控件发送点击消息。
//
// 参数:
//   - btnHwnd: 按钮控件的窗口句柄
//
// 返回:
//   - uintptr: 消息处理结果
func Window_ClickButton(btnHwnd syscall.Handle) uintptr {
	return Window_SendMessage(btnHwnd, BM_CLICK, 0, 0)
}

// Window_GetForeground 获取当前前台（具有输入焦点）窗口的句柄。
//
// 返回:
//   - syscall.Handle: 前台窗口句柄
func Window_GetForeground() syscall.Handle {
	ret, _, _ := procGetForegroundWnd.Call()
	return syscall.Handle(ret)
}

// Window_SetForeground 将指定窗口设为前台窗口。
//
// 参数:
//   - hWnd: 要设为前台的窗口句柄
//
// 返回:
//   - bool: 设置成功返回 true
func Window_SetForeground(hWnd syscall.Handle) bool {
	ret, _, _ := procSetForegroundWnd.Call(uintptr(hWnd))
	return ret != 0
}

// Window_IsVisible 检查窗口是否可见。
//
// 参数:
//   - hWnd: 目标窗口句柄
//
// 返回:
//   - bool: 窗口可见返回 true
func Window_IsVisible(hWnd syscall.Handle) bool {
	ret, _, _ := procIsWindowVisible.Call(uintptr(hWnd))
	return ret != 0
}

// Window_IsValid 检查窗口句柄是否有效（窗口是否存在）。
//
// 参数:
//   - hWnd: 要检查的窗口句柄
//
// 返回:
//   - bool: 窗口存在返回 true
func Window_IsValid(hWnd syscall.Handle) bool {
	ret, _, _ := procIsWindow.Call(uintptr(hWnd))
	return ret != 0
}

// Window_GetProcessID 获取窗口所属的进程 ID。
//
// 参数:
//   - hWnd: 目标窗口句柄
//
// 返回:
//   - uint32: 进程 ID
//   - uint32: 线程 ID
func Window_GetProcessID(hWnd syscall.Handle) (uint32, uint32) {
	var processID uint32
	threadID, _, _ := procGetWindowTID.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&processID)))
	return processID, uint32(threadID)
}

// Window_GetParent 获取指定窗口的父窗口句柄。
//
// 参数:
//   - hWnd: 目标窗口句柄
//
// 返回:
//   - syscall.Handle: 父窗口句柄；无父窗口返回 0
func Window_GetParent(hWnd syscall.Handle) syscall.Handle {
	ret, _, _ := procGetParent.Call(uintptr(hWnd))
	return syscall.Handle(ret)
}

// Window_GetDesktop 获取桌面窗口的句柄。
//
// 返回:
//   - syscall.Handle: 桌面窗口句柄
func Window_GetDesktop() syscall.Handle {
	ret, _, _ := procGetDesktopWindow.Call()
	return syscall.Handle(ret)
}

// Window_Enable 启用或禁用窗口的鼠标和键盘输入。
//
// 参数:
//   - hWnd: 目标窗口句柄
//   - enable: true 启用窗口，false 禁用窗口
//
// 返回:
//   - bool: 窗口之前的状态（之前被禁用返回 true）
func Window_Enable(hWnd syscall.Handle, enable bool) bool {
	var enableVal uintptr
	if enable {
		enableVal = 1
	}
	ret, _, _ := procEnableWindow.Call(uintptr(hWnd), enableVal)
	return ret != 0
}

// Window_GetNext 获取与指定窗口具有相同类型的下一个窗口（Z 顺序）。
//
// 参数:
//   - hWnd: 起始窗口句柄
//
// 返回:
//   - syscall.Handle: 下一个窗口句柄；没有返回 0
func Window_GetNext(hWnd syscall.Handle) syscall.Handle {
	ret, _, _ := procGetWindow.Call(uintptr(hWnd), GW_HWNDNEXT)
	return syscall.Handle(ret)
}

// Window_GetAllChildren 获取指定窗口的所有子窗口句柄。
//
// 参数:
//   - parent: 父窗口句柄
//
// 返回:
//   - []syscall.Handle: 子窗口句柄列表
func Window_GetAllChildren(parent syscall.Handle) []syscall.Handle {
	var children []syscall.Handle
	var child syscall.Handle
	for {
		child = Window_FindChild(parent, child, "", "")
		if child == 0 {
			break
		}
		children = append(children, child)
	}
	return children
}

// Window_SetText 向窗口设置文本内容（如编辑框）。
//
// 参数:
//   - hWnd: 目标窗口句柄
//   - text: 要设置的文本
//
// 返回:
//   - uintptr: 消息处理结果
func Window_SetText(hWnd syscall.Handle, text string) uintptr {
	textPtr, _ := syscall.UTF16PtrFromString(text)
	return Window_SendMessage(hWnd, WM_SETTEXT, 0, uintptr(unsafe.Pointer(textPtr)))
}

// Window_GetText 从窗口获取文本内容（如编辑框）。
//
// 参数:
//   - hWnd: 目标窗口句柄
//
// 返回:
//   - string: 窗口文本内容
func Window_GetText(hWnd syscall.Handle) string {
	length := Window_SendMessage(hWnd, WM_GETTEXT, 0, 0)
	if length == 0 {
		return ""
	}
	buf := make([]uint16, length+1)
	Window_SendMessage(hWnd, WM_GETTEXT, uintptr(length+1), uintptr(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf)
}