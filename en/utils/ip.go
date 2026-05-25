package utils

import (
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// IP_DecToIP 将 10 进制整数转换为点分十进制 IP 地址。
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
//	IP_DecToIP(3232235777)  // "192.168.1.1"
func IP_DecToIP(decimal int) string {
	ip := make([]string, 4)
	for i := 3; i >= 0; i-- {
		ip[i] = strconv.Itoa(decimal & 255)
		decimal >>= 8
	}
	return strings.Join(ip, ".")
}

// IP_GetLocalIPs 获取本机所有内网（局域网）IPv4 地址。
// 遍历所有网络接口，过滤回环地址和 IPv6 地址，返回可用的内网 IP 列表。
//
// 返回:
//   - []string: 内网 IPv4 地址列表；无可用地址时返回空切片
func IP_GetLocalIPs() []string {
	var result []string
	interfaces, err := net.Interfaces()
	if err != nil {
		return result
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() || ipNet.IP.To4() == nil {
				continue
			}
			result = append(result, ipNet.IP.String())
		}
	}
	return result
}

// IP_GetPreferredLocalIP 获取本机首选内网 IPv4 地址。
// 优先返回非 "169.254.x.x"（链路本地）的地址，通常为实际局域网地址。
//
// 返回:
//   - string: 首选内网 IPv4 地址；无可用地址返回空字符串
func IP_GetPreferredLocalIP() string {
	ips := IP_GetLocalIPs()
	if len(ips) == 0 {
		return ""
	}
	for _, ip := range ips {
		if !strings.HasPrefix(ip, "169.254.") {
			return ip
		}
	}
	return ips[0]
}

// IP_GetPublicIP 通过公共 IP 查询服务获取本机的外网（公网）IPv4 地址。
// 依次尝试多个查询服务，任一成功即返回。
//
// 返回:
//   - string: 外网 IPv4 地址；获取失败返回空字符串
func IP_GetPublicIP() string {
	queryURLs := []string{
		"https://api.ipify.org",
		"https://ifconfig.me/ip",
		"https://icanhazip.com",
		"https://checkip.amazonaws.com",
	}
	for _, url := range queryURLs {
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

// IP_GetPublicIPDetail 通过 ip-api.com 获取外网 IP 的详细信息，包括地理位置、ISP 等。
// 返回 JSON 格式的信息字符串。
//
// 返回:
//   - string: JSON 格式的 IP 详细信息；获取失败返回空字符串
func IP_GetPublicIPDetail() string {
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

// IP_IPToDec 将点分十进制 IPv4 地址转换为 10 进制整数。
//
// 参数:
//   - ip: 点分十进制格式的 IPv4 地址，如 "192.168.1.1"
//
// 返回:
//   - int: 10 进制表示的 IP 地址整数；格式无效返回 0
//
// 示例:
//
//	IP_IPToDec("192.168.1.1")  // 3232235777
func IP_IPToDec(ip string) int {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return 0
	}
	ipv4 := parsedIP.To4()
	if ipv4 == nil {
		return 0
	}
	result := 0
	for _, b := range ipv4 {
		result = result<<8 | int(b)
	}
	return result
}

// IP_IsPrivate 判断指定的 IPv4 地址是否为内网（私有）地址。
// 内网地址范围：10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16。
//
// 参数:
//   - ip: 点分十进制格式的 IPv4 地址
//
// 返回:
//   - bool: 是内网地址返回 true，否则返回 false
func IP_IsPrivate(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	privateRanges := []struct {
		network net.IP
		mask    net.IPMask
	}{
		{net.IPv4(10, 0, 0, 0), net.CIDRMask(8, 32)},
		{net.IPv4(172, 16, 0, 0), net.CIDRMask(12, 32)},
		{net.IPv4(192, 168, 0, 0), net.CIDRMask(16, 32)},
	}
	for _, network := range privateRanges {
		if parsedIP.Mask(network.mask).Equal(network.network.Mask(network.mask)) {
			return true
		}
	}
	return false
}

// IP_IsValid 判断指定字符串是否为有效的 IPv4 或 IPv6 地址。
//
// 参数:
//   - ip: 待验证的 IP 地址字符串
//
// 返回:
//   - bool: 有效返回 true，否则返回 false
func IP_IsValid(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IP_GetMACAddresses 获取本机所有网络接口的 MAC 地址。
//
// 返回:
//   - map[string]string: 键为接口名，值为 MAC 地址（小写冒号分隔）；获取失败返回空 map
func IP_GetMACAddresses() map[string]string {
	result := make(map[string]string)
	interfaces, err := net.Interfaces()
	if err != nil {
		return result
	}
	for _, iface := range interfaces {
		if iface.HardwareAddr != nil {
			result[iface.Name] = iface.HardwareAddr.String()
		}
	}
	return result
}

// IP_Ping 测试与指定主机的 TCP 连通性（非 ICMP ping）。
// 尝试在指定超时时间内建立 TCP 连接到目标主机的指定端口。
//
// 参数:
//   - host: 目标主机地址（IP 或域名）
//   - port: 目标端口号
//   - timeoutMs: 连接超时时间（毫秒）
//
// 返回:
//   - bool: 连接成功返回 true，否则返回 false
func IP_Ping(host string, port int, timeoutMs int) bool {
	addr := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", addr, time.Duration(timeoutMs)*time.Millisecond)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}