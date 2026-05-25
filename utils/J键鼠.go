//go:build windows

// 键盘鼠标操作模块
// 提供键盘模拟输入、鼠标模拟操作、按键状态检测等功能。
// 基于 Windows user32.dll API 实现，支持 SendInput / keybd_event / mouse_event / GetAsyncKeyState 等底层调用。
// 包含完整的虚拟键码（VK_*）常量定义，方便直接引用。
package utils

import (
	"syscall"
	"unsafe"
)

var (
	// DLL 引用
	user32J键鼠   = syscall.NewLazyDLL("user32.dll")
	kernel32J键鼠 = syscall.NewLazyDLL("kernel32.dll")

	// user32 API 函数指针
	procSendInput           = user32J键鼠.NewProc("SendInput")             // 发送输入事件（键盘/鼠标/硬件）
	procKeybdEvent          = user32J键鼠.NewProc("keybd_event")           // 模拟键盘事件
	procMouseEvent          = user32J键鼠.NewProc("mouse_event")           // 模拟鼠标事件
	procGetAsyncKeyState    = user32J键鼠.NewProc("GetAsyncKeyState")      // 异步获取按键状态（不阻塞）
	procGetKeyState         = user32J键鼠.NewProc("GetKeyState")           // 获取指定键的状态
	procMapVirtualKeyW      = user32J键鼠.NewProc("MapVirtualKeyW")        // 虚拟键码转扫描码
	procSetCursorPos        = user32J键鼠.NewProc("SetCursorPos")          // 设置光标位置
	procGetCursorPos        = user32J键鼠.NewProc("GetCursorPos")          // 获取光标位置
	procGetSystemMetrics    = user32J键鼠.NewProc("GetSystemMetrics")      // 获取系统度量（分辨率等）
	procBlockInput          = user32J键鼠.NewProc("BlockInput")            // 锁定/解锁输入设备

	// kernel32 API
	procGetModuleHandleW = kernel32J键鼠.NewProc("GetModuleHandleW") // 获取模块句柄
)

// ===================== WIN32 常量定义 =====================

const (
	// SendInput 事件类型
	INPUT_MOUSE    = 0 // 鼠标输入事件
	INPUT_KEYBOARD = 1 // 键盘输入事件
	INPUT_HARDWARE = 2 // 硬件输入事件

	// mouse_event 标志位
	MOUSEEVENTF_MOVE       = 0x0001 // 鼠标移动
	MOUSEEVENTF_LEFTDOWN   = 0x0002 // 左键按下
	MOUSEEVENTF_LEFTUP     = 0x0004 // 左键弹起
	MOUSEEVENTF_RIGHTDOWN  = 0x0008 // 右键按下
	MOUSEEVENTF_RIGHTUP    = 0x0010 // 右键弹起
	MOUSEEVENTF_MIDDLEDOWN = 0x0020 // 中键按下
	MOUSEEVENTF_MIDDLEUP   = 0x0040 // 中键弹起
	MOUSEEVENTF_WHEEL      = 0x0800 // 滚轮滚动
	MOUSEEVENTF_ABSOLUTE   = 0x8000 // 绝对坐标

	// keybd_event 标志位
	KEYEVENTF_EXTENDEDKEY = 0x0001 // 扩展键（右侧 Alt/Ctrl/方向键等）
	KEYEVENTF_KEYUP       = 0x0002 // 按键弹起

	// MapVirtualKey 映射类型
	MAPVK_VK_TO_VSC = 0 // 虚拟键码 → 硬件扫描码

	// GetSystemMetrics 索引
	SM_CXSCREEN = 0 // 主屏幕宽度（像素）
	SM_CYSCREEN = 1 // 主屏幕高度（像素）
)

// ===================== WIN32 结构体定义 =====================

