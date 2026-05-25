package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"errors"
	"io/ioutil"
)

// ByteData_HexToBytes 将十六进制字符串转换为字节集。
// 每两个十六进制字符对应一个字节。
//
// 参数:
//   - hexString: 十六进制格式的字符串（如 "48656c6c6f"）
//
// 返回:
//   - []byte: 转换后的字节集；转换失败返回空字节集
//
// 示例:
//
//	ByteData_HexToBytes("48656c6c6f")  // []byte("Hello")
func ByteData_HexToBytes(hexString string) []byte {
	data, _ := hex.DecodeString(hexString)
	return data
}

// ByteData_BytesToHex 将字节集转换为十六进制字符串（小写）。
//
// 参数:
//   - data: 待转换的字节集
//
// 返回:
//   - string: 小写十六进制格式的字符串
//
// 示例:
//
//	ByteData_BytesToHex([]byte("Hello"))  // "48656c6c6f"
func ByteData_BytesToHex(data []byte) string {
	return hex.EncodeToString(data)
}

// ByteData_Find 在字节集中查找子字节集的位置。
// 位置从 1 开始计数（兼容易语言习惯），未找到返回 -1。
//
// 参数:
//   - source: 被搜索的源字节集
//   - target: 要查找的目标字节集
//   - startPos: 开始搜索的位置（从 1 开始），无效值自动设为 1
//
// 返回:
//   - int: 找到的位置（从 1 开始）；未找到返回 -1
func ByteData_Find(source []byte, target []byte, startPos int) int {
	if len(source) == 0 || len(target) == 0 {
		return -1
	}

	if startPos <= 0 || startPos > len(source) {
		startPos = 1
	}

	for i := startPos - 1; i < len(source); i++ {
		if bytes.HasPrefix(source[i:], target) {
			return i + 1
		}
	}

	return -1
}

// ByteData_GzipDecompress 对 Gzip 压缩的字节集进行解压。
//
// 参数:
//   - data: Gzip 压缩的字节集
//
// 返回:
//   - decompressed: 解压后的字节集
//   - err: 解压失败时的错误信息（空输入或格式错误等）
func ByteData_GzipDecompress(data []byte) (decompressed []byte, err error) {
	if len(data) == 0 {
		err = errors.New("待解压数据不能为空")
		return
	}
	gzipReader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return
	}
	defer gzipReader.Close()
	decompressed, err = ioutil.ReadAll(gzipReader)
	return
}