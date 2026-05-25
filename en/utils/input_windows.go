//go:build windows

// Windows 键盘鼠标模拟模块
// 提供键盘模拟输入、鼠标模拟操作、按键状态检测等功能。
// 基于 Windows user32.dll API 实现，支持 SendInput / keybd_event / mouse_event 等底层调用。
package utils

import (
	"syscall"
	"unsafe"
)

var (
	user32Input   = syscall.NewLazyDLL("user32.dll")
	kernel32Input = syscall.NewLazyDLL("kernel32.dll")

	procSendInput               = user32Input.NewProc("SendInput")
	procKeybdEvent              = user32Input.NewProc("keybd_event")
	procMouseEvent              = user32Input.NewProc("mouse_event")
	procGetAsyncKeyState        = user32Input.NewProc("GetAsyncKeyState")
	procGetKeyState             = user32Input.NewProc("GetKeyState")
	procMapVirtualKeyW          = user32Input.NewProc("MapVirtualKeyW")
	procSetCursorPos            = user32Input.NewProc("SetCursorPos")
	procGetCursorPos            = user32Input.NewProc("GetCursorPos")
	procGetSystemMetrics        = user32Input.NewProc("GetSystemMetrics")
	procBlockInput              = user32Input.NewProc("BlockInput")

	procGetModuleHandleW = kernel32Input.NewProc("GetModuleHandleW")
)

const (
	INPUT_MOUSE    = 0
	INPUT_KEYBOARD = 1
	INPUT_HARDWARE = 2

	MOUSEEVENTF_MOVE       = 0x0001
	MOUSEEVENTF_LEFTDOWN   = 0x0002
	MOUSEEVENTF_LEFTUP     = 0x0004
	MOUSEEVENTF_RIGHTDOWN  = 0x0008
	MOUSEEVENTF_RIGHTUP    = 0x0010
	MOUSEEVENTF_MIDDLEDOWN = 0x0020
	MOUSEEVENTF_MIDDLEUP   = 0x0040
	MOUSEEVENTF_WHEEL      = 0x0800
	MOUSEEVENTF_ABSOLUTE   = 0x8000

	KEYEVENTF_EXTENDEDKEY = 0x0001
	KEYEVENTF_KEYUP       = 0x0002

	MAPVK_VK_TO_VSC = 0

	SM_CXSCREEN = 0
	SM_CYSCREEN = 1
)

type MOUSEINPUT struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

type HARDWAREINPUT struct {
	UMsg    uint32
	WParamL uint16
	WParamH uint16
}

type INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
	Mi   MOUSEINPUT
	Hi   HARDWAREINPUT
}

type POINT struct {
	X int32
	Y int32
}

func inputMapVirtualKey(uCode, uMapType uint32) uint32 {
	ret, _, _ := procMapVirtualKeyW.Call(uintptr(uCode), uintptr(uMapType))
	return uint32(ret)
}

func inputVkKeyScan(ch rune) uint16 {
	user32 := syscall.NewLazyDLL("user32.dll")
	procVkKeyScanW := user32.NewProc("VkKeyScanW")
	ret, _, _ := procVkKeyScanW.Call(uintptr(ch))
	return uint16(ret)
}

// Input_Key 模拟按下或弹起一个键。
func Input_Key(vk int, down bool) {
	var dwFlags uint32
	if !down {
		dwFlags = KEYEVENTF_KEYUP
	}
	scanCode := inputMapVirtualKey(uint32(vk), MAPVK_VK_TO_VSC)
	procKeybdEvent.Call(uintptr(vk), uintptr(scanCode), uintptr(dwFlags), 0)
}

// Input_KeyCombo 模拟组合键（如 Ctrl+C）。
func Input_KeyCombo(vk1, vk2 int) {
	Input_Key(vk1, true)
	Input_Key(vk2, true)
	Input_Key(vk2, false)
	Input_Key(vk1, false)
}

// Input_KeyPress 模拟按键的完整按下+弹起过程。
func Input_KeyPress(vk int) {
	Input_Key(vk, true)
	Input_Key(vk, false)
}

// Input_GetKeyState 异步获取按键当前是否被按下。
func Input_GetKeyState(vk int) int16 {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vk))
	return int16(ret)
}

// Input_IsKeyDown 判断指定键是否处于按下状态。
func Input_IsKeyDown(vk int) bool {
	return uint16(Input_GetKeyState(vk))&0x8000 != 0
}

// Input_GetKeyToggle 从消息队列获取键状态。
func Input_GetKeyToggle(vk int) int16 {
	ret, _, _ := procGetKeyState.Call(uintptr(vk))
	return int16(ret)
}

// Input_TypeText 逐字符模拟键盘输入中文/英文文本。
func Input_TypeText(text string) {
	for _, c := range text {
		vk := uintptr(inputVkKeyScan(c))
		shift := (vk >> 8) & 1
		if shift != 0 {
			Input_Key(VK_SHIFT, true)
		}
		Input_KeyPress(int(vk & 0xFF))
		if shift != 0 {
			Input_Key(VK_SHIFT, false)
		}
	}
}

// Input_MoveMouse 将鼠标光标移动到指定屏幕坐标。
func Input_MoveMouse(x, y int) {
	procSetCursorPos.Call(uintptr(x), uintptr(y))
}

// Input_GetMousePos 获取当前鼠标光标位置。
func Input_GetMousePos() (int, int) {
	var pt POINT
	procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	return int(pt.X), int(pt.Y)
}

func Input_LeftClick() {
	x, y := Input_GetMousePos()
	procMouseEvent.Call(uintptr(MOUSEEVENTF_LEFTDOWN|MOUSEEVENTF_LEFTUP), uintptr(x), uintptr(y), 0, 0)
}

