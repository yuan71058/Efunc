package utils

import (
	"strconv"
	"strings"
)

func IP_10进制转IP(decimal int) string {
	ip := make([]string, 4)
	for i := 3; i >= 0; i-- {
		ip[i] = strconv.Itoa(decimal & 255)
		decimal >>= 8
	}
	return strings.Join(ip, ".")
}
