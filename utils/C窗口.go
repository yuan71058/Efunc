//go:build windows

package utils

import (
	"syscall"
	"unsafe"
)

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	procFindWindowW     = user32.NewProc("FindWindowW")
	procFindWindowExW   = user32.NewProc("FindWindowExW")
	procGetWindowTextW  = user32.NewProc("GetWindowTextW")
	procGetWindowTextLengthW = user32.NewProc("GetWindowTextLengthW")
	procSetWindowTextW  = user32.NewProc("SetWindowTextW")
	procGetClassNameW   = user32.NewProc("GetClassNameW")
	procGetWindowRect   = user32.NewProc("GetWindowRect")
	procMoveWindow      = user32.NewProc("MoveWindow")
	procShowWindow      = user32.NewProc("ShowWindow")
	procSendMessageW    = user32.NewProc("SendMessageW")
	procPostMessageW    = user32.NewProc("PostMessageW")
	procGetForegroundWindow = user32.NewProc("GetForegroundWindow")
	procSetForegroundWindow = user32.NewProc("SetForegroundWindow")
	procEnumWindows     = user32.NewProc("EnumWindows")
	procIsWindowVisible = user32.NewProc("IsWindowVisible")
	procIsWindow        = user32.NewProc("IsWindow")
	procGetWindowThreadProcessId = user32.NewProc("GetWindowThreadProcessId")
	procGetParent       = user32.NewProc("GetParent")
	procGetDesktopWindow = user32.NewProc("GetDesktopWindow")
	procCloseWindow     = user32.NewProc("CloseWindow")
	procDestroyWindow   = user32.NewProc("DestroyWindow")
	procGetWindow       = user32.NewProc("GetWindow")
	procEnableWindow    = user32.NewProc("EnableWindow")
)

type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

const (
	GW_HWNDNEXT    = 2
	GW_HWNDPREV    = 3
	GW_OWNER       = 4
	SW_HIDE        = 0
	SW_SHOW        = 5
	SW_MINIMIZE    = 6
	SW_MAXIMIZE    = 3
	SW_RESTORE     = 9
	WM_CLOSE       = 0x0010
	WM_SETTEXT     = 0x000C
	WM_GETTEXT     = 0x000D
	WM_CLICK       = 0x00F5
	BM_CLICK       = 0x00F5
)

// C窗口_查找 按类名和窗口标题查找顶层窗口。
// 类名或标题可为空字符串（表示不限制该条件）。
//
// 参数:
//   - 类名: 窗口类名，如 "Notepad"；空字符串表示忽略类名
//   - 标题: 窗口标题；空字符串表示忽略标题
//
// 返回:
//   - syscall.Handle: 窗口句柄；未找到返回 0
func C窗口_查找(类名 string, 标题 string) syscall.Handle {
	var 类名Ptr, 标题Ptr *uint16
	if 类名 != "" {
		类名Ptr, _ = syscall.UTF16PtrFromString(类名)
	}
	if 标题 != "" {
		标题Ptr, _ = syscall.UTF16PtrFromString(标题)
	}
	ret, _, _ := procFindWindowW.Call(uintptr(unsafe.Pointer(类名Ptr)), uintptr(unsafe.Pointer(标题Ptr)))
	return syscall.Handle(ret)
}

// C窗口_查找子窗口 在父窗口中查找子窗口。
// 类名或标题可为空字符串（表示不限制该条件）。
//
// 参数:
//   - 父窗口: 父窗口句柄
//   - 子窗口后: 从该子窗口之后开始查找，0 表示从第一个开始
//   - 类名: 子窗口类名
//   - 标题: 子窗口标题
//
// 返回:
//   - syscall.Handle: 子窗口句柄；未找到返回 0
func C窗口_查找子窗口(父窗口 syscall.Handle, 子窗口后 syscall.Handle, 类名 string, 标题 string) syscall.Handle {
	var 类名Ptr, 标题Ptr *uint16
	if 类名 != "" {
		类名Ptr, _ = syscall.UTF16PtrFromString(类名)
	}
	if 标题 != "" {
		标题Ptr, _ = syscall.UTF16PtrFromString(标题)
	}
	ret, _, _ := procFindWindowExW.Call(
		uintptr(父窗口), uintptr(子窗口后),
		uintptr(unsafe.Pointer(类名Ptr)), uintptr(unsafe.Pointer(标题Ptr)))
	return syscall.Handle(ret)
}

// C窗口_取标题 获取窗口的标题文本。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//
// 返回:
//   - string: 窗口标题文本
func C窗口_取标题(窗口句柄 syscall.Handle) string {
	长度, _, _ := procGetWindowTextLengthW.Call(uintptr(窗口句柄))
	if 长度 == 0 {
		return ""
	}
	buf := make([]uint16, 长度+1)
	procGetWindowTextW.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(&buf[0])), uintptr(长度+1))
	return syscall.UTF16ToString(buf)
}

