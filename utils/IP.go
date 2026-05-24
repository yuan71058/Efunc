package utils

import (
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// IP_10进制转IP 将 10 进制整数转换为点分十进制 IP 地址。
// 整数按大端序拆分为 4 个字节，每个字节对应 IP 的一段。
//
// 参数:
//   - decimal: 10 进制表示的 IP 地址整数
//
// 返回:
//   - string: 点分十进制格式的 IP 地址
//
// 示例:
//
//	IP_10进制转IP(3232235777)  // "192.168.1.1"
func IP_10进制转IP(decimal int) string {
	ip := make([]string, 4)
	for i := 3; i >= 0; i-- {
		ip[i] = strconv.Itoa(decimal & 255)
		decimal >>= 8
	}
	return strings.Join(ip, ".")
}

// IP_取内网IP 获取本机所有内网（局域网）IPv4 地址。
// 遍历所有网络接口，过滤回环地址和 IPv6 地址，返回可用的内网 IP 列表。
//
// 返回:
//   - []string: 内网 IPv4 地址列表；无可用地址时返回空切片
func IP_取内网IP() []string {
	var 结果 []string
	接口列表, err := net.Interfaces()
	if err != nil {
		return 结果
	}
	for _, 接口 := range 接口列表 {
		地址列表, err := 接口.Addrs()
		if err != nil {
			continue
		}
		for _, 地址 := range 地址列表 {
			ipNet, ok := 地址.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() || ipNet.IP.To4() == nil {
				continue
			}
			结果 = append(结果, ipNet.IP.String())
		}
	}
	return 结果
}

// IP_取首选内网IP 获取本机首选内网 IPv4 地址。
// 优先返回非 "169.254.x.x"（链路本地）的地址，通常为实际局域网地址。
//
// 返回:
//   - string: 首选内网 IPv4 地址；无可用地址返回空字符串
func IP_取首选内网IP() string {
	地址列表 := IP_取内网IP()
	if len(地址列表) == 0 {
		return ""
	}
	for _, ip := range 地址列表 {
		if !strings.HasPrefix(ip, "169.254.") {
			return ip
		}
	}
	return 地址列表[0]
}

// IP_取外网IP 通过公共 IP 查询服务获取本机的外网（公网）IPv4 地址。
// 依次尝试多个查询服务，任一成功即返回。
//
// 返回:
//   - string: 外网 IPv4 地址；获取失败返回空字符串
func IP_取外网IP() string {
	查询地址 := []string{
		"https://api.ipify.org",
		"https://ifconfig.me/ip",
		"https://icanhazip.com",
		"https://checkip.amazonaws.com",
	}
	for _, url := range 查询地址 {
		resp, err := http.DefaultClient.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			continue
		}
		buf := make([]byte, 64)
		n, err := resp.Body.Read(buf)
		if err != nil && n == 0 {
			continue
		}
		ip := strings.TrimSpace(string(buf[:n]))
		if net.ParseIP(ip) != nil {
			return ip
		}
	}
	return ""
}

// IP_取外网IP详细信息 通过 ip-api.com 获取外网 IP 的详细信息，包括地理位置、ISP 等。
// 返回 JSON 格式的信息字符串。
//
// 返回:
//   - string: JSON 格式的 IP 详细信息；获取失败返回空字符串
func IP_取外网IP详细信息() string {
	resp, err := http.DefaultClient.Get("http://ip-api.com/json/?lang=zh-CN")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return ""
	}
	buf := make([]byte, 4096)
	n, err := resp.Body.Read(buf)
	if err != nil && n == 0 {
		return ""
	}
	return strings.TrimSpace(string(buf[:n]))
}

// IP_IP转10进制 将点分十进制 IPv4 地址转换为 10 进制整数。
//
// 参数:
//   - ip: 点分十进制格式的 IPv4 地址，如 "192.168.1.1"
//
// 返回:
//   - int: 10 进制表示的 IP 地址整数；格式无效返回 0
//
// 示例:
//
//	IP_IP转10进制("192.168.1.1")  // 3232235777
func IP_IP转10进制(ip string) int {
	解析ip := net.ParseIP(ip)
	if 解析ip == nil {
		return 0
	}
	ipv4 := 解析ip.To4()
	if ipv4 == nil {
		return 0
	}
	result := 0
	for _, b := range ipv4 {
		result = result<<8 | int(b)
	}
	return result
}

// IP_是否内网IP 判断指定的 IPv4 地址是否为内网（私有）地址。
// 内网地址范围：10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16。
//
// 参数:
//   - ip: 点分十进制格式的 IPv4 地址
//
// 返回:
//   - bool: 是内网地址返回 true，否则返回 false
func IP_是否内网IP(ip string) bool {
	解析ip := net.ParseIP(ip)
	if 解析ip == nil {
		return false
	}
	私有网段 := []struct {
		网络号   net.IP
		掩码长度 net.IPMask
	}{
		{net.IPv4(10, 0, 0, 0), net.CIDRMask(8, 32)},
		{net.IPv4(172, 16, 0, 0), net.CIDRMask(12, 32)},
		{net.IPv4(192, 168, 0, 0), net.CIDRMask(16, 32)},
	}
	for _, 网段 := range 私有网段 {
		if 解析ip.Mask(网段.掩码长度).Equal(网段.网络号.Mask(网段.掩码长度)) {
			return true
		}
	}
	return false
}

// IP_是否有效IP 判断指定字符串是否为有效的 IPv4 或 IPv6 地址。
//
// 参数:
//   - ip: 待验证的 IP 地址字符串
//
// 返回:
//   - bool: 有效返回 true，否则返回 false
func IP_是否有效IP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IP_取MAC地址 获取本机所有网络接口的 MAC 地址。
//
// 返回:
//   - map[string]string: 键为接口名，值为 MAC 地址（小写冒号分隔）；获取失败返回空 map
func IP_取MAC地址() map[string]string {
	结果 := make(map[string]string)
	接口列表, err := net.Interfaces()
	if err != nil {
		return 结果
	}
	for _, 接口 := range 接口列表 {
		if 接口.HardwareAddr != nil {
			结果[接口.Name] = 接口.HardwareAddr.String()
		}
	}
	return 结果
}

// IP_Ping测试 测试与指定主机的 TCP 连通性（非 ICMP ping）。
// 尝试在指定超时时间内建立 TCP 连接到目标主机的指定端口。
//
// 参数:
//   - 主机: 目标主机地址（IP 或域名）
//   - 端口: 目标端口号
//   - 超时毫秒: 连接超时时间（毫秒）
//
// 返回:
//   - bool: 连接成功返回 true，否则返回 false
func IP_Ping测试(主机 string, 端口 int, 超时毫秒 int) bool {
	地址 := net.JoinHostPort(主机, strconv.Itoa(端口))
	conn, err := net.DialTimeout("tcp", 地址, time.Duration(超时毫秒)*time.Millisecond)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
