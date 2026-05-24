package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash/crc32"
	"strings"
)

// J校验_取md5 计算字节集数据的 MD5 哈希值，返回 32 位十六进制字符串。
//
// 参数:
//   - 字节集数据: 待计算哈希的字节集
//   - 返回值转成大写: true 返回大写哈希，false 返回小写哈希
//
// 返回:
//   - string: 32 位十六进制 MD5 哈希值
//
// 示例:
//
//	J校验_取md5([]byte("hello"), false)  // "5d41402abc4b2a76b9719d911017c592"
//	J校验_取md5([]byte("hello"), true)   // "5D41402ABC4B2A76B9719D911017C592"
func J校验_取md5(字节集数据 []byte, 返回值转成大写 bool) string {

	has := md5.Sum(字节集数据)
	md5str := fmt.Sprintf("%x", has)
	if 返回值转成大写 {
		md5str = strings.ToUpper(md5str)
	}
	return md5str
}

// J校验_取md5_文本 计算文本字符串的 MD5 哈希值。
// 内部将文本转换为字节集后调用 J校验_取md5。
//
// 参数:
//   - 文本数据: 待计算哈希的文本
//   - 返回值转成大写: true 返回大写哈希，false 返回小写哈希
//
// 返回:
//   - string: 32 位十六进制 MD5 哈希值
func J校验_取md5_文本(文本数据 string, 返回值转成大写 bool) string {
	return J校验_取md5([]byte(文本数据), 返回值转成大写)
}

// J校验_取Crc32 计算字节集数据的 CRC32 校验值，返回 8 位十六进制字符串。
//
// 参数:
//   - 数据: 待计算校验值的字节集
//   - 返回值转成大写: true 返回大写，false 返回小写
//
// 返回:
//   - string: 8 位十六进制 CRC32 校验值；计算失败返回空串
func J校验_取Crc32(数据 []byte, 返回值转成大写 bool) string {
	crc32Hash := crc32.NewIEEE()

	_, err := crc32Hash.Write(数据)
	if err != nil {
		return ""
	}
	checksum := crc32Hash.Sum32()

	结果 := fmt.Sprintf("%08X", checksum)

	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}

	return 结果
}

// J校验_取sha1 计算字节集数据的 SHA1 哈希值，返回 40 位十六进制字符串。
//
// 参数:
//   - 数据: 待计算哈希的字节集
//   - 返回值转成大写: true 返回大写，false 返回小写
//
// 返回:
//   - string: 40 位十六进制 SHA1 哈希值
func J校验_取sha1(数据 []byte, 返回值转成大写 bool) string {
	sha1Hash := sha1.New()
	sha1Hash.Write(数据)
	checksum := sha1Hash.Sum(nil)

	结果 := fmt.Sprintf("%x", checksum)

	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// J校验_取sha256 计算字节集数据的 SHA256 哈希值，返回 64 位十六进制字符串。
//
// 参数:
//   - 数据: 待计算哈希的字节集
//   - 返回值转成大写: true 返回大写，false 返回小写
//
// 返回:
//   - string: 64 位十六进制 SHA256 哈希值
func J校验_取sha256(数据 []byte, 返回值转成大写 bool) string {
	sha1Hash := sha256.New()
	sha1Hash.Write(数据)
	checksum := sha1Hash.Sum(nil)

	结果 := fmt.Sprintf("%x", checksum)

	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// J校验_取sha512 计算字节集数据的 SHA512 哈希值，返回 128 位十六进制字符串。
//
// 参数:
//   - 数据: 待计算哈希的字节集
//   - 返回值转成大写: true 返回大写，false 返回小写
//
// 返回:
//   - string: 128 位十六进制 SHA512 哈希值
func J校验_取sha512(数据 []byte, 返回值转成大写 bool) string {
	sha1Hash := sha512.New()
	sha1Hash.Write(数据)
	checksum := sha1Hash.Sum(nil)

	结果 := fmt.Sprintf("%x", checksum)

	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}
