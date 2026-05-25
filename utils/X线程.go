//go:build windows

// 多线程模块
// 提供线程创建与管理、临界区同步、事件同步、互斥体、信号量等功能。
// 基于 Windows kernel32.dll API 实现，支持 CreateThread、Sleep、WaitForSingleObject、WaitForMultipleObjects 等底层调用。
// 包含完整的线程同步原语封装（CriticalSection、Event、Mutex、Semaphore），适用于多线程编程场景。
package utils

import (
	"errors"
	"syscall"
	"unsafe"
)

var (
	kernel32X线程 = syscall.NewLazyDLL("kernel32.dll") // kernel32.dll 引用

	// 线程创建与管理 API
	procCreateThread         = kernel32X线程.NewProc("CreateThread")         // 创建新线程
	procGetCurrentThreadId   = kernel32X线程.NewProc("GetCurrentThreadId")   // 获取当前线程 ID
	procGetCurrentThreadX线程 = kernel32X线程.NewProc("GetCurrentThread")     // 获取当前线程伪句柄
	procSuspendThread        = kernel32X线程.NewProc("SuspendThread")        // 挂起线程
	procResumeThread         = kernel32X线程.NewProc("ResumeThread")         // 恢复线程
	procTerminateThread      = kernel32X线程.NewProc("TerminateThread")      // 强制终止线程

	// 线程等待与同步
	procWaitForSingleObject   = kernel32X线程.NewProc("WaitForSingleObject")    // 等待单个对象
	procWaitForMultipleObjects = kernel32X线程.NewProc("WaitForMultipleObjects") // 等待多个对象
	procSleep                 = kernel32X线程.NewProc("Sleep")                  // 暂停当前线程

	// 临界区（Critical Section）
	procInitializeCriticalSection   = kernel32X线程.NewProc("InitializeCriticalSection")   // 初始化临界区
	procDeleteCriticalSection       = kernel32X线程.NewProc("DeleteCriticalSection")       // 删除临界区
	procEnterCriticalSection        = kernel32X线程.NewProc("EnterCriticalSection")        // 进入临界区
	procLeaveCriticalSection        = kernel32X线程.NewProc("LeaveCriticalSection")        // 离开临界区
	procInitializeCriticalSectionEx = kernel32X线程.NewProc("InitializeCriticalSectionEx") // 初始化临界区（带旋转计数）

	// 事件（Event）
	procCreateEventW     = kernel32X线程.NewProc("CreateEventW")     // 创建事件对象
	procSetEvent         = kernel32X线程.NewProc("SetEvent")         // 设置事件为有信号
	procResetEvent       = kernel32X线程.NewProc("ResetEvent")       // 设置事件为无信号
	procPulseEvent       = kernel32X线程.NewProc("PulseEvent")       // 脉冲触发事件

	// 互斥体（Mutex）
	procCreateMutexW      = kernel32X线程.NewProc("CreateMutexW")      // 创建互斥体
	procReleaseMutex      = kernel32X线程.NewProc("ReleaseMutex")      // 释放互斥体
	procOpenMutexW        = kernel32X线程.NewProc("OpenMutexW")        // 打开已存在的互斥体

	// 信号量（Semaphore）
	procCreateSemaphoreW  = kernel32X线程.NewProc("CreateSemaphoreW")  // 创建信号量
	procReleaseSemaphore  = kernel32X线程.NewProc("ReleaseSemaphore")  // 释放信号量（增加计数）
	procOpenSemaphoreW    = kernel32X线程.NewProc("OpenSemaphoreW")    // 打开已存在的信号量

	// 句柄操作
	procCloseHandleX线程 = kernel32X线程.NewProc("CloseHandle") // 关闭句柄
)

// ===================== 等待常量 =====================