// MOUSEINPUT 鼠标输入事件结构体
type MOUSEINPUT struct {
	Dx          int32  // 绝对/相对 X 坐标
	Dy          int32  // 绝对/相对 Y 坐标
	MouseData   uint32 // 鼠标滚轮数据
	DwFlags     uint32 // 事件标志位（MOUSEEVENTF_*）
	Time        uint32 // 事件时间戳
	DwExtraInfo uintptr // 附加信息
}

// KEYBDINPUT 键盘输入事件结构体
type KEYBDINPUT struct {
	WVk         uint16 // 虚拟键码
	WScan       uint16 // 硬件扫描码
	DwFlags     uint32 // 事件标志位（KEYEVENTF_*）
	Time        uint32 // 事件时间戳
	DwExtraInfo uintptr // 附加信息
}

// HARDWAREINPUT 硬件输入事件结构体
type HARDWAREINPUT struct {
	UMsg    uint32 // 硬件消息
	WParamL uint16 // 低字参数
	WParamH uint16 // 高字参数
}

// INPUT 输入事件联合体（SendInput 使用）
type INPUT struct {
	Type uint32       // 事件类型（INPUT_MOUSE / INPUT_KEYBOARD / INPUT_HARDWARE）
	Ki   KEYBDINPUT   // 键盘输入数据
	Mi   MOUSEINPUT   // 鼠标输入数据
	Hi   HARDWAREINPUT // 硬件输入数据
}

// POINT 坐标结构体
type POINT struct {
	X int32 // X 坐标
	Y int32 // Y 坐标
}

// ===================== 键盘操作 =====================

// J键鼠_按键 模拟按下或弹起一个键。
// 参数 虚拟键码：Win32 虚拟键码常量（VK_*）
// 参数 按下：true 为按下，false 为弹起
func J键鼠_按键(虚拟键码 int, 按下 bool) {
	var dwFlags uint32
	if !按下 {
		dwFlags = KEYEVENTF_KEYUP
	}
	scanCode := MapVirtualKey(uint32(虚拟键码), MAPVK_VK_TO_VSC)
	procKeybdEvent.Call(
		uintptr(虚拟键码),
		uintptr(scanCode),
		uintptr(dwFlags),
		0,
	)
}

// J键鼠_按键组合 模拟组合键（如 Ctrl+C：VK_CONTROL + VK_C）。
// 先按下两个键，再按相反顺序弹起。
// 参数 虚拟键码1：修饰键（Ctrl/Alt/Shift）
// 参数 虚拟键码2：目标键
func J键鼠_按键组合(虚拟键码1, 虚拟键码2 int) {
	J键鼠_按键(虚拟键码1, true)
	J键鼠_按键(虚拟键码2, true)
	J键鼠_按键(虚拟键码2, false)
	J键鼠_按键(虚拟键码1, false)
}

// J键鼠_模拟按键 模拟按键的完整按下 + 弹起过程。
// 参数 虚拟键码：Win32 虚拟键码常量
func J键鼠_模拟按键(虚拟键码 int) {
	J键鼠_按键(虚拟键码, true)
	J键鼠_按键(虚拟键码, false)
}

// J键鼠_取按键状态 异步获取按键当前是否被按下（不依赖消息队列）。
// 参数 虚拟键码：Win32 虚拟键码常量
// 返回 int16：最高位（0x8000）为 1 表示当前按下，最低位（0x0001）表示自上次查询后被按过
func J键鼠_取按键状态(虚拟键码 int) int16 {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(虚拟键码))
	return int16(ret)
}

// J键鼠_是否按下 判断指定键是否处于按下状态。
// 参数 虚拟键码：Win32 虚拟键码常量
// 返回 bool：true 表示当前按下
func J键鼠_是否按下(虚拟键码 int) bool {
	return J键鼠_取按键状态(虚拟键码)&0x8000 != 0
}

