package utils

import (
	"bytes"
	"encoding/hex"
)

func Z字节集_十六进制到字节集(原始16进制文本 string) []byte {
	字节集, _ := hex.DecodeString(原始16进制文本)
	return 字节集
}

func Z字节集_字节集到十六进制(字节集 []byte) string {
	return hex.EncodeToString(字节集)
}
func Z字节集_寻找(被搜寻的字节集 []byte, 欲寻找的字节集 []byte, 起始搜寻位置 int) int {
	if len(被搜寻的字节集) == 0 || len(欲寻找的字节集) == 0 {
		return -1
	}

	if 起始搜寻位置 <= 0 || 起始搜寻位置 > len(被搜寻的字节集) {
		起始搜寻位置 = 1
	}

	for i := 起始搜寻位置 - 1; i < len(被搜寻的字节集); i++ {
		if bytes.HasPrefix(被搜寻的字节集[i:], 欲寻找的字节集) {
			return i + 1
		}
	}

	return -1
}
