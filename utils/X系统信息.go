package utils

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
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
