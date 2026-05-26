package main

import (
	"encoding/json"

	"github.com/yuan71058/Efunc/utils"
)

func registerCryptoFunctions() {
	r := globalRegistry

	r.Register("J加解密_AES_CBC加密", []string{"明文", "密钥", "IV"}, "AES-CBC加密(输入Base64字节集,返回Base64密文)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				明文 string `json:"明文"`
				密钥 string `json:"密钥"`
				IV  string `json:"IV"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			明文, err := decodeBase64(v.明文)
			if err != nil {
				return errResultf("明文base64解码失败: %v", err)
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			iv, err := decodeBase64(v.IV)
			if err != nil {
				return errResultf("IV base64解码失败: %v", err)
			}
			result, err := utils.J加解密_AES_CBC加密(明文, 密钥, iv)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("J加解密_AES_CBC解密", []string{"密文Base64", "密钥", "IV"}, "AES-CBC解密(返回Base64字节集)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				密文 string `json:"密文Base64"`
				密钥 string `json:"密钥"`
				IV  string `json:"IV"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			iv, err := decodeBase64(v.IV)
			if err != nil {
				return errResultf("IV base64解码失败: %v", err)
			}
			result, err := utils.J加解密_AES_CBC解密(v.密文, 密钥, iv)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(encodeBase64(result))
		})

	r.Register("J加解密_AES_GCM加密", []string{"明文", "密钥", "附加数据"}, "AES-GCM加密",
		func(p json.RawMessage) *CallResult {
			var v struct {
				明文   string `json:"明文"`
				密钥   string `json:"密钥"`
				附加数据 string `json:"附加数据"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			明文, err := decodeBase64(v.明文)
			if err != nil {
				return errResultf("明文base64解码失败: %v", err)
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			附加数据, _ := decodeBase64(v.附加数据)
			result, err := utils.J加解密_AES_GCM加密(明文, 密钥, 附加数据)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("J加解密_AES_GCM解密", []string{"密文Base64", "密钥", "附加数据"}, "AES-GCM解密",
		func(p json.RawMessage) *CallResult {
			var v struct {
				密文   string `json:"密文Base64"`
				密钥   string `json:"密钥"`
				附加数据 string `json:"附加数据"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			附加数据, _ := decodeBase64(v.附加数据)
			result, err := utils.J加解密_AES_GCM解密(v.密文, 密钥, 附加数据)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(encodeBase64(result))
		})

	r.Register("J加解密_AES_ECB加密", []string{"明文", "密钥"}, "AES-ECB加密",
		func(p json.RawMessage) *CallResult {
			var v struct {
				明文 string `json:"明文"`
				密钥 string `json:"密钥"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			明文, err := decodeBase64(v.明文)
			if err != nil {
				return errResultf("明文base64解码失败: %v", err)
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			result, err := utils.J加解密_AES_ECB加密(明文, 密钥)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("J加解密_AES_ECB解密", []string{"密文Base64", "密钥"}, "AES-ECB解密",
		func(p json.RawMessage) *CallResult {
			var v struct {
				密文 string `json:"密文Base64"`
				密钥 string `json:"密钥"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			result, err := utils.J加解密_AES_ECB解密(v.密文, 密钥)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(encodeBase64(result))
		})

	r.Register("J加解密_AES_CTR加密", []string{"明文", "密钥"}, "AES-CTR加密",
		func(p json.RawMessage) *CallResult {
			var v struct {
				明文 string `json:"明文"`
				密钥 string `json:"密钥"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			明文, err := decodeBase64(v.明文)
			if err != nil {
				return errResultf("明文base64解码失败: %v", err)
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			result, err := utils.J加解密_AES_CTR加密(明文, 密钥)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("J加解密_AES_CTR解密", []string{"密文Base64", "密钥"}, "AES-CTR解密",
		func(p json.RawMessage) *CallResult {
			var v struct {
				密文 string `json:"密文Base64"`
				密钥 string `json:"密钥"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			result, err := utils.J加解密_AES_CTR解密(v.密文, 密钥)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(encodeBase64(result))
		})

	r.Register("J加解密_RC4", []string{"数据", "密钥"}, "RC4加解密(返回Base64)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数据 string `json:"数据"`
				密钥 string `json:"密钥"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			数据, err := decodeBase64(v.数据)
			if err != nil {
				return errResultf("数据base64解码失败: %v", err)
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			return okResult(utils.J加解密_RC4(数据, 密钥))
		})

	r.Register("J加解密_XOR", []string{"数据", "密钥"}, "XOR加解密(返回Base64字节集)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数据 string `json:"数据"`
				密钥 string `json:"密钥"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			数据, err := decodeBase64(v.数据)
			if err != nil {
				return errResultf("数据base64解码失败: %v", err)
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			result := utils.J加解密_XOR(数据, 密钥)
			return okResult(encodeBase64(result))
		})

	r.Register("J加解密_TEA加密", []string{"明文", "密钥"}, "TEA加密",
		func(p json.RawMessage) *CallResult {
			var v struct {
				明文 string `json:"明文"`
				密钥 string `json:"密钥"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			明文, err := decodeBase64(v.明文)
			if err != nil {
				return errResultf("明文base64解码失败: %v", err)
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			result, err := utils.J加解密_TEA加密(明文, 密钥)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("J加解密_TEA解密", []string{"密文Base64", "密钥"}, "TEA解密",
		func(p json.RawMessage) *CallResult {
			var v struct {
				密文 string `json:"密文Base64"`
				密钥 string `json:"密钥"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			result, err := utils.J加解密_TEA解密(v.密文, 密钥)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(encodeBase64(result))
		})

	r.Register("J加解密_XXTEA加密", []string{"明文", "密钥"}, "XXTEA加密",
		func(p json.RawMessage) *CallResult {
			var v struct {
				明文 string `json:"明文"`
				密钥 string `json:"密钥"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			明文, err := decodeBase64(v.明文)
			if err != nil {
				return errResultf("明文base64解码失败: %v", err)
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			result, err := utils.J加解密_XXTEA加密(明文, 密钥)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("J加解密_XXTEA解密", []string{"密文Base64", "密钥"}, "XXTEA解密",
		func(p json.RawMessage) *CallResult {
			var v struct {
				密文 string `json:"密文Base64"`
				密钥 string `json:"密钥"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			result, err := utils.J加解密_XXTEA解密(v.密文, 密钥)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(encodeBase64(result))
		})

	r.Register("J加解密_生成AES密钥", []string{"长度"}, "生成AES密钥(返回Base64)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				长度 int `json:"长度"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			result, err := utils.J加解密_生成AES密钥(v.长度)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(encodeBase64(result))
		})

	r.Register("J加解密_生成IV", []string{"块大小"}, "生成随机IV(返回Base64)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				块大小 int `json:"块大小"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			result, err := utils.J加解密_生成IV(v.块大小)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(encodeBase64(result))
		})
}

func registerChecksumFunctions() {
	r := globalRegistry

	r.Register("J校验_取md5", []string{"字节集数据", "返回值转成大写"}, "MD5哈希(输入Base64字节集)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数据 string `json:"字节集数据"`
				大写 bool   `json:"返回值转成大写"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data, err := decodeBase64(v.数据)
			if err != nil {
				return errResultf("base64解码失败: %v", err)
			}
			return okResult(utils.J校验_取md5(data, v.大写))
		})

	r.Register("J校验_取md5_文本", []string{"文本数据", "返回值转成大写"}, "MD5哈希(文本输入)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"文本数据"`
				大写 bool   `json:"返回值转成大写"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.J校验_取md5_文本(v.文本, v.大写))
		})

	r.Register("J校验_取md5_文件", []string{"文件路径", "返回值转成大写"}, "MD5哈希(文件)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				路径 string `json:"文件路径"`
				大写 bool   `json:"返回值转成大写"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			result, err := utils.J校验_取md5_文件(v.路径, v.大写)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})

	r.Register("J校验_取sha256", []string{"数据", "返回值转成大写"}, "SHA256哈希(输入Base64字节集)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数据 string `json:"数据"`
				大写 bool   `json:"返回值转成大写"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data, err := decodeBase64(v.数据)
			if err != nil {
				return errResultf("base64解码失败: %v", err)
			}
			return okResult(utils.J校验_取sha256(data, v.大写))
		})

	r.Register("J校验_HMAC_SHA256", []string{"密钥", "数据", "返回值转成大写"}, "HMAC-SHA256",
		func(p json.RawMessage) *CallResult {
			var v struct {
				密钥 string `json:"密钥"`
				数据 string `json:"数据"`
				大写 bool   `json:"返回值转成大写"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			密钥, err := decodeBase64(v.密钥)
			if err != nil {
				return errResultf("密钥base64解码失败: %v", err)
			}
			data, err := decodeBase64(v.数据)
			if err != nil {
				return errResultf("数据base64解码失败: %v", err)
			}
			return okResult(utils.J校验_HMAC_SHA256(密钥, data, v.大写))
		})

	r.Register("J校验_取CRC64", []string{"数据", "返回值转成大写"}, "CRC64校验",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数据 string `json:"数据"`
				大写 bool   `json:"返回值转成大写"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data, err := decodeBase64(v.数据)
			if err != nil {
				return errResultf("base64解码失败: %v", err)
			}
			return okResult(utils.J校验_取CRC64(data, v.大写))
		})
}
