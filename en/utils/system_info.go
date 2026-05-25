// 系统信息监控工具
// 基于 shirou/gopsutil/v3 库，提供 CPU、内存、磁盘、网络、进程等系统信息监控。
// 跨平台支持 Windows/Linux/macOS。
package utils

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

// SystemInfo_CPUInfo 获取 CPU 信息，包括核心数、型号等。
//
// 返回:
//   - []cpu.InfoStat: CPU 信息列表（多 CPU 时有多个元素）
//   - error: 获取失败时返回错误
func SystemInfo_CPUInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

// SystemInfo_CPULogicalCount 获取 CPU 逻辑核心数。
//
// 返回:
//   - int: 逻辑核心数，如 8
//   - error: 获取失败时返回错误
func SystemInfo_CPULogicalCount() (int, error) {
	return cpu.Counts(true)
}

// SystemInfo_CPUPhysCount 获取 CPU 物理核心数。
//
// 返回:
//   - int: 物理核心数
//   - error: 获取失败时返回错误
func SystemInfo_CPUPhysCount() (int, error) {
	return cpu.Counts(false)
}

// SystemInfo_CPUUsage 获取 CPU 使用率百分比。
//
// 参数:
//   - interval: 采样间隔（秒），0 表示瞬时值
//
// 返回:
//   - []float64: 各核心的使用率（0-100）
//   - error: 获取失败时返回错误
func SystemInfo_CPUUsage(interval int) ([]float64, error) {
	return cpu.Percent(time.Duration(interval)*time.Second, true)
}

// SystemInfo_CPUTotalUsage 获取总体 CPU 使用率百分比。
//
// 参数:
//   - interval: 采样间隔（秒），0 表示瞬时值
//
// 返回:
//   - float64: 总体 CPU 使用率（0-100）
//   - error: 获取失败时返回错误
func SystemInfo_CPUTotalUsage(interval int) (float64, error) {
	percentages, err := cpu.Percent(time.Duration(interval)*time.Second, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) == 0 {
		return 0, nil
	}
	return percentages[0], nil
}

// SystemInfo_MemInfo 获取内存使用情况，包括总量、已用、可用等。
//
// 返回:
//   - *mem.VirtualMemoryStat: 内存信息
//   - error: 获取失败时返回错误
func SystemInfo_MemInfo() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

// SystemInfo_SwapInfo 获取交换区（Swap）使用情况。
//
// 返回:
//   - *mem.SwapMemoryStat: 交换区信息
//   - error: 获取失败时返回错误
func SystemInfo_SwapInfo() (*mem.SwapMemoryStat, error) {
	return mem.SwapMemory()
}

// SystemInfo_DiskPartitions 获取磁盘分区使用情况。
//
// 返回:
//   - []disk.PartitionStat: 磁盘分区列表
//   - error: 获取失败时返回错误
func SystemInfo_DiskPartitions() ([]disk.PartitionStat, error) {
	return disk.Partitions(true)
}

// SystemInfo_DiskUsage 获取指定路径的磁盘使用量。
//
// 参数:
//   - path: 磁盘路径，如 "/" 或 "C:\\"
//
// 返回:
//   - *disk.UsageStat: 磁盘使用信息
//   - error: 获取失败时返回错误
func SystemInfo_DiskUsage(path string) (*disk.UsageStat, error) {
	return disk.Usage(path)
}

// SystemInfo_DiskIO 获取磁盘 IO 统计信息（读写字节数、读写次数等）。
//
// 返回:
//   - map[string]disk.IOCountersStat: 键为磁盘设备名，值为 IO 统计
//   - error: 获取失败时返回错误
func SystemInfo_DiskIO() (map[string]disk.IOCountersStat, error) {
	return disk.IOCounters()
}

// SystemInfo_HostInfo 获取主机信息，包括主机名、操作系统、内核版本等。
//
// 返回:
//   - *host.InfoStat: 主机信息
//   - error: 获取失败时返回错误
func SystemInfo_HostInfo() (*host.InfoStat, error) {
	return host.Info()
}

// SystemInfo_NetInterfaces 获取网络接口列表。
//
// 返回:
//   - []net.InterfaceStat: 网络接口信息列表
//   - error: 获取失败时返回错误
func SystemInfo_NetInterfaces() ([]net.InterfaceStat, error) {
	return net.Interfaces()
}

// SystemInfo_NetConnections 获取当前活动的网络连接列表。
//
// 返回:
//   - []net.ConnectionStat: 网络连接信息列表
//   - error: 获取失败时返回错误
func SystemInfo_NetConnections() ([]net.ConnectionStat, error) {
	return net.Connections("all")
}

// SystemInfo_NetIO 获取网络接口的 IO 统计信息（收发字节数、包数等）。
//
// 返回:
//   - []net.IOCountersStat: 网络 IO 统计列表
//   - error: 获取失败时返回错误
func SystemInfo_NetIO() ([]net.IOCountersStat, error) {
	return net.IOCounters(true)
}

// SystemInfo_Uptime 获取系统开机时长（秒）。
//
// 返回:
//   - uint64: 开机时长（秒）
//   - error: 获取失败时返回错误
func SystemInfo_Uptime() (uint64, error) {
	info, err := host.Info()
	if err != nil {
		return 0, err
	}
	return info.Uptime, nil
}

// SystemInfo_LoadAvg 获取系统平均负载（1分钟、5分钟、15分钟）。
// Linux/macOS 下返回真实负载值，Windows 下通过进程数模拟。
//
// 返回:
//   - *load.AvgStat: 系统负载信息
//   - error: 获取失败时返回错误
func SystemInfo_LoadAvg() (*load.AvgStat, error) {
	return load.Avg()
}

// SystemInfo_Processes 获取当前系统中所有正在运行的进程列表。
//
// 返回:
//   - []*process.Process: 进程对象列表
//   - error: 获取失败时返回错误
func SystemInfo_Processes() ([]*process.Process, error) {
	return process.Processes()
}

// SystemInfo_ProcessByPID 根据进程 PID 获取进程详细信息。
// 包括进程名、状态、CPU/内存占用、创建时间等。
//
// 参数:
//   - pid: 进程 ID
//
// 返回:
//   - *process.Process: 进程对象；进程不存在返回 nil
//   - error: 获取失败时返回错误
func SystemInfo_ProcessByPID(pid int32) (*process.Process, error) {
	return process.NewProcess(pid)
}