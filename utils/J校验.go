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

/*
.版本 2
.子程序 校验_取md5, 文本型, , 取数据MD5
.参数 字节集数据, 字节集, , 要取数据摘要的字节集
.参数 返回值转成大写, 逻辑型, 可空 , 可空，默认为假。假=小写  真=大写
*/
func J校验_取md5(字节集数据 []byte, 返回值转成大写 bool) string {

	has := md5.Sum(字节集数据)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	if 返回值转成大写 {
		md5str = strings.ToUpper(md5str)
	}
	return md5str
}
func J校验_取md5_文本(文本数据 string, 返回值转成大写 bool) string {
	return J校验_取md5([]byte(文本数据), 返回值转成大写)
}

// 用于取crc32，结果为16进制文本
func J校验_取Crc32(数据 []byte, 返回值转成大写 bool) string {
	// 创建CRC32校验器
	crc32Hash := crc32.NewIEEE()

	// 计算CRC32校验值
	_, err := crc32Hash.Write(数据)
	if err != nil {
		return ""
	}
	checksum := crc32Hash.Sum32()

	// 将校验值转换为16进制文本
	结果 := fmt.Sprintf("%08X", checksum)

	// 根据返回值转成大写的标志进行处理
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}

	return 结果
}

// 返回40位的校验数据
func J校验_取sha1(数据 []byte, 返回值转成大写 bool) string {
	// 计算SHA1校验值
	sha1Hash := sha1.New()
	sha1Hash.Write(数据)
	checksum := sha1Hash.Sum(nil)

	// 将校验值转换为16进制文本
	结果 := fmt.Sprintf("%x", checksum)

	// 根据返回值转成大写的标志进行处理
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// 返回校验数据
func J校验_取sha256(数据 []byte, 返回值转成大写 bool) string {
	// 计算SHA1校验值
	sha1Hash := sha256.New()
	sha1Hash.Write(数据)
	checksum := sha1Hash.Sum(nil)

	// 将校验值转换为16进制文本
	结果 := fmt.Sprintf("%x", checksum)

	// 根据返回值转成大写的标志进行处理
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}

// 返回校验数据
func J校验_取sha512(数据 []byte, 返回值转成大写 bool) string {
	// 计算SHA1校验值
	sha1Hash := sha512.New()
	sha1Hash.Write(数据)
	checksum := sha1Hash.Sum(nil)

	// 将校验值转换为16进制文本
	结果 := fmt.Sprintf("%x", checksum)

	// 根据返回值转成大写的标志进行处理
	if 返回值转成大写 {
		结果 = strings.ToUpper(结果)
	}
	return 结果
}
