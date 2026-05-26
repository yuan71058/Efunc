//go:build windows

// 内存操作模块
// 提供进程内存读写、内存搜索（AOB 特征码搜索）、虚拟内存管理、进程枚举与管理等功能。
// 基于 Windows kernel32.dll API 实现，支持 ReadProcessMemory、WriteProcessMemory、VirtualAllocEx、VirtualProtectEx 等底层调用。
// 包含进程快照（Toolhelp32Snapshot）枚举和多级指针读取能力。
package utils

import (
	"errors"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	kernel32H内存 = syscall.NewLazyDLL("kernel32.dll") // kernel32.dll 引用

	// 进程/内存操作 API 函数指针
	procOpenProcessH              = kernel32H内存.NewProc("OpenProcess")              // 打开进程句柄
	procCloseHandleH              = kernel32H内存.NewProc("CloseHandle")              // 关闭句柄
	procReadProcessMemory         = kernel32H内存.NewProc("ReadProcessMemory")        // 读取进程内存
	procWriteProcessMemory        = kernel32H内存.NewProc("WriteProcessMemory")       // 写入进程内存
	procVirtualAllocEx            = kernel32H内存.NewProc("VirtualAllocEx")            // 在远程进程分配内存
	procVirtualFreeEx             = kernel32H内存.NewProc("VirtualFreeEx")             // 释放远程进程内存
	procCreateRemoteThread        = kernel32H内存.NewProc("CreateRemoteThread")        // 创建远程线程
	procVirtualProtectEx          = kernel32H内存.NewProc("VirtualProtectEx")          // 修改远程进程内存保护属性
	procVirtualQueryEx            = kernel32H内存.NewProc("VirtualQueryEx")            // 查询远程进程内存信息

	// 进程/模块快照 API（procCreateToolhelp32Snapshot/procProcess32FirstW/procProcess32NextW 见 C进程.go）
	procModule32FirstW            = kernel32H内存.NewProc("Module32FirstW")            // 枚举第一个模块
	procModule32NextW             = kernel32H内存.NewProc("Module32NextW")             // 枚举下一个模块
	procGetModuleBaseNameW        = kernel32H内存.NewProc("GetModuleBaseNameW")        // 获取模块基础名称

	// 进程控制 API
	procGetExitCodeProcessH       = kernel32H内存.NewProc("GetExitCodeProcess")       // 获取进程退出码
	procTerminateProcessH         = kernel32H内存.NewProc("TerminateProcess")         // 终止进程
)

// ===================== 进程权限常量 =====================

const (
	PROCESS_VM_READ           = 0x0010  // 读取进程内存权限
	PROCESS_VM_WRITE          = 0x0020  // 写入进程内存权限
	PROCESS_VM_OPERATION      = 0x0008  // 操作进程内存权限
	PROCESS_CREATE_THREAD     = 0x0002  // 创建远程线程权限
	PROCESS_ALL_ACCESS        = 0x1F0FFF // 所有权限

	// 内存保护属性常量
	PAGE_NOACCESS          = 0x01 // 禁止访问
	PAGE_READONLY          = 0x02 // 只读
	PAGE_READWRITE         = 0x04 // 读写
	PAGE_WRITECOPY         = 0x08 // 写入时复制
	PAGE_EXECUTE           = 0x10 // 可执行
	PAGE_EXECUTE_READ      = 0x20 // 可执行 + 可读
	PAGE_EXECUTE_READWRITE = 0x40 // 可执行 + 可读 + 可写

	// 快照标志
	TH32CS_SNAPMODULE   = 0x00000008 // 模块快照
	TH32CS_SNAPMODULE32 = 0x00000010 // 32位模块快照

	// 名称长度限制
	MAX_MODULE_NAME32 = 255 // 模块名称最大长度
	MAX_PATH          = 260 // 路径最大长度
)

// ===================== WIN32 结构体定义 =====================

// MEMORY_BASIC_INFORMATION 内存区域信息结构体
type MEMORY_BASIC_INFORMATION struct {
	BaseAddress       uintptr // 区域基址
	AllocationBase    uintptr // 分配基址
	AllocationProtect uint32 // 分配时的保护属性
	RegionSize        uintptr // 区域大小
	State             uint32  // 状态（COMMIT/FREE/RESERVE）
	Protect           uint32  // 当前保护属性
	Type              uint32  // 类型（PRIVATE/MAPPED/IMAGE）
}