const (
	WAIT_OBJECT_0    = 0x00000000 // 等待对象进入有信号状态
	WAIT_TIMEOUT     = 0x00000102 // 等待超时
	WAIT_FAILED      = 0xFFFFFFFF // 等待失败
	INFINITE         = 0xFFFFFFFF // 无限等待
	ERROR_ALREADY_EXISTS = 183   // 对象已存在（用于互斥体防重复运行）

	// 事件创建标志
	EVENT_MODIFY_STATE = 0x0002 // 修改事件状态权限
	EVENT_ALL_ACCESS   = 0x1F0003 // 所有事件权限
	CREATE_EVENT_MANUAL_RESET  = 0x00000001 // 手动重置事件（需显式 ResetEvent）
	CREATE_EVENT_INITIAL_SET   = 0x00000002 // 初始状态为有信号

	// 信号量权限
	SEMAPHORE_MODIFY_STATE = 0x0002 // 修改信号量状态权限
	SEMAPHORE_ALL_ACCESS   = 0x1F0003 // 所有信号量权限

	// 互斥体权限
	MUTEX_MODIFY_STATE = 0x0001 // 修改互斥体状态权限
	MUTEX_ALL_ACCESS   = 0x1F0001 // 所有互斥体权限
)

// ===================== 结构体定义 =====================

// RTL_CRITICAL_SECTION 临界区结构体（Windows 内部使用）
type RTL_CRITICAL_SECTION struct {
	DebugInfo      uintptr   // 调试信息指针
	LockCount      int32     // 锁定计数（-1 未锁定，>=0 被锁定的次数）
	RecursionCount int32     // 递归计数
	OwningThread   syscall.Handle // 拥有线程句柄
	LockSemaphore  syscall.Handle // 锁定信号量
	SpinCount      uintptr   // 自旋计数（多处理器下优化）
}

// CRITICAL_SECTION 导出名，与 RTL_CRITICAL_SECTION 同一类型
type CRITICAL_SECTION = RTL_CRITICAL_SECTION

// ===================== 线程创建与管理 =====================

// X线程_创建 创建新线程并立即执行指定函数。
// 参数 线程函数：线程入口函数地址（可通过 syscall.NewCallback 获取）
// 参数 参数：传递给线程函数的参数
// 返回 syscall.Handle：线程句柄（使用完毕后需调用 X线程_关闭句柄）
// 返回 error：失败时返回错误信息
func X线程_创建(线程函数 uintptr, 参数 uintptr) (syscall.Handle, error) {
	var threadID uint32
	handle, _, err := procCreateThread.Call(0, 0, 线程函数, 参数, 0, uintptr(unsafe.Pointer(&threadID)))
	if handle == 0 {
		return 0, err
	}
	return syscall.Handle(handle), nil
}

// X线程_取当前ID 获取当前线程的 ID。
// 返回 uint32：当前线程 ID
func X线程_取当前ID() uint32 {
	id, _, _ := procGetCurrentThreadId.Call()
	return uint32(id)
}

// X线程_取当前伪句柄 获取当前线程的伪句柄（不可跨进程使用，使用后无需 CloseHandle）。
// 返回 syscall.Handle：当前线程伪句柄
func X线程_取当前伪句柄() syscall.Handle {
	handle, _, _ := procGetCurrentThreadX线程.Call()
	return syscall.Handle(handle)
}

// X线程_挂起 挂起指定线程（增加挂起计数）。
// 参数 hThread：线程句柄
// 返回 uint32：之前的挂起计数
// 返回 error：失败时返回错误信息
func X线程_挂起(hThread syscall.Handle) (uint32, error) {
	ret, _, err := procSuspendThread.Call(uintptr(hThread))
	if ret == 0xFFFFFFFF {
		return 0, err
	}
	return uint32(ret), nil
}

// X线程_恢复 恢复（减少挂起计数）指定线程。
// 参数 hThread：线程句柄
// 返回 uint32：之前的挂起计数
// 返回 error：失败时返回错误信息
func X线程_恢复(hThread syscall.Handle) (uint32, error) {
	ret, _, err := procResumeThread.Call(uintptr(hThread))
	if ret == 0xFFFFFFFF {
		return 0, err
	}
	return uint32(ret), nil
}

// X线程_终止 强制终止线程（不推荐使用，可能导致资源泄漏）。
// 参数 hThread：线程句柄
// 参数 退出码：线程退出码
// 返回 error：失败时返回错误信息
func X线程_终止(hThread syscall.Handle, 退出码 uint32) error {
	ret, _, err := procTerminateThread.Call(uintptr(hThread), uintptr(退出码))
	if ret == 0 {
		return err
	}
	return nil
}

