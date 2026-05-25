//go:build windows

// Windows 多线程模块
// 提供线程创建与管理、临界区同步、事件同步、互斥体、信号量等功能。
// 基于 Windows kernel32.dll API 实现。
package utils

import (
	"syscall"
	"unsafe"
)

var (
	kernel32Thread = syscall.NewLazyDLL("kernel32.dll")

	procCreateThread            = kernel32Thread.NewProc("CreateThread")
	procGetCurrentThreadId      = kernel32Thread.NewProc("GetCurrentThreadId")
	procGetCurrentThread        = kernel32Thread.NewProc("GetCurrentThread")
	procSuspendThread           = kernel32Thread.NewProc("SuspendThread")
	procResumeThread            = kernel32Thread.NewProc("ResumeThread")
	procTerminateThread         = kernel32Thread.NewProc("TerminateThread")

	procWaitForSingleObjectT     = kernel32Thread.NewProc("WaitForSingleObject")
	procWaitForMultipleObjectsT  = kernel32Thread.NewProc("WaitForMultipleObjects")
	procSleep                    = kernel32Thread.NewProc("Sleep")

	procInitializeCriticalSection   = kernel32Thread.NewProc("InitializeCriticalSection")
	procDeleteCriticalSection       = kernel32Thread.NewProc("DeleteCriticalSection")
	procEnterCriticalSection        = kernel32Thread.NewProc("EnterCriticalSection")
	procLeaveCriticalSection        = kernel32Thread.NewProc("LeaveCriticalSection")

	procCreateEventW    = kernel32Thread.NewProc("CreateEventW")
	procSetEvent        = kernel32Thread.NewProc("SetEvent")
	procResetEvent      = kernel32Thread.NewProc("ResetEvent")
	procPulseEvent      = kernel32Thread.NewProc("PulseEvent")

	procCreateMutexW    = kernel32Thread.NewProc("CreateMutexW")
	procReleaseMutex    = kernel32Thread.NewProc("ReleaseMutex")
	procOpenMutexW      = kernel32Thread.NewProc("OpenMutexW")

	procCreateSemaphoreW  = kernel32Thread.NewProc("CreateSemaphoreW")
	procReleaseSemaphore  = kernel32Thread.NewProc("ReleaseSemaphore")
	procOpenSemaphoreW    = kernel32Thread.NewProc("OpenSemaphoreW")

	procCloseHandleThread = kernel32Thread.NewProc("CloseHandle")
)

const (
	WAIT_OBJECT_0    = 0x00000000
	WAIT_TIMEOUT     = 0x00000102
	WAIT_FAILED      = 0xFFFFFFFF
	ERROR_ALREADY_EXISTS = 183

	EVENT_MODIFY_STATE       = 0x0002
	EVENT_ALL_ACCESS         = 0x1F0003
	CREATE_EVENT_MANUAL_RESET  = 0x00000001
	CREATE_EVENT_INITIAL_SET   = 0x00000002

	SEMAPHORE_MODIFY_STATE = 0x0002
	SEMAPHORE_ALL_ACCESS   = 0x1F0003

	MUTEX_MODIFY_STATE = 0x0001
	MUTEX_ALL_ACCESS   = 0x1F0001
)

type RTL_CRITICAL_SECTION struct {
	DebugInfo      uintptr
	LockCount      int32
	RecursionCount int32
	OwningThread   syscall.Handle
	LockSemaphore  syscall.Handle
	SpinCount      uintptr
}

type CRITICAL_SECTION = RTL_CRITICAL_SECTION

// Thread_Create 创建新线程并立即执行指定函数。
func Thread_Create(threadFunc uintptr, param uintptr) (syscall.Handle, error) {
	var threadID uint32
	handle, _, err := procCreateThread.Call(0, 0, threadFunc, param, 0, uintptr(unsafe.Pointer(&threadID)))
	if handle == 0 {
		return 0, err
	}
	return syscall.Handle(handle), nil
}

func Thread_CurrentID() uint32 {
	id, _, _ := procGetCurrentThreadId.Call()
	return uint32(id)
}

func Thread_CurrentPseudoHandle() syscall.Handle {
	handle, _, _ := procGetCurrentThread.Call()
	return syscall.Handle(handle)
}

func Thread_Suspend(hThread syscall.Handle) (uint32, error) {
	ret, _, err := procSuspendThread.Call(uintptr(hThread))
	if ret == 0xFFFFFFFF {
		return 0, err
	}
	return uint32(ret), nil
}

func Thread_Resume(hThread syscall.Handle) (uint32, error) {
	ret, _, err := procResumeThread.Call(uintptr(hThread))
	if ret == 0xFFFFFFFF {
		return 0, err
	}
	return uint32(ret), nil
}

func Thread_Terminate(hThread syscall.Handle, exitCode uint32) error {
	ret, _, err := procTerminateThread.Call(uintptr(hThread), uintptr(exitCode))
	if ret == 0 {
		return err
	}
	return nil
}

func Thread_CloseHandle(handle syscall.Handle) error {
	ret, _, err := procCloseHandleThread.Call(uintptr(handle))
	if ret == 0 {
		return err
	}
	return nil
}

func Thread_Sleep(milliseconds uint32) {
	procSleep.Call(uintptr(milliseconds))
}