// PROCESSENTRY32W 见 C进程.go 中的定义

// MODULEENTRY32W 模块列表条目结构体
type MODULEENTRY32W struct {
	DwSize        uint32                       // 结构体大小（调用前必须设置）
	Th32ModuleID  uint32                       // 模块 ID
	Th32ProcessID uint32                       // 所属进程 ID
	GlblcntUsage  uint32                       // 全局引用计数
	ProccntUsage  uint32                       // 进程引用计数
	ModBaseAddr   uintptr                      // 模块基址（重要）
	ModBaseSize   uint32                       // 模块大小
	HModule       syscall.Handle               // 模块句柄
	SzModule      [MAX_MODULE_NAME32 + 1]uint16 // 模块名称（UTF-16）
	SzExePath     [MAX_PATH]uint16             // 模块完整路径（UTF-16）
}

// ===================== 进程内存读写 =====================

// H内存_打开进程 打开目标进程并获取句柄。
// 参数 进程ID：目标进程 PID
// 参数 权限：进程访问权限（PROCESS_VM_READ | PROCESS_VM_WRITE 等）
// 返回 syscall.Handle：进程句柄（使用完毕后需调用 H内存_关闭句柄）
// 返回 error：失败时返回错误信息
func H内存_打开进程(进程ID uint32, 权限 uint32) (syscall.Handle, error) {
	hProcess, _, err := procOpenProcessH.Call(uintptr(权限), 0, uintptr(进程ID))
	if hProcess == 0 {
		return 0, err
	}
	return syscall.Handle(hProcess), nil
}

// H内存_读整数 从目标进程读取 32 位有符号整数。
// 参数 hProcess：进程句柄
// 参数 地址：目标内存地址
// 返回 int32：读取的整数值
// 返回 error：失败时返回错误信息
func H内存_读整数(hProcess syscall.Handle, 地址 uintptr) (int32, error) {
	var buffer [4]byte
	var numberOfBytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess), 地址,
		uintptr(unsafe.Pointer(&buffer[0])), 4,
		uintptr(unsafe.Pointer(&numberOfBytesRead)),
	)
	if ret == 0 {
		return 0, err
	}
	return int32(buffer[0]) | int32(buffer[1])<<8 | int32(buffer[2])<<16 | int32(buffer[3])<<24, nil
}

// H内存_读整数64 从目标进程读取 64 位有符号整数。
// 参数 hProcess：进程句柄
// 参数 地址：目标内存地址
// 返回 int64：读取的整数值
// 返回 error：失败时返回错误信息
func H内存_读整数64(hProcess syscall.Handle, 地址 uintptr) (int64, error) {
	var buffer [8]byte
	var numberOfBytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess), 地址,
		uintptr(unsafe.Pointer(&buffer[0])), 8,
		uintptr(unsafe.Pointer(&numberOfBytesRead)),
	)
	if ret == 0 {
		return 0, err
	}
	// 小端字节序：逐字节拼接为 int64
	return int64(buffer[0]) | int64(buffer[1])<<8 | int64(buffer[2])<<16 | int64(buffer[3])<<24 |
		int64(buffer[4])<<32 | int64(buffer[5])<<40 | int64(buffer[6])<<48 | int64(buffer[7])<<56, nil
}

// H内存_读字节集 从目标进程读取指定长度的原始字节数据。
// 参数 hProcess：进程句柄
// 参数 地址：目标内存地址
// 参数 长度：要读取的字节数
// 返回 []byte：读取的字节数据
// 返回 error：失败时返回错误信息
func H内存_读字节集(hProcess syscall.Handle, 地址 uintptr, 长度 uint32) ([]byte, error) {
	buffer := make([]byte, 长度)
	var numberOfBytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess), 地址,
		uintptr(unsafe.Pointer(&buffer[0])), uintptr(长度),
		uintptr(unsafe.Pointer(&numberOfBytesRead)),
	)
	if ret == 0 {
		return nil, err
	}
	return buffer[:numberOfBytesRead], nil
}