// X线程_关闭句柄 关闭内核对象句柄，释放系统资源。
// 参数 handle：要关闭的句柄
// 返回 error：失败时返回错误信息
func X线程_关闭句柄(handle syscall.Handle) error {
	ret, _, err := procCloseHandleX线程.Call(uintptr(handle))
	if ret == 0 {
		return err
	}
	return nil
}

// X线程_延时 暂停当前线程指定毫秒数。
// 参数 毫秒：暂停时间（毫秒）
func X线程_延时(毫秒 uint32) {
	procSleep.Call(uintptr(毫秒))
}

// ===================== 线程等待 =====================

// X线程_等待单个 等待单个内核对象变为有信号状态。
// 参数 handle：内核对象句柄（线程、事件、互斥体、信号量等）
// 参数 超时毫秒：超时时间（毫秒），INFINITE 表示无限等待
// 返回 uint32：WAIT_OBJECT_0 表示成功，WAIT_TIMEOUT 表示超时，WAIT_FAILED 表示失败
func X线程_等待单个(handle syscall.Handle, 超时毫秒 uint32) uint32 {
	ret, _, _ := procWaitForSingleObject.Call(uintptr(handle), uintptr(超时毫秒))
	return uint32(ret)
}

// X线程_等待多个 等待多个内核对象中任意一个或全部变为有信号状态。
// 参数 handles：内核对象句柄切片
// 参数 等待全部：true 等待全部，false 等待任意一个
// 参数 超时毫秒：超时时间（毫秒），INFINITE 表示无限等待
// 返回 uint32：WAIT_OBJECT_0 + n 表示第 n 个对象有信号，WAIT_TIMEOUT 表示超时
func X线程_等待多个(handles []syscall.Handle, 等待全部 bool, 超时毫秒 uint32) uint32 {
	var bWaitAll uintptr
	if 等待全部 {
		bWaitAll = 1
	}
	ret, _, _ := procWaitForMultipleObjects.Call(
		uintptr(len(handles)),
		uintptr(unsafe.Pointer(&handles[0])),
		bWaitAll,
		uintptr(超时毫秒),
	)
	return uint32(ret)
}

// ===================== 临界区（Critical Section） =====================

// X线程_临界区_创建 创建并初始化临界区对象。
// 返回 *CRITICAL_SECTION：临界区指针（使用完毕后需调用 X线程_临界区_销毁）
// 返回 error：失败时返回错误信息
func X线程_临界区_创建() (*CRITICAL_SECTION, error) {
	var cs CRITICAL_SECTION
	ret, _, err := procInitializeCriticalSection.Call(uintptr(unsafe.Pointer(&cs)))
	if ret == 0 {
		return nil, err
	}
	return &cs, nil
}

// X线程_临界区_销毁 销毁临界区并释放系统资源。
// 参数 cs：临界区指针
func X线程_临界区_销毁(cs *CRITICAL_SECTION) {
	procDeleteCriticalSection.Call(uintptr(unsafe.Pointer(cs)))
}

// X线程_临界区_进入 进入临界区（如果已被占用则阻塞等待）。
// 参数 cs：临界区指针
func X线程_临界区_进入(cs *CRITICAL_SECTION) {
	procEnterCriticalSection.Call(uintptr(unsafe.Pointer(cs)))
}

// X线程_临界区_离开 离开临界区，释放锁。
// 参数 cs：临界区指针
func X线程_临界区_离开(cs *CRITICAL_SECTION) {
	procLeaveCriticalSection.Call(uintptr(unsafe.Pointer(cs)))
}

// ===================== 事件（Event） =====================

// X线程_事件_创建 创建一个同步事件对象。
// 参数 手动重置：true 为手动重置事件，false 为自动重置
// 参数 初始状态：true 为初始有信号
// 参数 名称：事件名称（空字符串表示匿名事件）
// 返回 syscall.Handle：事件句柄
// 返回 error：失败时返回错误信息
func X线程_事件_创建(手动重置 bool, 初始状态 bool, 名称 string) (syscall.Handle, error) {
	var (
		bManualReset uintptr
		bInitialState uintptr
	)
	if 手动重置 {
		bManualReset = 1
	}
	if 初始状态 {
		bInitialState = 1
	}

	var namePtr *uint16
	if 名称 != "" {
		namePtr, _ = syscall.UTF16PtrFromString(名称)
	}

	handle, _, err := procCreateEventW.Call(0, bManualReset, bInitialState, uintptr(unsafe.Pointer(namePtr)))
	if handle == 0 {
		return 0, err
	}
	return syscall.Handle(handle), nil
}

