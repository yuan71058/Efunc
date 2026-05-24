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
	"hash/adler32"
	"io"
	"os"
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
func J校验_取md5(字节集数据 []byte, 返回值转成大写 bool) string {
	has := md5.Sum(字节集数据)
	md5str := fmt.Sprintf("%x", has)
	if 返回值转成大写 {
		md5str = strings.ToUpper(md5str)
	}
	return md5str
}

// J校验_取md5_文本 计算文本字符串的 MD5 哈希值。
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

// J校验_取md5_16位 计算 16 位 MD5 哈希值（取 32 位的中间 16 位）。
//
// 参数:
//   - 字节集数据: 待计算哈希的字节集
//   - 返回值转成大写: true 返回大写
//
// 返回:
//   - string: 16 位十六进制 MD5 哈希值
func J校验_取md5_16位(字节集数据 []byte, 返回值转成大写 bool) string {
	完整 := J校验_取md5(字节集数据, 返回值转成大写)
	return 完整[8:24]
}

// J校验_取md5_文件 计算文件的 MD5 哈希值（流式读取，支持大文件）。
//
// 参数:
//   - 文件路径: 文件路径
//   - 返回值转成大写: true 返回大写
//
// 返回:
//   - string: 32 位十六进制 MD5 哈希值
//   - error: 文件读取失败时返回错误
func J校验_取md5_文件(文件路径 string, 返回值转成大写 bool) (string, error) {
	return 计算文件哈希(md5.New(), 文件路径, 返回值转成大写)
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
	结果 := fmt.Sprintf("%08x", checksum)
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// J校验_取Crc32_文件 计算文件的 CRC32 校验值。
//
// 参数:
//   - 文件路径: 文件路径
//   - 返回值转成大写: true 返回大写
//
// 返回:
//   - string: 8 位十六进制 CRC32 校验值
//   - error: 文件读取失败时返回错误
func J校验_取Crc32_文件(文件路径 string, 返回值转成大写 bool) (string, error) {
	file, err := os.Open(文件路径)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := crc32.NewIEEE()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	结果 := fmt.Sprintf("%08x", hash.Sum32())
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果, nil
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

// J校验_取sha1_文件 计算文件的 SHA1 哈希值。
//
// 参数:
//   - 文件路径: 文件路径
//   - 返回值转成大写: true 返回大写
//
// 返回:
//   - string: 40 位十六进制 SHA1 哈希值
//   - error: 文件读取失败时返回错误
func J校验_取sha1_文件(文件路径 string, 返回值转成大写 bool) (string, error) {
	return 计算文件哈希(sha1.New(), 文件路径, 返回值转成大写)
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
	sha256Hash := sha256.New()
	sha256Hash.Write(数据)
	checksum := sha256Hash.Sum(nil)
	结果 := fmt.Sprintf("%x", checksum)
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// J校验_取sha256_文件 计算文件的 SHA256 哈希值。
//
// 参数:
//   - 文件路径: 文件路径
//   - 返回值转成大写: true 返回大写
//
// 返回:
//   - string: 64 位十六进制 SHA256 哈希值
//   - error: 文件读取失败时返回错误
func J校验_取sha256_文件(文件路径 string, 返回值转成大写 bool) (string, error) {
	return 计算文件哈希(sha256.New(), 文件路径, 返回值转成大写)
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
	sha512Hash := sha512.New()
	sha512Hash.Write(数据)
	checksum := sha512Hash.Sum(nil)
	结果 := fmt.Sprintf("%x", checksum)
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// J校验_取sha512_文件 计算文件的 SHA512 哈希值。
//
// 参数:
//   - 文件路径: 文件路径
//   - 返回值转成大写: true 返回大写
//
// 返回:
//   - string: 128 位十六进制 SHA512 哈希值
//   - error: 文件读取失败时返回错误
func J校验_取sha512_文件(文件路径 string, 返回值转成大写 bool) (string, error) {
	return 计算文件哈希(sha512.New(), 文件路径, 返回值转成大写)
}

// J校验_HMAC_SHA256 使用 HMAC-SHA256 计算消息认证码。
// 常用于 API 签名验证。
//
// 参数:
//   - 密钥: HMAC 密钥
//   - 数据: 待认证的数据
//   - 返回值转成大写: true 返回大写
//
// 返回:
//   - string: 64 位十六进制 HMAC-SHA256 值
func J校验_HMAC_SHA256(密钥 []byte, 数据 []byte, 返回值转成大写 bool) string {
	mac := hmac.New(sha256.New, 密钥)
	mac.Write(数据)
	结果 := hex.EncodeToString(mac.Sum(nil))
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// J校验_HMAC_SHA1 使用 HMAC-SHA1 计算消息认证码。
//
// 参数:
//   - 密钥: HMAC 密钥
//   - 数据: 待认证的数据
//   - 返回值转成大写: true 返回大写
//
// 返回:
//   - string: 40 位十六进制 HMAC-SHA1 值
func J校验_HMAC_SHA1(密钥 []byte, 数据 []byte, 返回值转成大写 bool) string {
	mac := hmac.New(sha1.New, 密钥)
	mac.Write(数据)
	结果 := hex.EncodeToString(mac.Sum(nil))
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// J校验_HMAC_MD5 使用 HMAC-MD5 计算消息认证码。
//
// 参数:
//   - 密钥: HMAC 密钥
//   - 数据: 待认证的数据
//   - 返回值转成大写: true 返回大写
//
// 返回:
//   - string: 32 位十六进制 HMAC-MD5 值
func J校验_HMAC_MD5(密钥 []byte, 数据 []byte, 返回值转成大写 bool) string {
	mac := hmac.New(md5.New, 密钥)
	mac.Write(数据)
	结果 := hex.EncodeToString(mac.Sum(nil))
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// J校验_HMAC_SHA512 使用 HMAC-SHA512 计算消息认证码。
//
// 参数:
//   - 密钥: HMAC 密钥
//   - 数据: 待认证的数据
//   - 返回值转成大写: true 返回大写
//
// 返回:
//   - string: 128 位十六进制 HMAC-SHA512 值
func J校验_HMAC_SHA512(密钥 []byte, 数据 []byte, 返回值转成大写 bool) string {
	mac := hmac.New(sha512.New, 密钥)
	mac.Write(数据)
	结果 := hex.EncodeToString(mac.Sum(nil))
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// J校验_校验MD5 校验数据的 MD5 哈希值是否与预期一致。
//
// 参数:
//   - 数据: 待校验的字节集
//   - 预期值: 预期的 MD5 哈希值（不区分大小写）
//
// 返回:
//   - bool: 一致返回 true
func J校验_校验MD5(数据 []byte, 预期值 string) bool {
	实际值 := J校验_取md5(数据, false)
	return strings.EqualFold(实际值, 预期值)
}

// J校验_校验SHA256 校验数据的 SHA256 哈希值是否与预期一致。
//
// 参数:
//   - 数据: 待校验的字节集
//   - 预期值: 预期的 SHA256 哈希值（不区分大小写）
//
// 返回:
//   - bool: 一致返回 true
func J校验_校验SHA256(数据 []byte, 预期值 string) bool {
	实际值 := J校验_取sha256(数据, false)
	return strings.EqualFold(实际值, 预期值)
}

// J校验_校验文件MD5 校验文件的 MD5 哈希值是否与预期一致。
//
// 参数:
//   - 文件路径: 文件路径
//   - 预期值: 预期的 MD5 哈希值
//
// 返回:
//   - bool: 一致返回 true
//   - error: 文件读取失败时返回错误
func J校验_校验文件MD5(文件路径 string, 预期值 string) (bool, error) {
	实际值, err := J校验_取md5_文件(文件路径, false)
	if err != nil {
		return false, err
	}
	return strings.EqualFold(实际值, 预期值), nil
}

// J校验_校验文件SHA256 校验文件的 SHA256 哈希值是否与预期一致。
//
// 参数:
//   - 文件路径: 文件路径
//   - 预期值: 预期的 SHA256 哈希值
//
// 返回:
//   - bool: 一致返回 true
//   - error: 文件读取失败时返回错误
func J校验_校验文件SHA256(文件路径 string, 预期值 string) (bool, error) {
	实际值, err := J校验_取sha256_文件(文件路径, false)
	if err != nil {
		return false, err
	}
	return strings.EqualFold(实际值, 预期值), nil
}

// J校验_校验HMAC 校验 HMAC-SHA256 是否与预期一致。
//
// 参数:
//   - 密钥: HMAC 密钥
//   - 数据: 待认证的数据
//   - 预期值: 预期的 HMAC 值
//
// 返回:
//   - bool: 一致返回 true
func J校验_校验HMAC(密钥 []byte, 数据 []byte, 预期值 string) bool {
	实际值 := J校验_HMAC_SHA256(密钥, 数据, false)
	return hmac.Equal([]byte(实际值), []byte(strings.ToLower(预期值)))
}

// J校验_取CRC64 计算字节集数据的 CRC64 校验值。
// 使用 ECMA 多项式。
//
// 参数:
//   - 数据: 待计算的字节集
//   - 返回值转成大写: true 返回大写
//
// 返回:
//   - string: 16 位十六进制 CRC64 校验值
func J校验_取CRC64(数据 []byte, 返回值转成大写 bool) string {
	table := crc64.MakeTable(crc64.ECMA)
	checksum := crc64.Checksum(数据, table)
	结果 := fmt.Sprintf("%016x", checksum)
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// J校验_取Adler32 计算字节集数据的 Adler-32 校验值。
//
// 参数:
//   - 数据: 待计算的字节集
//   - 返回值转成大写: true 返回大写
//
// 返回:
//   - string: 8 位十六进制 Adler-32 校验值
func J校验_取Adler32(数据 []byte, 返回值转成大写 bool) string {
	checksum := adler32.Checksum(数据)
	结果 := fmt.Sprintf("%08x", checksum)
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// 计算文件哈希 通用文件哈希计算辅助函数。
func 计算文件哈希(哈希器 hash.Hash, 文件路径 string, 大写 bool) (string, error) {
	file, err := os.Open(文件路径)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(哈希器, file); err != nil {
		return "", err
	}
	结果 := hex.EncodeToString(哈希器.Sum(nil))
	if 大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果, nil
}