func Input_LeftDown() {
	x, y := Input_GetMousePos()
	procMouseEvent.Call(uintptr(MOUSEEVENTF_LEFTDOWN), uintptr(x), uintptr(y), 0, 0)
}

func Input_LeftUp() {
	x, y := Input_GetMousePos()
	procMouseEvent.Call(uintptr(MOUSEEVENTF_LEFTUP), uintptr(x), uintptr(y), 0, 0)
}

func Input_RightClick() {
	x, y := Input_GetMousePos()
	procMouseEvent.Call(uintptr(MOUSEEVENTF_RIGHTDOWN|MOUSEEVENTF_RIGHTUP), uintptr(x), uintptr(y), 0, 0)
}

func Input_MiddleClick() {
	x, y := Input_GetMousePos()
	procMouseEvent.Call(uintptr(MOUSEEVENTF_MIDDLEDOWN|MOUSEEVENTF_MIDDLEUP), uintptr(x), uintptr(y), 0, 0)
}

// Input_MouseWheel 在当前鼠标位置滚动滚轮。
func Input_MouseWheel(delta int) {
	x, y := Input_GetMousePos()
	mouseData := int32(delta) * 120
	procMouseEvent.Call(uintptr(MOUSEEVENTF_WHEEL), uintptr(x), uintptr(y), uintptr(mouseData), 0)
}

func Input_ScreenWidth() int {
	ret, _, _ := procGetSystemMetrics.Call(uintptr(SM_CXSCREEN))
	return int(ret)
}

func Input_ScreenHeight() int {
	ret, _, _ := procGetSystemMetrics.Call(uintptr(SM_CYSCREEN))
	return int(ret)
}

// Input_Block 锁定/解锁所有键盘鼠标输入。
func Input_Block(block bool) {
	if block {
		procBlockInput.Call(1)
	} else {
		procBlockInput.Call(0)
	}
}

// ===================== 虚拟键码常量 =====================

const (
	VK_LBUTTON = 0x01
	VK_RBUTTON = 0x02
	VK_MBUTTON = 0x04
	VK_CANCEL  = 0x03
	VK_BACK    = 0x08
	VK_TAB     = 0x09
	VK_CLEAR   = 0x0C
	VK_RETURN  = 0x0D
	VK_SHIFT   = 0x10
	VK_CONTROL = 0x11
	VK_MENU    = 0x12
	VK_PAUSE   = 0x13
	VK_CAPITAL = 0x14
	VK_ESCAPE  = 0x1B
	VK_SPACE   = 0x20
	VK_PRIOR   = 0x21
	VK_NEXT    = 0x22
	VK_END     = 0x23
	VK_HOME    = 0x24
	VK_LEFT    = 0x25
	VK_UP      = 0x26
	VK_RIGHT   = 0x27
	VK_DOWN    = 0x28
	VK_SELECT  = 0x29
	VK_PRINT   = 0x2A
	VK_EXECUTE = 0x2B
	VK_SNAPSHOT = 0x2C
	VK_INSERT  = 0x2D
	VK_DELETE  = 0x2E
	VK_HELP    = 0x2F
	VK_0 = 0x30
	VK_1 = 0x31
	VK_2 = 0x32
	VK_3 = 0x33
	VK_4 = 0x34
	VK_5 = 0x35
	VK_6 = 0x36
	VK_7 = 0x37
	VK_8 = 0x38
	VK_9 = 0x39
	VK_A = 0x41
	VK_B = 0x42
	VK_C = 0x43
	VK_D = 0x44
	VK_E = 0x45
	VK_F = 0x46
	VK_G = 0x47
	VK_H = 0x48
	VK_I = 0x49
	VK_J = 0x4A
	VK_K = 0x4B
	VK_L = 0x4C
	VK_M = 0x4D
	VK_N = 0x4E
	VK_O = 0x4F
	VK_P = 0x50
	VK_Q = 0x51
	VK_R = 0x52
	VK_S = 0x53
	VK_T = 0x54
	VK_U = 0x55
	VK_V = 0x56
	VK_W = 0x57
	VK_X = 0x58
	VK_Y = 0x59
	VK_Z = 0x5A
	VK_LWIN = 0x5B
	VK_RWIN = 0x5C
	VK_APPS = 0x5D
	VK_NUMPAD0  = 0x60
	VK_NUMPAD1  = 0x61
	VK_NUMPAD2  = 0x62
	VK_NUMPAD3  = 0x63
	VK_NUMPAD4  = 0x64
	VK_NUMPAD5  = 0x65
	VK_NUMPAD6  = 0x66
	VK_NUMPAD7  = 0x67
	VK_NUMPAD8  = 0x68
	VK_NUMPAD9  = 0x69
	VK_MULTIPLY = 0x6A
	VK_ADD      = 0x6B
	VK_SEPARATOR = 0x6C
	VK_SUBTRACT = 0x6D
	VK_DECIMAL  = 0x6E
	VK_DIVIDE   = 0x6F
	VK_F1  = 0x70
	VK_F2  = 0x71
	VK_F3  = 0x72
	VK_F4  = 0x73
	VK_F5  = 0x74
	VK_F6  = 0x75
	VK_F7  = 0x76
	VK_F8  = 0x77
	VK_F9  = 0x78
	VK_F10 = 0x79
	VK_F11 = 0x7A
	VK_F12 = 0x7B
	VK_NUMLOCK = 0x90
	VK_SCROLL  = 0x91
	VK_LSHIFT   = 0xA0
	VK_RSHIFT   = 0xA1
	VK_LCONTROL = 0xA2
	VK_RCONTROL = 0xA3
	VK_LMENU    = 0xA4
	VK_RMENU    = 0xA5
)