// X线程_事件_设置 将事件设置为有信号状态，唤醒等待线程。
// 参数 hEvent：事件句柄
// 返回 error：失败时返回错误信息
func X线程_事件_设置(hEvent syscall.Handle) error {
	ret, _, err := procSetEvent.Call(uintptr(hEvent))
	if ret == 0 {
		return err
	}
	return nil
}

// X线程_事件_重置 将事件重置为无信号状态。
// 参数 hEvent：事件句柄
// 返回 error：失败时返回错误信息
func X线程_事件_重置(hEvent syscall.Handle) error {
	ret, _, err := procResetEvent.Call(uintptr(hEvent))
	if ret == 0 {
		return err
	}
	return nil
}

// X线程_事件_脉冲 脉冲触发事件：立即设置并重置，唤醒所有等待线程。
// 参数 hEvent：事件句柄
// 返回 error：失败时返回错误信息
func X线程_事件_脉冲(hEvent syscall.Handle) error {
	ret, _, err := procPulseEvent.Call(uintptr(hEvent))
	if ret == 0 {
		return err
	}
	return nil
}

// ===================== 互斥体（Mutex） =====================

// X线程_互斥体_创建 创建一个互斥体对象，常用于防止程序多开。
// 参数 名称：互斥体名称（空字符串表示匿名）
// 返回 syscall.Handle：互斥体句柄
// 返回 bool：true 表示已存在（已有同名互斥体），false 表示新创建
// 返回 error：失败时返回错误信息
func X线程_互斥体_创建(名称 string) (syscall.Handle, bool, error) {
	var namePtr *uint16
	if 名称 != "" {
		namePtr, _ = syscall.UTF16PtrFromString(名称)
	}

	handle, _, err := procCreateMutexW.Call(0, 0, uintptr(unsafe.Pointer(namePtr)))
	if handle == 0 {
		return 0, false, err
	}

	// 检查是否已存在（ERROR_ALREADY_EXISTS = 183）
	alreadyExists := err == syscall.Errno(ERROR_ALREADY_EXISTS)
	return syscall.Handle(handle), alreadyExists, nil
}

// X线程_互斥体_打开 打开一个已存在的命名互斥体。
// 参数 名称：互斥体名称
// 返回 syscall.Handle：互斥体句柄
// 返回 error：失败时返回错误信息
func X线程_互斥体_打开(名称 string) (syscall.Handle, error) {
	namePtr, _ := syscall.UTF16PtrFromString(名称)
	handle, _, err := procOpenMutexW.Call(uintptr(MUTEX_ALL_ACCESS), 0, uintptr(unsafe.Pointer(namePtr)))
	if handle == 0 {
		return 0, err
	}
	return syscall.Handle(handle), nil
}

// X线程_互斥体_释放 释放互斥体所有权。
// 参数 hMutex：互斥体句柄
// 返回 error：失败时返回错误信息
func X线程_互斥体_释放(hMutex syscall.Handle) error {
	ret, _, err := procReleaseMutex.Call(uintptr(hMutex))
	if ret == 0 {
		return err
	}
	return nil
}

// ===================== 信号量（Semaphore） =====================

// X线程_信号量_创建 创建一个信号量对象。
// 参数 初始计数：信号量初始可用资源数
// 参数 最大计数：信号量最大可用资源数
// 参数 名称：信号量名称（空字符串表示匿名）
// 返回 syscall.Handle：信号量句柄
// 返回 error：失败时返回错误信息
func X线程_信号量_创建(初始计数 int32, 最大计数 int32, 名称 string) (syscall.Handle, error) {
	var namePtr *uint16
	if 名称 != "" {
		namePtr, _ = syscall.UTF16PtrFromString(名称)
	}

	handle, _, err := procCreateSemaphoreW.Call(
		0,
		uintptr(初始计数),
		uintptr(最大计数),
		uintptr(unsafe.Pointer(namePtr)),
	)
	if handle == 0 {
		return 0, err
	}
	return syscall.Handle(handle), nil
}