// J键鼠_取键状态 从消息队列获取键的状态（包括切换状态和按下状态）。
// 参数 虚拟键码：Win32 虚拟键码常量
// 返回 int16：键状态值
func J键鼠_取键状态(虚拟键码 int) int16 {
	ret, _, _ := procGetKeyState.Call(uintptr(虚拟键码))
	return int16(ret)
}

// J键鼠_模拟文本输入 逐字符模拟键盘输入中文/英文文本。
// 自动处理 Shift 大小写切换，使用 VkKeyScan 获取每个字符的虚拟键码。
// 参数 text：要输入的文本字符串
func J键鼠_模拟文本输入(text string) {
	for _, c := range text {
		vk := uintptr(VkKeyScan(c))
		shift := (vk >> 8) & 1
		if shift != 0 {
			J键鼠_按键(VK_SHIFT, true)
		}
		J键鼠_模拟按键(int(vk & 0xFF))
		if shift != 0 {
			J键鼠_按键(VK_SHIFT, false)
		}
	}
}

// ===================== 鼠标操作 =====================

// J键鼠_移动鼠标 将鼠标光标移动到指定屏幕坐标。
// 参数 x：屏幕 X 坐标（像素）
// 参数 y：屏幕 Y 坐标（像素）
func J键鼠_移动鼠标(x, y int) {
	procSetCursorPos.Call(uintptr(x), uintptr(y))
}

// J键鼠_取鼠标位置 获取当前鼠标光标的屏幕坐标。
// 返回 (int, int)：X 和 Y 坐标（像素）
func J键鼠_取鼠标位置() (int, int) {
	var pt POINT
	procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	return int(pt.X), int(pt.Y)
}

// J键鼠_鼠标左键单击 在当前鼠标位置模拟左键单击（按下 + 弹起）。
func J键鼠_鼠标左键单击() {
	x, y := J键鼠_取鼠标位置()
	procMouseEvent.Call(
		uintptr(MOUSEEVENTF_LEFTDOWN|MOUSEEVENTF_LEFTUP),
		uintptr(x), uintptr(y), 0, 0,
	)
}

// J键鼠_鼠标左键按下 在当前鼠标位置模拟左键按下（不弹起）。
func J键鼠_鼠标左键按下() {
	x, y := J键鼠_取鼠标位置()
	procMouseEvent.Call(
		uintptr(MOUSEEVENTF_LEFTDOWN),
		uintptr(x), uintptr(y), 0, 0,
	)
}

// J键鼠_鼠标左键弹起 在当前鼠标位置模拟左键弹起。
func J键鼠_鼠标左键弹起() {
	x, y := J键鼠_取鼠标位置()
	procMouseEvent.Call(
		uintptr(MOUSEEVENTF_LEFTUP),
		uintptr(x), uintptr(y), 0, 0,
	)
}

// J键鼠_鼠标右键单击 在当前鼠标位置模拟右键单击（按下 + 弹起）。
func J键鼠_鼠标右键单击() {
	x, y := J键鼠_取鼠标位置()
	procMouseEvent.Call(
		uintptr(MOUSEEVENTF_RIGHTDOWN|MOUSEEVENTF_RIGHTUP),
		uintptr(x), uintptr(y), 0, 0,
	)
}

// J键鼠_鼠标中键单击 在当前鼠标位置模拟中键单击（按下 + 弹起）。
func J键鼠_鼠标中键单击() {
	x, y := J键鼠_取鼠标位置()
	procMouseEvent.Call(
		uintptr(MOUSEEVENTF_MIDDLEDOWN|MOUSEEVENTF_MIDDLEUP),
		uintptr(x), uintptr(y), 0, 0,
	)
}

// J键鼠_鼠标滚轮 在当前鼠标位置滚动滚轮。
// 参数 delta：正值向上滚动，负值向下滚动（1 单位 = 120 像素滚动距离）
func J键鼠_鼠标滚轮(delta int) {
	x, y := J键鼠_取鼠标位置()
	// 每个单位对应 WHEEL_DELTA(120)
	mouseData := int32(delta) * 120
	procMouseEvent.Call(
		uintptr(MOUSEEVENTF_WHEEL),
		uintptr(x), uintptr(y),
		uintptr(mouseData),
		0,
	)
}

