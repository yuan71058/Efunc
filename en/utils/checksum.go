package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"io"
	"os"
	"strings"
)

// Checksum_MD5 计算字节集数据的 MD5 哈希值，返回 32 位十六进制字符串。
//
// 参数:
//   - data: 待计算哈希的字节集
//   - upperCase: true 返回大写哈希，false 返回小写哈希
//
// 返回:
//   - string: 32 位十六进制 MD5 哈希值
func Checksum_MD5(data []byte, upperCase bool) string {
	hash := md5.Sum(data)
	result := fmt.Sprintf("%x", hash)
	if upperCase {
		result = strings.ToUpper(result)
	}
	return result
}

// Checksum_MD5Text 计算文本字符串的 MD5 哈希值。
//
// 参数:
//   - text: 待计算哈希的文本
//   - upperCase: true 返回大写哈希，false 返回小写哈希
//
// 返回:
//   - string: 32 位十六进制 MD5 哈希值
func Checksum_MD5Text(text string, upperCase bool) string {
	return Checksum_MD5([]byte(text), upperCase)
}

// Checksum_MD5Short 计算 16 位 MD5 哈希值（取 32 位的中间 16 位）。
//
// 参数:
//   - data: 待计算哈希的字节集
//   - upperCase: true 返回大写
//
// 返回:
//   - string: 16 位十六进制 MD5 哈希值
func Checksum_MD5Short(data []byte, upperCase bool) string {
	full := Checksum_MD5(data, upperCase)
	return full[8:24]
}

// Checksum_MD5File 计算文件的 MD5 哈希值（流式读取，支持大文件）。
//
// 参数:
//   - filePath: 文件路径
//   - upperCase: true 返回大写
//
// 返回:
//   - string: 32 位十六进制 MD5 哈希值
//   - error: 文件读取失败时返回错误
func Checksum_MD5File(filePath string, upperCase bool) (string, error) {
	return computeFileHash(md5.New(), filePath, upperCase)
}

// Checksum_CRC32 计算字节集数据的 CRC32 校验值，返回 8 位十六进制字符串。
//
// 参数:
//   - data: 待计算校验值的字节集
//   - upperCase: true 返回大写，false 返回小写
//
// 返回:
//   - string: 8 位十六进制 CRC32 校验值；计算失败返回空串
func Checksum_CRC32(data []byte, upperCase bool) string {
	crc32Hash := crc32.NewIEEE()
	_, err := crc32Hash.Write(data)
	if err != nil {
		return ""
	}
	checksum := crc32Hash.Sum32()
	result := fmt.Sprintf("%08x", checksum)
	if upperCase {
		result = strings.ToUpper(result)
	}
	return result
}

// Checksum_CRC32File 计算文件的 CRC32 校验值。
//
// 参数:
//   - filePath: 文件路径
//   - upperCase: true 返回大写
//
// 返回:
//   - string: 8 位十六进制 CRC32 校验值
//   - error: 文件读取失败时返回错误
func Checksum_CRC32File(filePath string, upperCase bool) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := crc32.NewIEEE()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	result := fmt.Sprintf("%08x", hash.Sum32())
	if upperCase {
		result = strings.ToUpper(result)
	}
	return result, nil
}

// Checksum_SHA1 计算字节集数据的 SHA1 哈希值，返回 40 位十六进制字符串。
//
// 参数:
//   - data: 待计算哈希的字节集
//   - upperCase: true 返回大写，false 返回小写
//
// 返回:
//   - string: 40 位十六进制 SHA1 哈希值
func Checksum_SHA1(data []byte, upperCase bool) string {
	sha1Hash := sha1.New()
	sha1Hash.Write(data)
	checksum := sha1Hash.Sum(nil)
	result := fmt.Sprintf("%x", checksum)
	if upperCase {
		result = strings.ToUpper(result)
	}
	return result
}

