package utils

import (
	"unsafe"

	"github.com/yuan71058/Efunc/utils/HHook"
)

func HHook_初始化() error {
	return HHook.Init()
}

func HHook_卸载() error {
	return HHook.Uninit()
}

// HHook_安装Hook 为目标函数安装钩子，在禁用状态下创建。
//
// 参数:
//   - 目标地址: 目标函数的地址（uintptr）
//   - 回调地址: 回调函数的地址（uintptr），目标函数被调用时将跳转到此地址
//
// 返回:
//   - uintptr: 原始函数地址（trampoline），用于调用原始函数
//   - error: 错误信息
func HHook_安装Hook(目标地址, 回调地址 uintptr) (uintptr, error) {
	orig, err := HHook.CreateHook(unsafe.Pointer(目标地址), unsafe.Pointer(回调地址))
	return uintptr(orig), err
}

// HHook_安装ApiHook 通过模块名和函数名为指定 API 安装钩子。
//
// 参数:
//   - 模块名: DLL 模块名称，如 "kernel32.dll"
//   - 函数名: 导出函数名称，如 "MessageBoxW"
//   - 回调地址: 回调函数的地址（uintptr）
//
// 返回:
//   - uintptr: 原始函数地址（trampoline），用于调用原始函数
//   - error: 错误信息
func HHook_安装ApiHook(模块名, 函数名 string, 回调地址 uintptr) (uintptr, error) {
	orig, err := HHook.CreateHookApi(模块名, 函数名, unsafe.Pointer(回调地址))
	return uintptr(orig), err
}

// HHook_安装ApiHookEx 通过模块名和函数名为指定 API 安装钩子（扩展版）。
//
// 参数:
//   - 模块名: DLL 模块名称
//   - 函数名: 导出函数名称
//   - 回调地址: 回调函数的地址（uintptr）
//
// 返回:
//   - uintptr: 原始函数地址（trampoline）
//   - uintptr: 目标函数地址
//   - error: 错误信息
func HHook_安装ApiHookEx(模块名, 函数名 string, 回调地址 uintptr) (原地址, 目标地址 uintptr, err error) {
	orig, tgt, e := HHook.CreateHookApiEx(模块名, 函数名, unsafe.Pointer(回调地址))
	return uintptr(orig), uintptr(tgt), e
}

func HHook_卸载Hook(目标地址 uintptr) error {
	return HHook.RemoveHook(unsafe.Pointer(目标地址))
}

func HHook_启用Hook(目标地址 uintptr) error {
	return HHook.EnableHook(unsafe.Pointer(目标地址))
}

func HHook_禁用Hook(目标地址 uintptr) error {
	return HHook.DisableHook(unsafe.Pointer(目标地址))
}

func HHook_启用全部Hook() error {
	return HHook.EnableAllHooks()
}

func HHook_禁用全部Hook() error {
	return HHook.DisableAllHooks()
}

func HHook_排队启用Hook(目标地址 uintptr) error {
	return HHook.QueueEnableHook(unsafe.Pointer(目标地址))
}

func HHook_排队禁用Hook(目标地址 uintptr) error {
	return HHook.QueueDisableHook(unsafe.Pointer(目标地址))
}

func HHook_应用排队() error {
	return HHook.ApplyQueued()
}

// HHook_取状态文本 将状态码转换为中文描述文本。
func HHook_取状态文本(状态码 int) string {
	return HHook.StatusToString(状态码)
}