package HHook

/*
#cgo CFLAGS: -Wall -O2
#cgo LDFLAGS: -lkernel32

#include "MinHook.h"

static inline MH_STATUS go_MH_CreateHook(void *pTarget, void *pDetour, LPVOID *ppOriginal) { return MH_CreateHook((LPVOID)pTarget, (LPVOID)pDetour, ppOriginal); }
static inline MH_STATUS go_MH_RemoveHook(void *pTarget) { return MH_RemoveHook((LPVOID)pTarget); }
static inline MH_STATUS go_MH_EnableHook(void *pTarget) { return MH_EnableHook((LPVOID)pTarget); }
static inline MH_STATUS go_MH_DisableHook(void *pTarget) { return MH_DisableHook((LPVOID)pTarget); }
static inline MH_STATUS go_MH_QueueEnableHook(void *pTarget) { return MH_QueueEnableHook((LPVOID)pTarget); }
static inline MH_STATUS go_MH_QueueDisableHook(void *pTarget) { return MH_QueueDisableHook((LPVOID)pTarget); }
static inline MH_STATUS go_MH_EnableAllHooks(void) { return MH_EnableHook(MH_ALL_HOOKS); }
static inline MH_STATUS go_MH_DisableAllHooks(void) { return MH_DisableHook(MH_ALL_HOOKS); }
*/
import "C"
import (
	"errors"
	"unsafe"

	"golang.org/x/sys/windows"
)

var statusMessages = map[int]string{
	-1: "未知错误",
	0:  "成功",
	1:  "已初始化",
	2:  "未初始化",
	3:  "已创建",
	4:  "未创建",
	5:  "已启用",
	6:  "已禁用",
	7:  "不可执行",
	8:  "不支持的函数",
	9:  "内存分配失败",
	10: "内存保护失败",
	11: "模块未找到",
	12: "函数未找到",
}

func statusToError(status C.MH_STATUS) error {
	if status == C.MH_OK {
		return nil
	}
	msg, ok := statusMessages[int(status)]
	if !ok {
		msg = C.GoString(C.MH_StatusToString(status))
	}
	return errors.New(msg)
}

func Init() error {
	return statusToError(C.MH_Initialize())
}

func Uninit() error {
	return statusToError(C.MH_Uninitialize())
}

func CreateHook(target, detour unsafe.Pointer) (original unsafe.Pointer, err error) {
	var orig C.LPVOID
	status := C.go_MH_CreateHook(target, detour, &orig)
	if status != C.MH_OK {
		return nil, statusToError(status)
	}
	return unsafe.Pointer(orig), nil
}

func CreateHookApi(module, procName string, detour unsafe.Pointer) (original unsafe.Pointer, err error) {
	moduleWide, err := windows.UTF16PtrFromString(module)
	if err != nil {
		return nil, err
	}

	cProcName := C.CString(procName)
	defer C.free(unsafe.Pointer(cProcName))

	var orig C.LPVOID
	status := C.MH_CreateHookApi(
		(*C.WCHAR)(unsafe.Pointer(moduleWide)),
		cProcName,
		C.LPVOID(detour),
		&orig,
	)
	if status != C.MH_OK {
		return nil, statusToError(status)
	}
	return unsafe.Pointer(orig), nil
}

func CreateHookApiEx(module, procName string, detour unsafe.Pointer) (original, target unsafe.Pointer, err error) {
	moduleWide, err := windows.UTF16PtrFromString(module)
	if err != nil {
		return nil, nil, err
	}

	cProcName := C.CString(procName)
	defer C.free(unsafe.Pointer(cProcName))

	var orig C.LPVOID
	var tgt C.LPVOID
	status := C.MH_CreateHookApiEx(
		(*C.WCHAR)(unsafe.Pointer(moduleWide)),
		cProcName,
		C.LPVOID(detour),
		&orig,
		&tgt,
	)
	if status != C.MH_OK {
		return nil, nil, statusToError(status)
	}
	return unsafe.Pointer(orig), unsafe.Pointer(tgt), nil
}

func RemoveHook(target unsafe.Pointer) error {
	return statusToError(C.go_MH_RemoveHook(target))
}

func EnableHook(target unsafe.Pointer) error {
	return statusToError(C.go_MH_EnableHook(target))
}

func DisableHook(target unsafe.Pointer) error {
	return statusToError(C.go_MH_DisableHook(target))
}

func EnableAllHooks() error {
	return statusToError(C.go_MH_EnableAllHooks())
}

func DisableAllHooks() error {
	return statusToError(C.go_MH_DisableAllHooks())
}

func QueueEnableHook(target unsafe.Pointer) error {
	return statusToError(C.go_MH_QueueEnableHook(target))
}

func QueueDisableHook(target unsafe.Pointer) error {
	return statusToError(C.go_MH_QueueDisableHook(target))
}

func ApplyQueued() error {
	return statusToError(C.MH_ApplyQueued())
}

func StatusToString(status int) string {
	msg, ok := statusMessages[status]
	if !ok {
		return C.GoString(C.MH_StatusToString(C.MH_STATUS(status)))
	}
	return msg
}