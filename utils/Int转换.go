package utils

import (
	"bytes"
	"encoding/binary"
)

// Int取绝对值 返回整数的绝对值。
// 正数和零直接返回，负数取反。
//
// 参数:
//   - 值: 输入的整数
//
// 返回:
//   - int: 绝对值
func Int取绝对值(值 int) int {
	if 值 >= 0 {
		return 值
	}
	return -值
}

// Int32ToBytes 将 int32 整数转换为 4 字节的大端序字节集。
// 使用 binary.BigEndian 编码。
//
// 参数:
//   - i: 待转换的 int32 值
//
// 返回:
//   - []byte: 4 字节大端序字节集；转换失败返回空字节集
func Int32ToBytes(i int32) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, i)
	if err != nil {
		return []byte{}
	}
	return buf.Bytes()
}