// H内存_读浮点数 从目标进程读取 32 位浮点数（float32）。
// 参数 hProcess：进程句柄
// 参数 地址：目标内存地址
// 返回 float32：读取的浮点数值
// 返回 error：失败时返回错误信息
func H内存_读浮点数(hProcess syscall.Handle, 地址 uintptr) (float32, error) {
	var buffer [4]byte
	var numberOfBytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess), 地址,
		uintptr(unsafe.Pointer(&buffer[0])), 4,
		uintptr(unsafe.Pointer(&numberOfBytesRead)),
	)
	if ret == 0 {
		return 0, err
	}
	bits := uint32(buffer[0]) | uint32(buffer[1])<<8 | uint32(buffer[2])<<16 | uint32(buffer[3])<<24
	return *(*float32)(unsafe.Pointer(&bits)), nil
}

// H内存_读浮点数64 从目标进程读取 64 位浮点数（float64）。
// 参数 hProcess：进程句柄
// 参数 地址：目标内存地址
// 返回 float64：读取的浮点数值
// 返回 error：失败时返回错误信息
func H内存_读浮点数64(hProcess syscall.Handle, 地址 uintptr) (float64, error) {
	var buffer [8]byte
	var numberOfBytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess), 地址,
		uintptr(unsafe.Pointer(&buffer[0])), 8,
		uintptr(unsafe.Pointer(&numberOfBytesRead)),
	)
	if ret == 0 {
		return 0, err
	}
	bits := uint64(buffer[0]) | uint64(buffer[1])<<8 | uint64(buffer[2])<<16 | uint64(buffer[3])<<24 |
		uint64(buffer[4])<<32 | uint64(buffer[5])<<40 | uint64(buffer[6])<<48 | uint64(buffer[7])<<56
	return *(*float64)(unsafe.Pointer(&bits)), nil
}

// H内存_写整数 向目标进程写入 32 位有符号整数。
// 参数 hProcess：进程句柄
// 参数 地址：目标内存地址
// 参数 值：要写入的整数值
// 返回 error：失败时返回错误信息
func H内存_写整数(hProcess syscall.Handle, 地址 uintptr, 值 int32) error {
	buffer := [4]byte{
		byte(值),
		byte(值 >> 8),
		byte(值 >> 16),
		byte(值 >> 24),
	}
	var numberOfBytesWritten uintptr
	ret, _, err := procWriteProcessMemory.Call(
		uintptr(hProcess), 地址,
		uintptr(unsafe.Pointer(&buffer[0])), 4,
		uintptr(unsafe.Pointer(&numberOfBytesWritten)),
	)
	if ret == 0 {
		return err
	}
	return nil
}

// H内存_写整数64 向目标进程写入 64 位有符号整数。
// 参数 hProcess：进程句柄
// 参数 地址：目标内存地址
// 参数 值：要写入的 64 位整数值
// 返回 error：失败时返回错误信息
func H内存_写整数64(hProcess syscall.Handle, 地址 uintptr, 值 int64) error {
	buffer := [8]byte{
		byte(值), byte(值 >> 8), byte(值 >> 16), byte(值 >> 24),
		byte(值 >> 32), byte(值 >> 40), byte(值 >> 48), byte(值 >> 56),
	}
	var numberOfBytesWritten uintptr
	ret, _, err := procWriteProcessMemory.Call(
		uintptr(hProcess), 地址,
		uintptr(unsafe.Pointer(&buffer[0])), 8,
		uintptr(unsafe.Pointer(&numberOfBytesWritten)),
	)
	if ret == 0 {
		return err
	}
	return nil
}

// H内存_写字节集 向目标进程写入原始字节数据。
// 参数 hProcess：进程句柄
// 参数 地址：目标内存地址
// 参数 数据：要写入的字节数据
// 返回 error：失败时返回错误信息
func H内存_写字节集(hProcess syscall.Handle, 地址 uintptr, 数据 []byte) error {
	var numberOfBytesWritten uintptr
	ret, _, err := procWriteProcessMemory.Call(
		uintptr(hProcess), 地址,
		uintptr(unsafe.Pointer(&数据[0])), uintptr(len(数据)),
		uintptr(unsafe.Pointer(&numberOfBytesWritten)),
	)
	if ret == 0 {
		return err
	}
	return nil
}