// Checksum_SHA1File 计算文件的 SHA1 哈希值。
//
// 参数:
//   - filePath: 文件路径
//   - upperCase: true 返回大写
//
// 返回:
//   - string: 40 位十六进制 SHA1 哈希值
//   - error: 文件读取失败时返回错误
func Checksum_SHA1File(filePath string, upperCase bool) (string, error) {
	return computeFileHash(sha1.New(), filePath, upperCase)
}

// Checksum_SHA256 计算字节集数据的 SHA256 哈希值，返回 64 位十六进制字符串。
//
// 参数:
//   - data: 待计算哈希的字节集
//   - upperCase: true 返回大写，false 返回小写
//
// 返回:
//   - string: 64 位十六进制 SHA256 哈希值
func Checksum_SHA256(data []byte, upperCase bool) string {
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	checksum := sha256Hash.Sum(nil)
	result := fmt.Sprintf("%x", checksum)
	if upperCase {
		result = strings.ToUpper(result)
	}
	return result
}

// Checksum_SHA256File 计算文件的 SHA256 哈希值。
//
// 参数:
//   - filePath: 文件路径
//   - upperCase: true 返回大写
//
// 返回:
//   - string: 64 位十六进制 SHA256 哈希值
//   - error: 文件读取失败时返回错误
func Checksum_SHA256File(filePath string, upperCase bool) (string, error) {
	return computeFileHash(sha256.New(), filePath, upperCase)
}

// Checksum_SHA512 计算字节集数据的 SHA512 哈希值，返回 128 位十六进制字符串。
//
// 参数:
//   - data: 待计算哈希的字节集
//   - upperCase: true 返回大写，false 返回小写
//
// 返回:
//   - string: 128 位十六进制 SHA512 哈希值
func Checksum_SHA512(data []byte, upperCase bool) string {
	sha512Hash := sha512.New()
	sha512Hash.Write(data)
	checksum := sha512Hash.Sum(nil)
	result := fmt.Sprintf("%x", checksum)
	if upperCase {
		result = strings.ToUpper(result)
	}
	return result
}

// Checksum_SHA512File 计算文件的 SHA512 哈希值。
//
// 参数:
//   - filePath: 文件路径
//   - upperCase: true 返回大写
//
// 返回:
//   - string: 128 位十六进制 SHA512 哈希值
//   - error: 文件读取失败时返回错误
func Checksum_SHA512File(filePath string, upperCase bool) (string, error) {
	return computeFileHash(sha512.New(), filePath, upperCase)
}

// Checksum_HMAC_SHA256 使用 HMAC-SHA256 计算消息认证码。
// 常用于 API 签名验证。
//
// 参数:
//   - key: HMAC 密钥
//   - data: 待认证的数据
//   - upperCase: true 返回大写
//
// 返回:
//   - string: 64 位十六进制 HMAC-SHA256 值
func Checksum_HMAC_SHA256(key []byte, data []byte, upperCase bool) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	result := hex.EncodeToString(mac.Sum(nil))
	if upperCase {
		result = strings.ToUpper(result)
	}
	return result
}

// Checksum_HMAC_SHA1 使用 HMAC-SHA1 计算消息认证码。
//
// 参数:
//   - key: HMAC 密钥
//   - data: 待认证的数据
//   - upperCase: true 返回大写
//
// 返回:
//   - string: 40 位十六进制 HMAC-SHA1 值
func Checksum_HMAC_SHA1(key []byte, data []byte, upperCase bool) string {
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	result := hex.EncodeToString(mac.Sum(nil))
	if upperCase {
		result = strings.ToUpper(result)
	}
	return result
}

// Checksum_HMAC_MD5 使用 HMAC-MD5 计算消息认证码。
//
// 参数:
//   - key: HMAC 密钥
//   - data: 待认证的数据
//   - upperCase: true 返回大写
//
// 返回:
//   - string: 32 位十六进制 HMAC-MD5 值
func Checksum_HMAC_MD5(key []byte, data []byte, upperCase bool) string {
	mac := hmac.New(md5.New, key)
	mac.Write(data)
	result := hex.EncodeToString(mac.Sum(nil))
	if upperCase {
		result = strings.ToUpper(result)
	}
	return result
}

