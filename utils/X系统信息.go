package utils

import (
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

// X系统_取CPU信息 获取 CPU 信息，包括核心数、型号等。
//
// 返回:
//   - []cpu.InfoStat: CPU 信息列表（多 CPU 时有多个元素）
//   - error: 获取失败时返回错误
func X系统_取CPU信息() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

// X系统_取CPU核心数 获取 CPU 逻辑核心数。
//
// 返回:
//   - int: 逻辑核心数，如 8
//   - error: 获取失败时返回错误
func X系统_取CPU核心数() (int, error) {
	return cpu.Counts(true)
}

// X系统_取CPU物理核心数 获取 CPU 物理核心数。
//
// 返回:
//   - int: 物理核心数
//   - error: 获取失败时返回错误
func X系统_取CPU物理核心数() (int, error) {
	return cpu.Counts(false)
}

// X系统_取CPU使用率 获取 CPU 使用率百分比。
//
// 参数:
//   - 间隔: 采样间隔（秒），0 表示瞬时值
//
// 返回:
//   - []float64: 各核心的使用率（0-100）
//   - error: 获取失败时返回错误
func X系统_取CPU使用率(间隔 int) ([]float64, error) {
	return cpu.Percent(time.Duration(间隔)*time.Second, true)
}

// X系统_取总CPU使用率 获取总体 CPU 使用率百分比。
//
// 参数:
//   - 间隔: 采样间隔（秒），0 表示瞬时值
//
// 返回:
//   - float64: 总体 CPU 使用率（0-100）
//   - error: 获取失败时返回错误
func X系统_取总CPU使用率(间隔 int) (float64, error) {
	percentages, err := cpu.Percent(time.Duration(间隔)*time.Second, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) == 0 {
		return 0, nil
	}
	return percentages[0], nil
}

// X系统_取内存信息 获取内存使用情况，包括总量、已用、可用等。
//
// 返回:
//   - *mem.VirtualMemoryStat: 内存信息
//   - error: 获取失败时返回错误
func X系统_取内存信息() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

// X系统_取交换区信息 获取交换区（Swap）使用情况。
//
// 返回:
//   - *mem.SwapMemoryStat: 交换区信息
//   - error: 获取失败时返回错误
func X系统_取交换区信息() (*mem.SwapMemoryStat, error) {
	return mem.SwapMemory()
}

// X系统_取磁盘信息 获取磁盘分区使用情况。
//
// 返回:
//   - []disk.PartitionStat: 磁盘分区列表
//   - error: 获取失败时返回错误
func X系统_取磁盘信息() ([]disk.PartitionStat, error) {
	return disk.Partitions(true)
}

// X系统_取磁盘使用量 获取指定路径的磁盘使用量。
//
// 参数:
//   - 路径: 磁盘路径，如 "/" 或 "C:\\"
//
// 返回:
//   - *disk.UsageStat: 磁盘使用信息
//   - error: 获取失败时返回错误
func X系统_取磁盘使用量(路径 string) (*disk.UsageStat, error) {
	return disk.Usage(路径)
}

// X系统_取磁盘IO信息 获取磁盘 IO 统计信息（读写字节数、读写次数等）。
//
// 返回:
//   - map[string]disk.IOCountersStat: 键为磁盘设备名，值为 IO 统计
//   - error: 获取失败时返回错误
func X系统_取磁盘IO信息() (map[string]disk.IOCountersStat, error) {
	return disk.IOCounters()
}

// X系统_取主机信息 获取主机信息，包括主机名、操作系统、内核版本等。
//
// 返回:
//   - *host.InfoStat: 主机信息
//   - error: 获取失败时返回错误
func X系统_取主机信息() (*host.InfoStat, error) {
	return host.Info()
}

// X系统_取网络接口信息 获取网络接口列表。
//
// 返回:
//   - []net.InterfaceStat: 网络接口信息列表
//   - error: 获取失败时返回错误
func X系统_取网络接口信息() ([]net.InterfaceStat, error) {
	return net.Interfaces()
}

// X系统_取网络连接信息 获取当前活动的网络连接列表。
//
// 返回:
//   - []net.ConnectionStat: 网络连接信息列表
//   - error: 获取失败时返回错误
func X系统_取网络连接信息() ([]net.ConnectionStat, error) {
	return net.Connections("all")
}

// X系统_取网络IO信息 获取网络接口的 IO 统计信息（收发字节数、包数等）。
//
// 返回:
//   - []net.IOCountersStat: 网络 IO 统计列表
//   - error: 获取失败时返回错误
func X系统_取网络IO信息() ([]net.IOCountersStat, error) {
	return net.IOCounters(true)
}

// X系统_取开机时间 获取系统开机时长（秒）。
//
// 返回:
//   - uint64: 开机时长（秒）
//   - error: 获取失败时返回错误
func X系统_取开机时间() (uint64, error) {
	信息, 错误 := host.Info()
	if 错误 != nil {
		return 0, 错误
	}
	return 信息.Uptime, nil
}

// X系统_取系统负载 获取系统平均负载（1分钟、5分钟、15分钟）。
// Linux/macOS 下返回真实负载值，Windows 下通过进程数模拟。
//
// 返回:
//   - *load.AvgStat: 系统负载信息
//   - error: 获取失败时返回错误
func X系统_取系统负载() (*load.AvgStat, error) {
	return load.Avg()
}

// X系统_取进程列表 获取当前系统中所有正在运行的进程列表。
//
// 返回:
//   - []process.Process: 进程对象列表
//   - error: 获取失败时返回错误
func X系统_取进程列表() ([]*process.Process, error) {
	return process.Processes()
}

// X系统_取进程信息 根据进程 PID 获取进程详细信息。
// 包括进程名、状态、CPU/内存占用、创建时间等。
//
// 参数:
//   - pid: 进程 ID
//
// 返回:
//   - *process.Process: 进程对象；进程不存在返回 nil
//   - error: 获取失败时返回错误
func X系统_取进程信息(pid int32) (*process.Process, error) {
	return process.NewProcess(pid)
}

// X系统_取当前进程ID 获取当前进程的 PID。
//
// 返回:
//   - int32: 当前进程 ID
func X系统_取当前进程ID() int32 {
	return int32(os.Getpid())
}

// X系统_取当前进程信息 获取当前进程的详细信息。
//
// 返回:
//   - *process.Process: 当前进程对象
//   - error: 获取失败时返回错误
func X系统_取当前进程信息() (*process.Process, error) {
	return process.NewProcess(X系统_取当前进程ID())
}

// X系统_取进程名 根据进程 PID 获取进程名称。
//
// 参数:
//   - pid: 进程 ID
//
// 返回:
//   - string: 进程名称；获取失败返回空字符串
func X系统_取进程名(pid int32) string {
	p, err := process.NewProcess(pid)
	if err != nil {
		return ""
	}
	name, err := p.Name()
	if err != nil {
		return ""
	}
	return name
}

// X系统_取进程内存占用 根据进程 PID 获取进程的内存占用（字节）。
//
// 参数:
//   - pid: 进程 ID
//
// 返回:
//   - float64: 内存占用（字节）；获取失败返回 0
func X系统_取进程内存占用(pid int32) float64 {
	p, err := process.NewProcess(pid)
	if err != nil {
		return 0
	}
	mem, err := p.MemoryInfo()
	if err != nil {
		return 0
	}
	return float64(mem.RSS)
}

// X系统_取进程CPU占用 根据进程 PID 获取进程的 CPU 使用率百分比。
//
// 参数:
//   - pid: 进程 ID
//
// 返回:
//   - float64: CPU 使用率（0-100）；获取失败返回 0
func X系统_取进程CPU占用(pid int32) float64 {
	p, err := process.NewProcess(pid)
	if err != nil {
		return 0
	}
	cpuPercent, err := p.CPUPercent()
	if err != nil {
		return 0
	}
	return cpuPercent
}

// X系统_是否64位系统 判断当前操作系统是否为 64 位架构。
//
// 返回:
//   - bool: 64 位返回 true，32 位返回 false
func X系统_是否64位系统() bool {
	return runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64" || runtime.GOARCH == "loong64"
}

// X系统_取系统架构 获取当前系统的 CPU 架构名称。
// 如 "amd64"、"arm64"、"386" 等。
//
// 返回:
//   - string: CPU 架构名称
func X系统_取系统架构() string {
	return runtime.GOARCH
}

// X系统_取操作系统类型 获取当前操作系统名称。
// 如 "windows"、"linux"、"darwin" 等。
//
// 返回:
//   - string: 操作系统名称
func X系统_取操作系统类型() string {
	return runtime.GOOS
}

// X系统_取逻辑处理器数 获取 Go 运行时可用的逻辑处理器数量。
// 等同于 GOMAXPROCS 的默认值。
//
// 返回:
//   - int: 逻辑处理器数量
func X系统_取逻辑处理器数() int {
	return runtime.NumCPU()
}

// X系统_取Go版本 获取当前 Go 编译器版本。
//
// 返回:
//   - string: Go 版本字符串，如 "go1.22.0"
func X系统_取Go版本() string {
	return runtime.Version()
}
