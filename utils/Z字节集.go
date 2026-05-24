package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"errors"
	"io/ioutil"
)

// Z字节集_十六进制到字节集 将十六进制字符串转换为字节集。
// 每两个十六进制字符对应一个字节。
//
// 参数:
//   - 原始16进制文本: 十六进制格式的字符串（如 "48656c6c6f"）
//
// 返回:
//   - []byte: 转换后的字节集；转换失败返回空字节集
//
// 示例:
//
//	Z字节集_十六进制到字节集("48656c6c6f")  // []byte("Hello")
func Z字节集_十六进制到字节集(原始16进制文本 string) []byte {
	字节集, _ := hex.DecodeString(原始16进制文本)
	return 字节集
}

// Z字节集_字节集到十六进制 将字节集转换为十六进制字符串（小写）。
//
// 参数:
//   - 字节集: 待转换的字节集
//
// 返回:
//   - string: 小写十六进制格式的字符串
//
// 示例:
//
//	Z字节集_字节集到十六进制([]byte("Hello"))  // "48656c6c6f"
func Z字节集_字节集到十六进制(字节集 []byte) string {
	return hex.EncodeToString(字节集)
}

// Z字节集_寻找 在字节集中查找子字节集的位置。
// 位置从 1 开始计数（兼容易语言习惯），未找到返回 -1。
//
// 参数:
//   - 被搜寻的字节集: 被搜索的源字节集
//   - 欲寻找的字节集: 要查找的目标字节集
//   - 起始搜寻位置: 开始搜索的位置（从 1 开始），无效值自动设为 1
//
// 返回:
//   - int: 找到的位置（从 1 开始）；未找到返回 -1
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

// Z字节集_Gzip解压 对 Gzip 压缩的字节集进行解压。
//
// 参数:
//   - 字节集: Gzip 压缩的字节集
//
// 返回:
//   - data: 解压后的字节集
//   - err: 解压失败时的错误信息（空输入或格式错误等）
func Z字节集_Gzip解压(字节集 []byte) (data []byte, err error) {
	if len(字节集) == 0 {
		err = errors.New("待解压数据不能为空")
		return
	}
	gzipReader, err := gzip.NewReader(bytes.NewReader(字节集))
	if err != nil {
		return
	}
	defer gzipReader.Close()
	data, err = ioutil.ReadAll(gzipReader)
	return
}
