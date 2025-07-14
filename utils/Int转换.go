package utils

import (
	"bytes"
	"encoding/binary"
)

// 防止精度丢失
func Int取绝对值(值 int) int {
	if 值 >= 0 {
		return 值
	}
	return -值
}

// 32位int 转换成字节数组 大端,长度4字节
func Int32ToBytes(i int32) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, i)
	if err != nil {
		return []byte{}
	}
	return buf.Bytes()
}