// J键鼠_取屏幕宽度 获取主显示器屏幕宽度。
// 返回 int：屏幕宽度（像素）
func J键鼠_取屏幕宽度() int {
	ret, _, _ := procGetSystemMetrics.Call(uintptr(SM_CXSCREEN))
	return int(ret)
}

// J键鼠_取屏幕高度 获取主显示器屏幕高度。
// 返回 int：屏幕高度（像素）
func J键鼠_取屏幕高度() int {
	ret, _, _ := procGetSystemMetrics.Call(uintptr(SM_CYSCREEN))
	return int(ret)
}

// J键鼠_锁定输入 锁定/解锁所有键盘鼠标输入。
// 参数 锁定：true 锁定输入，false 解锁恢复输入
func J键鼠_锁定输入(锁定 bool) {
	if 锁定 {
		procBlockInput.Call(1)
	} else {
		procBlockInput.Call(0)
	}
}

// ===================== 辅助函数 =====================

// MapVirtualKey 将虚拟键码映射为硬件扫描码。
// 参数 uCode：虚拟键码
// 参数 uMapType：映射类型（MAPVK_VK_TO_VSC 等）
// 返回 uint32：扫描码值
func MapVirtualKey(uCode, uMapType uint32) uint32 {
	ret, _, _ := procMapVirtualKeyW.Call(uintptr(uCode), uintptr(uMapType))
	return uint32(ret)
}

// VkKeyScan 将 Unicode 字符转换为虚拟键码和 Shift 状态。
// 参数 ch：Unicode 字符
// 返回 uint16：低字节为虚拟键码，高字节为 Shift 状态（1=需要 Shift）
func VkKeyScan(ch rune) uint16 {
	user32 := syscall.NewLazyDLL("user32.dll")
	procVkKeyScanW := user32.NewProc("VkKeyScanW")
	ret, _, _ := procVkKeyScanW.Call(uintptr(ch))
	return uint16(ret)
}

// ===================== 虚拟键码常量（完整定义） =====================