// X线程_信号量_打开 打开一个已存在的命名信号量。
// 参数 名称：信号量名称
// 返回 syscall.Handle：信号量句柄
// 返回 error：失败时返回错误信息
func X线程_信号量_打开(名称 string) (syscall.Handle, error) {
	namePtr, _ := syscall.UTF16PtrFromString(名称)
	handle, _, err := procOpenSemaphoreW.Call(uintptr(SEMAPHORE_ALL_ACCESS), 0, uintptr(unsafe.Pointer(namePtr)))
	if handle == 0 {
		return 0, err
	}
	return syscall.Handle(handle), nil
}

// X线程_信号量_释放 释放信号量（增加可用资源计数）。
// 参数 hSemaphore：信号量句柄
// 参数 释放数量：要增加的计数
// 返回 int32：释放前的可用计数
// 返回 error：失败时返回错误信息
func X线程_信号量_释放(hSemaphore syscall.Handle, 释放数量 int32) (int32, error) {
	var previousCount int32
	ret, _, err := procReleaseSemaphore.Call(
		uintptr(hSemaphore),
		uintptr(释放数量),
		uintptr(unsafe.Pointer(&previousCount)),
	)
	if ret == 0 {
		return 0, err
	}
	return previousCount, nil
}

// ===================== 高级线程封装（基于 goroutine） =====================

var (
	Err线程已创建 = errors.New("线程已经创建，无法重复创建")
	Err线程未创建 = errors.New("线程尚未创建")
	Err线程已结束 = errors.New("线程已经结束")
)

// G线程 基于 goroutine 的协作式线程封装。
// 提供创建、挂起、恢复、终止等操作，通过 channel 控制生命周期。
type G线程 struct {
	创建     chan bool // 创建信号
	挂起     chan bool // 挂起信号
	恢复     chan bool // 恢复信号
	结束     chan bool // 结束信号
	已结束   chan bool // 已结束通知
	执行函数 func()    // 线程执行函数
	已运行   bool       // 是否已启动
	已销毁   bool       // 是否已销毁
}

// X线程_协程_创建 创建基于 goroutine 的协作式线程。
// 参数 执行函数：线程需要执行的函数
// 返回 *G线程：线程控制对象
func X线程_协程_创建(执行函数 func()) *G线程 {
	return &G线程{
		创建:     make(chan bool),
		挂起:     make(chan bool),
		恢复:     make(chan bool),
		结束:     make(chan bool),
		已结束:   make(chan bool),
		执行函数: 执行函数,
	}
}

// 运行 启动线程执行（内部方法）。
func (t *G线程) 运行() error {
	if t.已运行 {
		return Err线程已创建
	}
	t.已运行 = true
	go func() {
		defer func() {
			t.已结束 <- true
		}()
		for {
			select {
			case <-t.创建:
				t.执行函数() // 执行一次用户函数
			case <-t.挂起:
				select {
				case <-t.恢复:
					// 继续循环
				case <-t.结束:
					return
				}
			case <-t.结束:
				return
			}
		}
	}()
	return nil
}

// 启动 启动线程并立即执行用户函数。
// 返回 error：重复创建时返回错误
func (t *G线程) 启动() error {
	if err := t.运行(); err != nil {
		return err
	}
	t.创建 <- true
	return nil
}

// 暂停 挂起线程执行。
// 返回 error：线程已结束时返回错误
func (t *G线程) 暂停() error {
	if t.已销毁 {
		return Err线程已结束
	}
	t.挂起 <- true
	return nil
}

// 继续 恢复挂起的线程。
// 返回 error：线程已结束时返回错误
func (t *G线程) 继续() error {
	if t.已销毁 {
		return Err线程已结束
	}
	t.恢复 <- true
	return nil
}

// 退出 请求线程退出（优雅关闭）。
func (t *G线程) 退出() {
	if !t.已销毁 {
		close(t.结束)
		<-t.已结束 // 等待线程真正结束
		t.已销毁 = true
	}
}