func Thread_WaitSingle(handle syscall.Handle, timeoutMs uint32) uint32 {
	ret, _, _ := procWaitForSingleObjectT.Call(uintptr(handle), uintptr(timeoutMs))
	return uint32(ret)
}

func Thread_WaitMultiple(handles []syscall.Handle, waitAll bool, timeoutMs uint32) uint32 {
	var bWaitAll uintptr
	if waitAll {
		bWaitAll = 1
	}
	ret, _, _ := procWaitForMultipleObjectsT.Call(
		uintptr(len(handles)),
		uintptr(unsafe.Pointer(&handles[0])),
		bWaitAll,
		uintptr(timeoutMs),
	)
	return uint32(ret)
}

// Thread_CriticalSection_Create 创建临界区。
func Thread_CriticalSection_Create() (*CRITICAL_SECTION, error) {
	var cs CRITICAL_SECTION
	ret, _, err := procInitializeCriticalSection.Call(uintptr(unsafe.Pointer(&cs)))
	if ret == 0 {
		return nil, err
	}
	return &cs, nil
}

func Thread_CriticalSection_Delete(cs *CRITICAL_SECTION) {
	procDeleteCriticalSection.Call(uintptr(unsafe.Pointer(cs)))
}

func Thread_CriticalSection_Enter(cs *CRITICAL_SECTION) {
	procEnterCriticalSection.Call(uintptr(unsafe.Pointer(cs)))
}

func Thread_CriticalSection_Leave(cs *CRITICAL_SECTION) {
	procLeaveCriticalSection.Call(uintptr(unsafe.Pointer(cs)))
}

// Thread_Event_Create 创建同步事件对象。
func Thread_Event_Create(manualReset bool, initialState bool, name string) (syscall.Handle, error) {
	var bManual, bInitial uintptr
	if manualReset {
		bManual = 1
	}
	if initialState {
		bInitial = 1
	}
	var namePtr *uint16
	if name != "" {
		namePtr, _ = syscall.UTF16PtrFromString(name)
	}
	handle, _, err := procCreateEventW.Call(0, bManual, bInitial, uintptr(unsafe.Pointer(namePtr)))
	if handle == 0 {
		return 0, err
	}
	return syscall.Handle(handle), nil
}

func Thread_Event_Set(hEvent syscall.Handle) error {
	ret, _, err := procSetEvent.Call(uintptr(hEvent))
	if ret == 0 {
		return err
	}
	return nil
}

func Thread_Event_Reset(hEvent syscall.Handle) error {
	ret, _, err := procResetEvent.Call(uintptr(hEvent))
	if ret == 0 {
		return err
	}
	return nil
}

func Thread_Event_Pulse(hEvent syscall.Handle) error {
	ret, _, err := procPulseEvent.Call(uintptr(hEvent))
	if ret == 0 {
		return err
	}
	return nil
}

// Thread_Mutex_Create 创建互斥体，常用于防止程序多开。
func Thread_Mutex_Create(name string) (syscall.Handle, bool, error) {
	var namePtr *uint16
	if name != "" {
		namePtr, _ = syscall.UTF16PtrFromString(name)
	}
	handle, _, err := procCreateMutexW.Call(0, 0, uintptr(unsafe.Pointer(namePtr)))
	if handle == 0 {
		return 0, false, err
	}
	alreadyExists := err == syscall.Errno(ERROR_ALREADY_EXISTS)
	return syscall.Handle(handle), alreadyExists, nil
}

func Thread_Mutex_Open(name string) (syscall.Handle, error) {
	namePtr, _ := syscall.UTF16PtrFromString(name)
	handle, _, err := procOpenMutexW.Call(uintptr(MUTEX_ALL_ACCESS), 0, uintptr(unsafe.Pointer(namePtr)))
	if handle == 0 {
		return 0, err
	}
	return syscall.Handle(handle), nil
}

func Thread_Mutex_Release(hMutex syscall.Handle) error {
	ret, _, err := procReleaseMutex.Call(uintptr(hMutex))
	if ret == 0 {
		return err
	}
	return nil
}

// Thread_Semaphore_Create 创建信号量对象。
func Thread_Semaphore_Create(initialCount int32, maxCount int32, name string) (syscall.Handle, error) {
	var namePtr *uint16
	if name != "" {
		namePtr, _ = syscall.UTF16PtrFromString(name)
	}
	handle, _, err := procCreateSemaphoreW.Call(0, uintptr(initialCount), uintptr(maxCount), uintptr(unsafe.Pointer(namePtr)))
	if handle == 0 {
		return 0, err
	}
	return syscall.Handle(handle), nil
}

func Thread_Semaphore_Open(name string) (syscall.Handle, error) {
	namePtr, _ := syscall.UTF16PtrFromString(name)
	handle, _, err := procOpenSemaphoreW.Call(uintptr(SEMAPHORE_ALL_ACCESS), 0, uintptr(unsafe.Pointer(namePtr)))
	if handle == 0 {
		return 0, err
	}
	return syscall.Handle(handle), nil
}

func Thread_Semaphore_Release(hSemaphore syscall.Handle, count int32) (int32, error) {
	var prevCount int32
	ret, _, err := procReleaseSemaphore.Call(uintptr(hSemaphore), uintptr(count), uintptr(unsafe.Pointer(&prevCount)))
	if ret == 0 {
		return 0, err
	}
	return prevCount, nil
}