// H内存_关闭句柄 关闭进程句柄，释放系统资源。
// 参数 hProcess：进程句柄
// 返回 error：失败时返回错误信息
func H内存_关闭句柄(hProcess syscall.Handle) error {
	ret, _, err := procCloseHandleH.Call(uintptr(hProcess))
	if ret == 0 {
		return err
	}
	return nil
}

// ===================== 内存搜索（AOB 特征码搜索） =====================

// H内存_搜索字节集 在目标进程内存中搜索二进制特征码（AOB 搜索）。
// 特征码中 0xFF 表示通配符（匹配任意字节）。
// 参数 hProcess：进程句柄
// 参数 特征码：要搜索的字节模式（0xFF 为通配符）
// 参数 起始地址：搜索起始内存地址
// 参数 长度：搜索范围大小（字节）
// 返回 []uintptr：所有匹配地址的切片
// 返回 error：失败时返回错误信息
func H内存_搜索字节集(hProcess syscall.Handle, 特征码 []byte, 起始地址 uintptr, 长度 uintptr) ([]uintptr, error) {
	buf := make([]byte, 长度)
	var read uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess), 起始地址,
		uintptr(unsafe.Pointer(&buf[0])), 长度,
		uintptr(unsafe.Pointer(&read)),
	)
	if ret == 0 {
		return nil, err
	}

	var results []uintptr
	patLen := len(特征码)
	for i := uintptr(0); i < read-uintptr(patLen)+1; i++ {
		match := true
		for j := 0; j < patLen; j++ {
			if 特征码[j] != 0xFF && buf[i+uintptr(j)] != 特征码[j] {
				match = false
				break
			}
		}
		if match {
			results = append(results, 起始地址+i)
		}
	}
	return results, nil
}

// H内存_搜索文本 在目标进程内存中搜索文本字符串。
// 参数 hProcess：进程句柄
// 参数 文本：要搜索的文本字符串
// 参数 起始地址：搜索起始内存地址
// 参数 长度：搜索范围大小（字节）
// 返回 []uintptr：所有匹配地址的切片
// 返回 error：失败时返回错误信息
func H内存_搜索文本(hProcess syscall.Handle, 文本 string, 起始地址 uintptr, 长度 uintptr) ([]uintptr, error) {
	return H内存_搜索字节集(hProcess, []byte(文本), 起始地址, 长度)
}

// ===================== 虚拟内存管理 =====================

// H内存_分配内存 在目标进程中分配可执行读写内存。
// 参数 hProcess：进程句柄
// 参数 大小：分配大小（字节）
// 返回 uintptr：分配的内存地址
// 返回 error：失败时返回错误信息
func H内存_分配内存(hProcess syscall.Handle, 大小 uintptr) (uintptr, error) {
	addr, _, err := procVirtualAllocEx.Call(
		uintptr(hProcess), 0, 大小,
		uintptr(uint32(windows.MEM_COMMIT|windows.MEM_RESERVE)),
		uintptr(PAGE_EXECUTE_READWRITE),
	)
	if addr == 0 {
		return 0, err
	}
	return addr, nil
}

// H内存_释放内存 释放目标进程中已分配的内存区域。
// 参数 hProcess：进程句柄
// 参数 地址：要释放的内存地址
// 返回 error：失败时返回错误信息
func H内存_释放内存(hProcess syscall.Handle, 地址 uintptr) error {
	ret, _, err := procVirtualFreeEx.Call(uintptr(hProcess), 地址, 0, uintptr(windows.MEM_RELEASE))
	if ret == 0 {
		return err
	}
	return nil
}

// H内存_修改保护 修改目标进程内存区域的保护属性。
// 参数 hProcess：进程句柄
// 参数 地址：目标内存地址
// 参数 大小：区域大小
// 参数 新保护：新保护属性（PAGE_READWRITE / PAGE_EXECUTE_READWRITE 等）
// 返回 uint32：修改前的旧保护属性
// 返回 error：失败时返回错误信息
func H内存_修改保护(hProcess syscall.Handle, 地址 uintptr, 大小 uintptr, 新保护 uint32) (uint32, error) {
	var oldProtect uint32
	ret, _, err := procVirtualProtectEx.Call(
		uintptr(hProcess), 地址, 大小,
		uintptr(新保护),
		uintptr(unsafe.Pointer(&oldProtect)),
	)
	if ret == 0 {
		return 0, err
	}
	return oldProtect, nil
}