// Checksum_HMAC_SHA512 使用 HMAC-SHA512 计算消息认证码。
//
// 参数:
//   - key: HMAC 密钥
//   - data: 待认证的数据
//   - upperCase: true 返回大写
//
// 返回:
//   - string: 128 位十六进制 HMAC-SHA512 值
func Checksum_HMAC_SHA512(key []byte, data []byte, upperCase bool) string {
	mac := hmac.New(sha512.New, key)
	mac.Write(data)
	result := hex.EncodeToString(mac.Sum(nil))
	if upperCase {
		result = strings.ToUpper(result)
	}
	return result
}

// Checksum_VerifyMD5 校验数据的 MD5 哈希值是否与预期一致。
//
// 参数:
//   - data: 待校验的字节集
//   - expected: 预期的 MD5 哈希值（不区分大小写）
//
// 返回:
//   - bool: 一致返回 true
func Checksum_VerifyMD5(data []byte, expected string) bool {
	actual := Checksum_MD5(data, false)
	return strings.EqualFold(actual, expected)
}

// Checksum_VerifySHA256 校验数据的 SHA256 哈希值是否与预期一致。
//
// 参数:
//   - data: 待校验的字节集
//   - expected: 预期的 SHA256 哈希值（不区分大小写）
//
// 返回:
//   - bool: 一致返回 true
func Checksum_VerifySHA256(data []byte, expected string) bool {
	actual := Checksum_SHA256(data, false)
	return strings.EqualFold(actual, expected)
}

// Checksum_VerifyFileMD5 校验文件的 MD5 哈希值是否与预期一致。
//
// 参数:
//   - filePath: 文件路径
//   - expected: 预期的 MD5 哈希值
//
// 返回:
//   - bool: 一致返回 true
//   - error: 文件读取失败时返回错误
func Checksum_VerifyFileMD5(filePath string, expected string) (bool, error) {
	actual, err := Checksum_MD5File(filePath, false)
	if err != nil {
		return false, err
	}
	return strings.EqualFold(actual, expected), nil
}

// Checksum_VerifyFileSHA256 校验文件的 SHA256 哈希值是否与预期一致。
//
// 参数:
//   - filePath: 文件路径
//   - expected: 预期的 SHA256 哈希值
//
// 返回:
//   - bool: 一致返回 true
//   - error: 文件读取失败时返回错误
func Checksum_VerifyFileSHA256(filePath string, expected string) (bool, error) {
	actual, err := Checksum_SHA256File(filePath, false)
	if err != nil {
		return false, err
	}
	return strings.EqualFold(actual, expected), nil
}

// Checksum_VerifyHMAC 校验 HMAC-SHA256 是否与预期一致。
//
// 参数:
//   - key: HMAC 密钥
//   - data: 待认证的数据
//   - expected: 预期的 HMAC 值
//
// 返回:
//   - bool: 一致返回 true
func Checksum_VerifyHMAC(key []byte, data []byte, expected string) bool {
	actual := Checksum_HMAC_SHA256(key, data, false)
	return hmac.Equal([]byte(actual), []byte(strings.ToLower(expected)))
}

// Checksum_CRC64 计算字节集数据的 CRC64 校验值。
// 使用 ECMA 多项式。
//
// 参数:
//   - data: 待计算的字节集
//   - upperCase: true 返回大写
//
// 返回:
//   - string: 16 位十六进制 CRC64 校验值
func Checksum_CRC64(data []byte, upperCase bool) string {
	table := crc64.MakeTable(crc64.ECMA)
	checksum := crc64.Checksum(data, table)
	result := fmt.Sprintf("%016x", checksum)
	if upperCase {
		result = strings.ToUpper(result)
	}
	return result
}

// computeFileHash 计算文件哈希的通用内部函数。
func computeFileHash(hash hash.Hash, filePath string, upperCase bool) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	result := fmt.Sprintf("%x", hash.Sum(nil))
	if upperCase {
		result = strings.ToUpper(result)
	}
	return result, nil
}