const (
	// 鼠标键
	VK_LBUTTON = 0x01 // 鼠标左键
	VK_RBUTTON = 0x02 // 鼠标右键
	VK_MBUTTON = 0x04 // 鼠标中键

	// 控制键
	VK_CANCEL  = 0x03 // Ctrl+Break
	VK_BACK    = 0x08 // Backspace（退格键）
	VK_TAB     = 0x09 // Tab
	VK_CLEAR   = 0x0C // Clear（数字键盘 5 当 NumLock 关闭时）
	VK_RETURN  = 0x0D // Enter（回车）
	VK_SHIFT   = 0x10 // Shift
	VK_CONTROL = 0x11 // Ctrl
	VK_MENU    = 0x12 // Alt
	VK_PAUSE   = 0x13 // Pause/Break
	VK_CAPITAL = 0x14 // Caps Lock
	VK_ESCAPE  = 0x1B // Esc
	VK_SPACE   = 0x20 // 空格

	// 导航键
	VK_PRIOR  = 0x21 // Page Up
	VK_NEXT   = 0x22 // Page Down
	VK_END    = 0x23 // End
	VK_HOME   = 0x24 // Home
	VK_LEFT   = 0x25 // 左方向键
	VK_UP     = 0x26 // 上方向键
	VK_RIGHT  = 0x27 // 右方向键
	VK_DOWN   = 0x28 // 下方向键

	// 其他键
	VK_SELECT   = 0x29 // Select
	VK_PRINT    = 0x2A // Print
	VK_EXECUTE  = 0x2B // Execute
	VK_SNAPSHOT = 0x2C // Print Screen
	VK_INSERT   = 0x2D // Insert
	VK_DELETE   = 0x2E // Delete
	VK_HELP     = 0x2F // Help

	// 数字键（主键盘顶部）
	VK_0 = 0x30 // 0
	VK_1 = 0x31 // 1
	VK_2 = 0x32 // 2
	VK_3 = 0x33 // 3
	VK_4 = 0x34 // 4
	VK_5 = 0x35 // 5
	VK_6 = 0x36 // 6
	VK_7 = 0x37 // 7
	VK_8 = 0x38 // 8
	VK_9 = 0x39 // 9

	// 字母键
	VK_A = 0x41 // A
	VK_B = 0x42 // B
	VK_C = 0x43 // C
	VK_D = 0x44 // D
	VK_E = 0x45 // E
	VK_F = 0x46 // F
	VK_G = 0x47 // G
	VK_H = 0x48 // H
	VK_I = 0x49 // I
	VK_J = 0x4A // J
	VK_K = 0x4B // K
	VK_L = 0x4C // L
	VK_M = 0x4D // M
	VK_N = 0x4E // N
	VK_O = 0x4F // O
	VK_P = 0x50 // P
	VK_Q = 0x51 // Q
	VK_R = 0x52 // R
	VK_S = 0x53 // S
	VK_T = 0x54 // T
	VK_U = 0x55 // U
	VK_V = 0x56 // V
	VK_W = 0x57 // W
	VK_X = 0x58 // X
	VK_Y = 0x59 // Y
	VK_Z = 0x5A // Z

	// Windows 键
	VK_LWIN = 0x5B // 左侧 Win 键
	VK_RWIN = 0x5C // 右侧 Win 键
	VK_APPS = 0x5D // 菜单键（Application）

	// 数字小键盘
	VK_NUMPAD0  = 0x60 // 小键盘 0
	VK_NUMPAD1  = 0x61 // 小键盘 1
	VK_NUMPAD2  = 0x62 // 小键盘 2
	VK_NUMPAD3  = 0x63 // 小键盘 3
	VK_NUMPAD4  = 0x64 // 小键盘 4
	VK_NUMPAD5  = 0x65 // 小键盘 5
	VK_NUMPAD6  = 0x66 // 小键盘 6
	VK_NUMPAD7  = 0x67 // 小键盘 7
	VK_NUMPAD8  = 0x68 // 小键盘 8
	VK_NUMPAD9  = 0x69 // 小键盘 9
	VK_MULTIPLY = 0x6A // 小键盘 *
	VK_ADD      = 0x6B // 小键盘 +
	VK_SEPARATOR = 0x6C // 分隔符
	VK_SUBTRACT = 0x6D // 小键盘 -
	VK_DECIMAL  = 0x6E // 小键盘 .
	VK_DIVIDE   = 0x6F // 小键盘 /

	// F 功能键
	VK_F1  = 0x70 // F1
	VK_F2  = 0x71 // F2
	VK_F3  = 0x72 // F3
	VK_F4  = 0x73 // F4
	VK_F5  = 0x74 // F5
	VK_F6  = 0x75 // F6
	VK_F7  = 0x76 // F7
	VK_F8  = 0x77 // F8
	VK_F9  = 0x78 // F9
	VK_F10 = 0x79 // F10
	VK_F11 = 0x7A // F11
	VK_F12 = 0x7B // F12

	// 锁定键
	VK_NUMLOCK = 0x90 // Num Lock
	VK_SCROLL  = 0x91 // Scroll Lock

	// 左右修饰键
	VK_LSHIFT   = 0xA0 // 左 Shift
	VK_RSHIFT   = 0xA1 // 右 Shift
	VK_LCONTROL = 0xA2 // 左 Ctrl
	VK_RCONTROL = 0xA3 // 右 Ctrl
	VK_LMENU    = 0xA4 // 左 Alt
	VK_RMENU    = 0xA5 // 右 Alt
)