// H内存_查询内存 查询目标进程指定地址的内存信息。
// 参数 hProcess：进程句柄
// 参数 地址：要查询的内存地址
// 返回 *MEMORY_BASIC_INFORMATION：内存区域信息（基址、大小、保护属性等）
// 返回 error：失败时返回错误信息
func H内存_查询内存(hProcess syscall.Handle, 地址 uintptr) (*MEMORY_BASIC_INFORMATION, error) {
	var mbi MEMORY_BASIC_INFORMATION
	ret, _, err := procVirtualQueryEx.Call(
		uintptr(hProcess), 地址,
		uintptr(unsafe.Pointer(&mbi)),
		unsafe.Sizeof(mbi),
	)
	if ret == 0 {
		return nil, err
	}
	return &mbi, nil
}

// ===================== 进程枚举与管理 =====================

// H内存_取进程ID 通过进程名称获取进程 ID。
// 参数 进程名：进程可执行文件名（如 "notepad.exe"）
// 返回 uint32：进程 ID
// 返回 error：未找到时返回错误
func H内存_取进程ID(进程名 string) (uint32, error) {
	snapshot, _, err := procCreateToolhelp32Snapshot.Call(uintptr(TH32CS_SNAPPROCESS), 0)
	if snapshot == uintptr(syscall.InvalidHandle) {
		return 0, err
	}
	defer procCloseHandleH.Call(snapshot)

	var pe32 PROCESSENTRY32W
	pe32.Size = uint32(unsafe.Sizeof(pe32))

	ret, _, _ := procProcess32FirstW.Call(snapshot, uintptr(unsafe.Pointer(&pe32)))
	if ret == 0 {
		return 0, errors.New("無法枚舉進程")
	}

	for {
		name := syscall.UTF16ToString(pe32.ExeFile[:])
		if name == 进程名 {
			return pe32.ProcessID, nil
		}
		ret, _, _ = procProcess32NextW.Call(snapshot, uintptr(unsafe.Pointer(&pe32)))
		if ret == 0 {
			break
		}
	}
	return 0, errors.New("未找到指定進程: " + 进程名)
}

// H内存_枚举进程 枚举系统中所有正在运行的进程。
// 返回 []PROCESSENTRY32W：进程信息列表
// 返回 error：失败时返回错误
func H内存_枚举进程() ([]PROCESSENTRY32W, error) {
	snapshot, _, err := procCreateToolhelp32Snapshot.Call(uintptr(TH32CS_SNAPPROCESS), 0)
	if snapshot == uintptr(syscall.InvalidHandle) {
		return nil, err
	}
	defer procCloseHandleH.Call(snapshot)

	var pe32 PROCESSENTRY32W
	pe32.Size = uint32(unsafe.Sizeof(pe32))

	var result []PROCESSENTRY32W
	ret, _, _ := procProcess32FirstW.Call(snapshot, uintptr(unsafe.Pointer(&pe32)))
	if ret == 0 {
		return nil, errors.New("無法枚舉進程")
	}

	for {
		result = append(result, pe32)
		ret, _, _ = procProcess32NextW.Call(snapshot, uintptr(unsafe.Pointer(&pe32)))
		if ret == 0 {
			break
		}
	}
	return result, nil
}

