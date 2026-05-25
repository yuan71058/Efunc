// API Hook 模块
// 基于 MinHook 库实现的 API 钩子（Hook）功能。
// 支持为目标函数/导出 API 安装钩子，拦截并修改函数行为。
package utils

import (
	"unsafe"

	"github.com/yuan71058/Efunc/en/utils/HHook"
)

// Hook_Init 初始化 MinHook 库。
func Hook_Init() error {
	return HHook.Init()
}

// Hook_Uninit 卸载 MinHook 库。
func Hook_Uninit() error {
	return HHook.Uninit()
}

// Hook_Install 为目标函数安装钩子，在禁用状态下创建。
//
// 参数:
//   - targetAddr: 目标函数的地址（uintptr）
//   - callbackAddr: 回调函数的地址（uintptr），目标函数被调用时将跳转到此地址
//
// 返回:
//   - uintptr: 原始函数地址（trampoline），用于调用原始函数
//   - error: 错误信息
func Hook_Install(targetAddr, callbackAddr uintptr) (uintptr, error) {
	orig, err := HHook.CreateHook(unsafe.Pointer(targetAddr), unsafe.Pointer(callbackAddr))
	return uintptr(orig), err
}

// Hook_InstallApi 通过模块名和函数名为指定 API 安装钩子。
//
// 参数:
//   - moduleName: DLL 模块名称，如 "kernel32.dll"
//   - funcName: 导出函数名称，如 "MessageBoxW"
//   - callbackAddr: 回调函数的地址（uintptr）
//
// 返回:
//   - uintptr: 原始函数地址（trampoline），用于调用原始函数
//   - error: 错误信息
func Hook_InstallApi(moduleName, funcName string, callbackAddr uintptr) (uintptr, error) {
	orig, err := HHook.CreateHookApi(moduleName, funcName, unsafe.Pointer(callbackAddr))
	return uintptr(orig), err
}

// Hook_InstallApiEx 通过模块名和函数名为指定 API 安装钩子（扩展版）。
//
// 参数:
//   - moduleName: DLL 模块名称
//   - funcName: 导出函数名称
//   - callbackAddr: 回调函数的地址（uintptr）
//
// 返回:
//   - origAddr: 原始函数地址（trampoline）
//   - targetAddr: 目标函数地址
//   - err: 错误信息
func Hook_InstallApiEx(moduleName, funcName string, callbackAddr uintptr) (origAddr, targetAddr uintptr, err error) {
	orig, tgt, e := HHook.CreateHookApiEx(moduleName, funcName, unsafe.Pointer(callbackAddr))
	return uintptr(orig), uintptr(tgt), e
}

// Hook_Remove 卸载指定钩子。
func Hook_Remove(targetAddr uintptr) error {
	return HHook.RemoveHook(unsafe.Pointer(targetAddr))
}

// Hook_Enable 启用指定钩子。
func Hook_Enable(targetAddr uintptr) error {
	return HHook.EnableHook(unsafe.Pointer(targetAddr))
}

// Hook_Disable 禁用指定钩子。
func Hook_Disable(targetAddr uintptr) error {
	return HHook.DisableHook(unsafe.Pointer(targetAddr))
}

// Hook_EnableAll 启用全部钩子。
func Hook_EnableAll() error {
	return HHook.EnableAllHooks()
}

// Hook_DisableAll 禁用全部钩子。
func Hook_DisableAll() error {
	return HHook.DisableAllHooks()
}

// Hook_QueueEnable 排队启用钩子。
func Hook_QueueEnable(targetAddr uintptr) error {
	return HHook.QueueEnableHook(unsafe.Pointer(targetAddr))
}

// Hook_QueueDisable 排队禁用钩子。
func Hook_QueueDisable(targetAddr uintptr) error {
	return HHook.QueueDisableHook(unsafe.Pointer(targetAddr))
}

// Hook_ApplyQueued 应用所有排队中的钩子操作。
func Hook_ApplyQueued() error {
	return HHook.ApplyQueued()
}

// Hook_GetStatusText 将状态码转换为中文描述文本。
func Hook_GetStatusText(statusCode int) string {
	return HHook.StatusToString(statusCode)
}