// C窗口_置标题 设置窗口的标题文本。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//   - 标题: 要设置的新标题
//
// 返回:
//   - bool: 设置成功返回 true
func C窗口_置标题(窗口句柄 syscall.Handle, 标题 string) bool {
	标题Ptr, _ := syscall.UTF16PtrFromString(标题)
	ret, _, _ := procSetWindowTextW.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(标题Ptr)))
	return ret != 0
}

// C窗口_取类名 获取窗口的类名。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//
// 返回:
//   - string: 窗口类名
func C窗口_取类名(窗口句柄 syscall.Handle) string {
	buf := make([]uint16, 256)
	ret, _, _ := procGetClassNameW.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(&buf[0])), 256)
	if ret == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// C窗口_取矩形 获取窗口在屏幕上的位置和大小。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//
// 返回:
//   - RECT: 窗口矩形区域（Left, Top, Right, Bottom）
//   - bool: 获取成功返回 true
func C窗口_取矩形(窗口句柄 syscall.Handle) (RECT, bool) {
	var rect RECT
	ret, _, _ := procGetWindowRect.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(&rect)))
	return rect, ret != 0
}

// C窗口_移动 移动并调整窗口的大小和位置。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//   - 左边: 窗口左上角 X 坐标
//   - 顶边: 窗口左上角 Y 坐标
//   - 宽度: 窗口宽度
//   - 高度: 窗口高度
//   - 重绘: 是否重绘窗口
//
// 返回:
//   - bool: 移动成功返回 true
func C窗口_移动(窗口句柄 syscall.Handle, 左边 int32, 顶边 int32, 宽度 int32, 高度 int32, 重绘 bool) bool {
	var 重绘值 uintptr
	if 重绘 {
		重绘值 = 1
	}
	ret, _, _ := procMoveWindow.Call(uintptr(窗口句柄), uintptr(左边), uintptr(顶边), uintptr(宽度), uintptr(高度), 重绘值)
	return ret != 0
}

// C窗口_显示 控制窗口的显示状态。
// 常用命令：SW_HIDE(隐藏)、SW_SHOW(显示)、SW_MINIMIZE(最小化)、SW_MAXIMIZE(最大化)、SW_RESTORE(还原)。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//   - 命令: 显示命令，如 SW_SHOW、SW_HIDE
//
// 返回:
//   - bool: 操作成功返回 true
func C窗口_显示(窗口句柄 syscall.Handle, 命令 int) bool {
	ret, _, _ := procShowWindow.Call(uintptr(窗口句柄), uintptr(命令))
	return ret != 0
}

// C窗口_发送消息 向窗口发送同步消息，等待消息处理完毕后返回。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//   - 消息: 消息标识符，如 WM_CLOSE、WM_SETTEXT
//   - 参数1: 消息的 wParam 参数
//   - 参数2: 消息的 lParam 参数
//
// 返回:
//   - uintptr: 消息处理结果
func C窗口_发送消息(窗口句柄 syscall.Handle, 消息 uint32, 参数1 uintptr, 参数2 uintptr) uintptr {
	ret, _, _ := procSendMessageW.Call(uintptr(窗口句柄), uintptr(消息), 参数1, 参数2)
	return ret
}

// C窗口_投递消息 向窗口投递异步消息，不等待处理结果立即返回。
// 适用于不需要返回值的消息，如 WM_CLOSE。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//   - 消息: 消息标识符
//   - 参数1: 消息的 wParam 参数
//   - 参数2: 消息的 lParam 参数
//
// 返回:
//   - bool: 投递成功返回 true
func C窗口_投递消息(窗口句柄 syscall.Handle, 消息 uint32, 参数1 uintptr, 参数2 uintptr) bool {
	ret, _, _ := procPostMessageW.Call(uintptr(窗口句柄), uintptr(消息), 参数1, 参数2)
	return ret != 0
}

// C窗口_关闭 发送 WM_CLOSE 消息关闭窗口。
// 窗口可以拦截此消息拒绝关闭。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//
// 返回:
//   - bool: 发送成功返回 true
func C窗口_关闭(窗口句柄 syscall.Handle) bool {
	return C窗口_投递消息(窗口句柄, WM_CLOSE, 0, 0)
}

// C窗口_点击按钮 向按钮控件发送点击消息。
//
// 参数:
//   - 按钮句柄: 按钮控件的窗口句柄
//
// 返回:
//   - uintptr: 消息处理结果
func C窗口_点击按钮(按钮句柄 syscall.Handle) uintptr {
	return C窗口_发送消息(按钮句柄, BM_CLICK, 0, 0)
}

// C窗口_取前台窗口 获取当前前台（具有输入焦点）窗口的句柄。
//
// 返回:
//   - syscall.Handle: 前台窗口句柄
func C窗口_取前台窗口() syscall.Handle {
	ret, _, _ := procGetForegroundWindow.Call()
	return syscall.Handle(ret)
}