// H内存_取模块基址 获取目标进程中指定模块的加载基址。
// 参数 进程ID：目标进程 ID
// 参数 模块名：模块文件名（如 "kernel32.dll"）
// 返回 uintptr：模块基址
// 返回 error：未找到时返回错误
func H内存_取模块基址(进程ID uint32, 模块名 string) (uintptr, error) {
	snapshot, _, err := procCreateToolhelp32Snapshot.Call(
		uintptr(TH32CS_SNAPMODULE|TH32CS_SNAPMODULE32),
		uintptr(进程ID),
	)
	if snapshot == uintptr(syscall.InvalidHandle) {
		return 0, err
	}
	defer procCloseHandleH.Call(snapshot)

	var me32 MODULEENTRY32W
	me32.DwSize = uint32(unsafe.Sizeof(me32))

	ret, _, _ := procModule32FirstW.Call(snapshot, uintptr(unsafe.Pointer(&me32)))
	if ret == 0 {
		return 0, errors.New("無法枚舉模塊")
	}

	// 大小写不敏感匹配
	want, _ := syscall.UTF16FromString(模块名)
	for {
		name := me32.SzModule[:]
		match := true
		for i := 0; i < len(want) && want[i] != 0; i++ {
			c1, c2 := want[i], name[i]
			if c1 >= 'A' && c1 <= 'Z' {
				c1 += 32
			}
			if c2 >= 'A' && c2 <= 'Z' {
				c2 += 32
			}
			if c1 != c2 {
				match = false
				break
			}
		}
		if match {
			return me32.ModBaseAddr, nil
		}
		ret, _, _ = procModule32NextW.Call(snapshot, uintptr(unsafe.Pointer(&me32)))
		if ret == 0 {
			break
		}
	}
	return 0, errors.New("未找到模塊: " + 模块名)
}

// H内存_终止进程 强制终止指定进程。
// 参数 进程ID：目标进程 ID
// 返回 error：失败时返回错误信息
func H内存_终止进程(进程ID uint32) error {
	hProcess, err := H内存_打开进程(进程ID, PROCESS_TERMINATE)
	if err != nil {
		return err
	}
	defer H内存_关闭句柄(hProcess)

	ret, _, err := procTerminateProcessH.Call(uintptr(hProcess), 0)
	if ret == 0 {
		return err
	}
	return nil
}

// H内存_取进程名称 获取指定进程 ID 对应的可执行文件名称。
// 参数 进程ID：目标进程 ID
// 返回 string：进程名称
// 返回 []uint16：原始 UTF-16 名称数组
// 返回 error：失败时返回错误
func H内存_取进程名称(进程ID uint32) (string, []uint16, error) {
	snapshot, _, err := procCreateToolhelp32Snapshot.Call(uintptr(TH32CS_SNAPPROCESS), 0)
	if snapshot == uintptr(syscall.InvalidHandle) {
		return "", nil, err
	}
	defer procCloseHandleH.Call(snapshot)

	var pe32 PROCESSENTRY32W
	pe32.Size = uint32(unsafe.Sizeof(pe32))

	ret, _, _ := procProcess32FirstW.Call(snapshot, uintptr(unsafe.Pointer(&pe32)))
	if ret == 0 {
		return "", nil, errors.New("無法枚舉進程")
	}

	for {
		if pe32.ProcessID == 进程ID {
			return syscall.UTF16ToString(pe32.ExeFile[:]), pe32.ExeFile[:], nil
		}
		ret, _, _ = procProcess32NextW.Call(snapshot, uintptr(unsafe.Pointer(&pe32)))
		if ret == 0 {
			break
		}
	}
	return "", nil, errors.New("未找到指定進程")
}

// H内存_是否存在进程 检查指定进程是否正在运行。
// 参数 进程ID：目标进程 ID
// 返回 bool：true 表示进程存在且正在运行（退出码为 STILL_ACTIVE=259）
func H内存_是否存在进程(进程ID uint32) bool {
	hProcess, err := H内存_打开进程(进程ID, PROCESS_QUERY_INFORMATION)
	if err != nil {
		return false
	}
	defer H内存_关闭句柄(hProcess)

	var exitCode uint32
	ret, _, _ := procGetExitCodeProcessH.Call(uintptr(hProcess), uintptr(unsafe.Pointer(&exitCode)))
	if ret == 0 {
		return false
	}
	return exitCode == 259 // STILL_ACTIVE
}

// H内存_取进程位数 检测目标进程是 32 位还是 64 位。
// 参数 进程ID：目标进程 ID
// 返回 int：32 或 64，失败时返回 0
func H内存_取进程位数(进程ID uint32) int {
	handle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION, false, 进程ID)
	if err != nil {
		return 0
	}
	defer windows.CloseHandle(handle)

	var isWow64 bool
	err = windows.IsWow64Process(handle, &isWow64)
	if err != nil {
		return 0
	}
	if isWow64 {
		return 32 // 32 位进程跑在 64 位系统上
	}
	return 64 // 原生 64 位进程
}