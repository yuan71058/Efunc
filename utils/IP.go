package utils

import (
	"strconv"
	"strings"
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
