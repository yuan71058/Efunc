package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

// ============================================================
// AES 加解密（常用模式：CBC / ECB / GCM / CTR / CFB / OFB）
// ============================================================

// Crypto_AES_CBC_Encrypt 使用 AES-CBC 模式加密数据。
// 采用 PKCS7 填充方案，返回 Base64 编码的密文。
//
// 参数:
//   - plaintext: 待加密的原始数据
//   - key: AES 密钥，长度必须为 16（AES-128）、24（AES-192）或 32（AES-256）字节
//   - iv: 初始化向量，长度必须为 16 字节
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func Crypto_AES_CBC_Encrypt(plaintext, key, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(iv) != block.BlockSize() {
		return "", errors.New("IV长度必须为16字节")
	}

	plaintext = pkcs7Padding(plaintext, block.BlockSize())

	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Crypto_AES_CBC_Decrypt AES-CBC 模式解密，输入为 Base64 密文。
//
// 参数:
//   - cipherBase64: Base64 编码的密文文本
//   - key: AES 密钥
//   - iv: 初始化向量
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func Crypto_AES_CBC_Decrypt(cipherBase64 string, key, iv []byte) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(iv) != block.BlockSize() {
		return nil, errors.New("IV长度必须为16字节")
	}
	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, errors.New("密文长度不是块大小的整数倍")
	}

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = pkcs7Unpadding(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// Crypto_AES_ECB_Encrypt 使用 AES-ECB 模式加密数据（不推荐在生产环境使用）。
// ECB 模式不使用 IV，安全性较低，相同明文块会生成相同密文块。
//
// 参数:
//   - plaintext: 待加密的原始数据
//   - key: AES 密钥
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func Crypto_AES_ECB_Encrypt(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext = pkcs7Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	blockSize := block.BlockSize()

	for i := 0; i < len(plaintext); i += blockSize {
		block.Encrypt(ciphertext[i:i+blockSize], plaintext[i:i+blockSize])
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Crypto_AES_ECB_Decrypt AES-ECB 模式解密。
//
// 参数:
//   - cipherBase64: Base64 编码的密文文本
//   - key: AES 密钥
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func Crypto_AES_ECB_Decrypt(cipherBase64 string, key []byte) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, errors.New("密文长度不是块大小的整数倍")
	}

	plaintext := make([]byte, len(ciphertext))
	blockSize := block.BlockSize()

	for i := 0; i < len(ciphertext); i += blockSize {
		block.Decrypt(plaintext[i:i+blockSize], ciphertext[i:i+blockSize])
	}

	plaintext, err = pkcs7Unpadding(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// Crypto_AES_GCM_Encrypt 使用 AES-GCM 认证加密模式加密数据。
// GCM 模式同时提供机密性和完整性验证，推荐用于安全敏感场景。
// 自动生成 12 字节随机 nonce，格式为 nonce + 密文。
//
// 参数:
//   - plaintext: 待加密的原始数据
//   - key: AES 密钥
//   - additionalData: 可选，参与认证但不加密的附加数据
//
// 返回:
//   - string: Base64 编码的密文（nonce + 密文 + tag）
//   - error: 加密失败时返回错误信息
func Crypto_AES_GCM_Encrypt(plaintext, key, additionalData []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, additionalData)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Crypto_AES_GCM_Decrypt AES-GCM 认证解密。
//
// 参数:
//   - cipherBase64: Base64 编码的密文
//   - key: AES 密钥
//   - additionalData: 加密时传入的附加数据，必须一致
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败或认证失败时返回错误信息
func Crypto_AES_GCM_Decrypt(cipherBase64 string, key, additionalData []byte) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("密文长度不足")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, additionalData)
}

// Crypto_AES_CTR_Encrypt 使用 AES-CTR 模式加密数据。
// CTR 模式为流加密模式，不需要填充，适合任意长度数据。
// 自动生成随机 IV，格式为 IV + 密文。
//
// 参数:
//   - plaintext: 待加密的原始数据
//   - key: AES 密钥
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func Crypto_AES_CTR_Encrypt(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := make([]byte, block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	result := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(result), nil
}

// Crypto_AES_CTR_Decrypt AES-CTR 模式解密。
//
// 参数:
//   - cipherBase64: Base64 编码的密文
//   - key: AES 密钥
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func Crypto_AES_CTR_Decrypt(cipherBase64 string, key []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(data) < blockSize {
		return nil, errors.New("密文长度不足")
	}

	iv := data[:blockSize]
	ciphertext := data[blockSize:]

	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}

// Crypto_AES_CFB_Encrypt 使用 AES-CFB 模式加密数据。
// CFB 模式为流加密模式，不需要填充。
// 自动生成随机 IV，格式为 IV + 密文。
//
// 参数:
//   - plaintext: 待加密的原始数据
//   - key: AES 密钥
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func Crypto_AES_CFB_Encrypt(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := make([]byte, block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	result := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(result), nil
}

// Crypto_AES_CFB_Decrypt AES-CFB 模式解密。
//
// 参数:
//   - cipherBase64: Base64 编码的密文
//   - key: AES 密钥
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func Crypto_AES_CFB_Decrypt(cipherBase64 string, key []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(data) < blockSize {
		return nil, errors.New("密文长度不足")
	}

	iv := data[:blockSize]
	ciphertext := data[blockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}

// Crypto_AES_OFB_Encrypt 使用 AES-OFB 模式加密数据。
// OFB 模式为流加密模式，不需要填充。
// 自动生成随机 IV，格式为 IV + 密文。
//
// 参数:
//   - plaintext: 待加密的原始数据
//   - key: AES 密钥
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func Crypto_AES_OFB_Encrypt(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := make([]byte, block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	stream := cipher.NewOFB(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	result := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(result), nil
}

// Crypto_AES_OFB_Decrypt AES-OFB 模式解密。
//
// 参数:
//   - cipherBase64: Base64 编码的密文
//   - key: AES 密钥
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func Crypto_AES_OFB_Decrypt(cipherBase64 string, key []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(data) < blockSize {
		return nil, errors.New("密文长度不足")
	}

	iv := data[:blockSize]
	ciphertext := data[blockSize:]

	stream := cipher.NewOFB(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}

// ============================================================
// DES / 3DES 加解密
// ============================================================

// Crypto_DES_CBC_Encrypt 使用 DES-CBC 模式加密数据。
// DES 密钥长度为 8 字节，IV 长度为 8 字节。
//
// 参数:
//   - plaintext: 待加密的原始数据
//   - key: DES 密钥（8 字节）
//   - iv: 初始化向量（8 字节）
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func Crypto_DES_CBC_Encrypt(plaintext, key, iv []byte) (string, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext = pkcs7Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Crypto_DES_CBC_Decrypt DES-CBC 模式解密。
//
// 参数:
//   - cipherBase64: Base64 编码的密文
//   - key: DES 密钥
//   - iv: 初始化向量
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func Crypto_DES_CBC_Decrypt(cipherBase64 string, key, iv []byte) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return nil, err
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = pkcs7Unpadding(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// Crypto_3DES_CBC_Encrypt 使用 3DES-CBC 模式加密数据。
// 3DES 密钥长度为 24 字节，IV 长度为 8 字节。
//
// 参数:
//   - plaintext: 待加密的原始数据
//   - key: 3DES 密钥（24 字节）
//   - iv: 初始化向量（8 字节）
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func Crypto_3DES_CBC_Encrypt(plaintext, key, iv []byte) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}

	plaintext = pkcs7Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Crypto_3DES_CBC_Decrypt 3DES-CBC 模式解密。
//
// 参数:
//   - cipherBase64: Base64 编码的密文
//   - key: 3DES 密钥
//   - iv: 初始化向量
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func Crypto_3DES_CBC_Decrypt(cipherBase64 string, key, iv []byte) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return nil, err
	}

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = pkcs7Unpadding(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// ============================================================
// 辅助函数：PKCS7 填充
// ============================================================

// pkcs7Padding 对数据进行 PKCS7 填充。
// 填充规则：若要填充 N 个字节（0 < N < 256），则每个填充字节的值均为 N。
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := make([]byte, padding)
	for i := range padText {
		padText[i] = byte(padding)
	}
	return append(data, padText...)
}

// pkcs7Unpadding 去除 PKCS7 填充。
func pkcs7Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("数据为空")
	}
	padding := int(data[length-1])
	if padding > length {
		return nil, errors.New("填充无效")
	}
	return data[:length-padding], nil
}