// C窗口_置前台窗口 将指定窗口设为前台窗口。
//
// 参数:
//   - 窗口句柄: 要设为前台的窗口句柄
//
// 返回:
//   - bool: 设置成功返回 true
func C窗口_置前台窗口(窗口句柄 syscall.Handle) bool {
	ret, _, _ := procSetForegroundWindow.Call(uintptr(窗口句柄))
	return ret != 0
}

// C窗口_是否可见 检查窗口是否可见。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//
// 返回:
//   - bool: 窗口可见返回 true
func C窗口_是否可见(窗口句柄 syscall.Handle) bool {
	ret, _, _ := procIsWindowVisible.Call(uintptr(窗口句柄))
	return ret != 0
}

// C窗口_是否有效 检查窗口句柄是否有效（窗口是否存在）。
//
// 参数:
//   - 窗口句柄: 要检查的窗口句柄
//
// 返回:
//   - bool: 窗口存在返回 true
func C窗口_是否有效(窗口句柄 syscall.Handle) bool {
	ret, _, _ := procIsWindow.Call(uintptr(窗口句柄))
	return ret != 0
}

// C窗口_取进程ID 获取窗口所属的进程 ID。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//
// 返回:
//   - uint32: 进程 ID
//   - uint32: 线程 ID
func C窗口_取进程ID(窗口句柄 syscall.Handle) (uint32, uint32) {
	var 进程ID uint32
	线程ID, _, _ := procGetWindowThreadProcessId.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(&进程ID)))
	return 进程ID, uint32(线程ID)
}

// C窗口_取父窗口 获取指定窗口的父窗口句柄。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//
// 返回:
//   - syscall.Handle: 父窗口句柄；无父窗口返回 0
func C窗口_取父窗口(窗口句柄 syscall.Handle) syscall.Handle {
	ret, _, _ := procGetParent.Call(uintptr(窗口句柄))
	return syscall.Handle(ret)
}

// C窗口_取桌面窗口 获取桌面窗口的句柄。
//
// 返回:
//   - syscall.Handle: 桌面窗口句柄
func C窗口_取桌面窗口() syscall.Handle {
	ret, _, _ := procGetDesktopWindow.Call()
	return syscall.Handle(ret)
}

// C窗口_启用 启用或禁用窗口的鼠标和键盘输入。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//   - 启用: true 启用窗口，false 禁用窗口
//
// 返回:
//   - bool: 窗口之前的状态（之前被禁用返回 true）
func C窗口_启用(窗口句柄 syscall.Handle, 启用 bool) bool {
	var 启用值 uintptr
	if 启用 {
		启用值 = 1
	}
	ret, _, _ := procEnableWindow.Call(uintptr(窗口句柄), 启用值)
	return ret != 0
}

// C窗口_取下一个 获取与指定窗口具有相同类型的下一个窗口（Z 顺序）。
//
// 参数:
//   - 窗口句柄: 起始窗口句柄
//
// 返回:
//   - syscall.Handle: 下一个窗口句柄；没有返回 0
func C窗口_取下一个(窗口句柄 syscall.Handle) syscall.Handle {
	ret, _, _ := procGetWindow.Call(uintptr(窗口句柄), GW_HWNDNEXT)
	return syscall.Handle(ret)
}

// C窗口_取所有子窗口 获取指定窗口的所有子窗口句柄。
//
// 参数:
//   - 父窗口: 父窗口句柄
//
// 返回:
//   - []syscall.Handle: 子窗口句柄列表
func C窗口_取所有子窗口(父窗口 syscall.Handle) []syscall.Handle {
	var 子窗口列表 []syscall.Handle
	var 子窗口 syscall.Handle
	for {
		子窗口 = C窗口_查找子窗口(父窗口, 子窗口, "", "")
		if 子窗口 == 0 {
			break
		}
		子窗口列表 = append(子窗口列表, 子窗口)
	}
	return 子窗口列表
}

// C窗口_置文本 向窗口设置文本内容（如编辑框）。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//   - 文本: 要设置的文本
//
// 返回:
//   - uintptr: 消息处理结果
func C窗口_置文本(窗口句柄 syscall.Handle, 文本 string) uintptr {
	文本Ptr, _ := syscall.UTF16PtrFromString(文本)
	return C窗口_发送消息(窗口句柄, WM_SETTEXT, 0, uintptr(unsafe.Pointer(文本Ptr)))
}

// C窗口_取文本 从窗口获取文本内容（如编辑框）。
//
// 参数:
//   - 窗口句柄: 目标窗口句柄
//
// 返回:
//   - string: 窗口文本内容
func C窗口_取文本(窗口句柄 syscall.Handle) string {
	长度 := C窗口_发送消息(窗口句柄, WM_GETTEXT, 0, 0)
	if 长度 == 0 {
		return ""
	}
	buf := make([]uint16, 长度+1)
	C窗口_发送消息(窗口句柄, WM_GETTEXT, uintptr(长度+1), uintptr(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf)
}
