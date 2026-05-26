// Package utils 提供加解密算法常用函数。
// 本文件集成 AES、DES、3DES、RC4、XOR、TEA/XXTEA 等常用对称加解密算法。
// 输出格式统一为 Base64 编码字符串，方便存储与传输。
//
// 注意：
//   - ECB 模式安全性较低，不建议用于生产环境的加密需求，推荐 GCM 或 CBC 模式。
//   - 所有函数均已添加详细注释，包括参数说明、返回值及注意事项。
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

// J加解密_AES_CBC加密 使用 AES-CBC 模式加密数据。
// 采用 PKCS7 填充方案，返回 Base64 编码的密文。
//
// 参数:
//   - 明文: 待加密的原始数据
//   - 密钥: AES 密钥，长度必须为 16（AES-128）、24（AES-192）或 32（AES-256）字节
//   - IV: 初始化向量，长度必须为 16 字节
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func J加解密_AES_CBC加密(明文, 密钥, IV []byte) (string, error) {
	block, err := aes.NewCipher(密钥)
	if err != nil {
		return "", err
	}
	if len(IV) != block.BlockSize() {
		return "", errors.New("IV长度必须为16字节")
	}

	// PKCS7 填充
	明文 = pkcs7Padding(明文, block.BlockSize())

	密文 := make([]byte, len(明文))
	mode := cipher.NewCBCEncrypter(block, IV)
	mode.CryptBlocks(密文, 明文)

	return base64.StdEncoding.EncodeToString(密文), nil
}

// J加解密_AES_CBC解密 AES-CBC 模式解密，输入为 Base64 密文。
//
// 参数:
//   - 密文Base64: Base64 编码的密文文本
//   - 密钥: AES 密钥
//   - IV: 初始化向量
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func J加解密_AES_CBC解密(密文Base64 string, 密钥, IV []byte) ([]byte, error) {
	密文, err := base64.StdEncoding.DecodeString(密文Base64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(密钥)
	if err != nil {
		return nil, err
	}
	if len(IV) != block.BlockSize() {
		return nil, errors.New("IV长度必须为16字节")
	}
	if len(密文)%block.BlockSize() != 0 {
		return nil, errors.New("密文长度不是块大小的整数倍")
	}

	明文 := make([]byte, len(密文))
	mode := cipher.NewCBCDecrypter(block, IV)
	mode.CryptBlocks(明文, 密文)

	// 去除 PKCS7 填充
	明文, err = pkcs7Unpadding(明文)
	if err != nil {
		return nil, err
	}

	return 明文, nil
}

// J加解密_AES_ECB加密 使用 AES-ECB 模式加密数据（不推荐在生产环境使用）。
// ECB 模式不使用 IV，安全性较低，相同明文块会生成相同密文块。
//
// 参数:
//   - 明文: 待加密的原始数据
//   - 密钥: AES 密钥
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func J加解密_AES_ECB加密(明文, 密钥 []byte) (string, error) {
	block, err := aes.NewCipher(密钥)
	if err != nil {
		return "", err
	}

	明文 = pkcs7Padding(明文, block.BlockSize())
	密文 := make([]byte, len(明文))
	块大小 := block.BlockSize()

	for i := 0; i < len(明文); i += 块大小 {
		block.Encrypt(密文[i:i+块大小], 明文[i:i+块大小])
	}

	return base64.StdEncoding.EncodeToString(密文), nil
}

// J加解密_AES_ECB解密 AES-ECB 模式解密。
//
// 参数:
//   - 密文Base64: Base64 编码的密文文本
//   - 密钥: AES 密钥
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func J加解密_AES_ECB解密(密文Base64 string, 密钥 []byte) ([]byte, error) {
	密文, err := base64.StdEncoding.DecodeString(密文Base64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(密钥)
	if err != nil {
		return nil, err
	}
	if len(密文)%block.BlockSize() != 0 {
		return nil, errors.New("密文长度不是块大小的整数倍")
	}

	明文 := make([]byte, len(密文))
	块大小 := block.BlockSize()

	for i := 0; i < len(密文); i += 块大小 {
		block.Decrypt(明文[i:i+块大小], 密文[i:i+块大小])
	}

	明文, err = pkcs7Unpadding(明文)
	if err != nil {
		return nil, err
	}

	return 明文, nil
}

// J加解密_AES_GCM加密 使用 AES-GCM 认证加密模式加密数据。
// GCM 模式同时提供机密性和完整性验证，推荐用于安全敏感场景。
// 自动生成 12 字节随机 nonce，格式为 nonce + 密文。
//
// 参数:
//   - 明文: 待加密的原始数据
//   - 密钥: AES 密钥
//   - 附加数据: 可选，参与认证但不加密的附加数据
//
// 返回:
//   - string: Base64 编码的密文（nonce + 密文 + tag）
//   - error: 加密失败时返回错误信息
func J加解密_AES_GCM加密(明文, 密钥, 附加数据 []byte) (string, error) {
	block, err := aes.NewCipher(密钥)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 生成随机 nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	// GCM 自动附加认证标签
	密文 := gcm.Seal(nonce, nonce, 明文, 附加数据)
	return base64.StdEncoding.EncodeToString(密文), nil
}

// J加解密_AES_GCM解密 AES-GCM 认证解密。
//
// 参数:
//   - 密文Base64: Base64 编码的密文
//   - 密钥: AES 密钥
//   - 附加数据: 加密时传入的附加数据，必须一致
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败或认证失败时返回错误信息
func J加解密_AES_GCM解密(密文Base64 string, 密钥, 附加数据 []byte) ([]byte, error) {
	密文, err := base64.StdEncoding.DecodeString(密文Base64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(密钥)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(密文) < nonceSize {
		return nil, errors.New("密文长度不足")
	}

	nonce, 密文 := 密文[:nonceSize], 密文[nonceSize:]
	return gcm.Open(nil, nonce, 密文, 附加数据)
}

// J加解密_AES_CTR加密 使用 AES-CTR 模式加密数据。
// CTR 模式为流加密模式，不需要填充，适合任意长度数据。
// 自动生成随机 IV，格式为 IV + 密文。
//
// 参数:
//   - 明文: 待加密的原始数据
//   - 密钥: AES 密钥
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func J加解密_AES_CTR加密(明文, 密钥 []byte) (string, error) {
	block, err := aes.NewCipher(密钥)
	if err != nil {
		return "", err
	}

	iv := make([]byte, block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	stream := cipher.NewCTR(block, iv)
	密文 := make([]byte, len(明文))
	stream.XORKeyStream(密文, 明文)

	// 将 IV 追加到密文头部
	result := make([]byte, len(iv)+len(密文))
	copy(result, iv)
	copy(result[len(iv):], 密文)

	return base64.StdEncoding.EncodeToString(result), nil
}

// J加解密_AES_CTR解密 AES-CTR 模式解密。
//
// 参数:
//   - 密文Base64: Base64 编码的密文
//   - 密钥: AES 密钥
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func J加解密_AES_CTR解密(密文Base64 string, 密钥 []byte) ([]byte, error) {
	密文, err := base64.StdEncoding.DecodeString(密文Base64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(密钥)
	if err != nil {
		return nil, err
	}

	块大小 := block.BlockSize()
	if len(密文) < 块大小 {
		return nil, errors.New("密文长度不足")
	}

	iv := 密文[:块大小]
	密文 = 密文[块大小:]

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(密文, 密文)

	return 密文, nil
}

// J加解密_AES_CFB加密 使用 AES-CFB 模式加密数据。
// CFB 模式为流加密模式，不需要填充，适合逐字节加密场景。
// 自动生成随机 IV，格式为 IV + 密文。
//
// 参数:
//   - 明文: 待加密的原始数据
//   - 密钥: AES 密钥
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func J加解密_AES_CFB加密(明文, 密钥 []byte) (string, error) {
	block, err := aes.NewCipher(密钥)
	if err != nil {
		return "", err
	}

	iv := make([]byte, block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	密文 := make([]byte, len(明文))
	stream.XORKeyStream(密文, 明文)

	result := make([]byte, len(iv)+len(密文))
	copy(result, iv)
	copy(result[len(iv):], 密文)

	return base64.StdEncoding.EncodeToString(result), nil
}

// J加解密_AES_CFB解密 AES-CFB 模式解密。
//
// 参数:
//   - 密文Base64: Base64 编码的密文
//   - 密钥: AES 密钥
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func J加解密_AES_CFB解密(密文Base64 string, 密钥 []byte) ([]byte, error) {
	密文, err := base64.StdEncoding.DecodeString(密文Base64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(密钥)
	if err != nil {
		return nil, err
	}

	块大小 := block.BlockSize()
	if len(密文) < 块大小 {
		return nil, errors.New("密文长度不足")
	}

	iv := 密文[:块大小]
	密文 = 密文[块大小:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(密文, 密文)

	return 密文, nil
}

// J加解密_AES_OFB加密 使用 AES-OFB 模式加密数据。
// OFB 将分组密码转换为同步流密码，比特错误不会传播。
// 自动生成随机 IV，格式为 IV + 密文。
//
// 参数:
//   - 明文: 待加密的原始数据
//   - 密钥: AES 密钥
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func J加解密_AES_OFB加密(明文, 密钥 []byte) (string, error) {
	block, err := aes.NewCipher(密钥)
	if err != nil {
		return "", err
	}

	iv := make([]byte, block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	stream := cipher.NewOFB(block, iv)
	密文 := make([]byte, len(明文))
	stream.XORKeyStream(密文, 明文)

	result := make([]byte, len(iv)+len(密文))
	copy(result, iv)
	copy(result[len(iv):], 密文)

	return base64.StdEncoding.EncodeToString(result), nil
}

// J加解密_AES_OFB解密 AES-OFB 模式解密。
//
// 参数:
//   - 密文Base64: Base64 编码的密文
//   - 密钥: AES 密钥
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func J加解密_AES_OFB解密(密文Base64 string, 密钥 []byte) ([]byte, error) {
	// OFB 解密与加密流程相同，直接复用加密逻辑
	密文, err := base64.StdEncoding.DecodeString(密文Base64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(密钥)
	if err != nil {
		return nil, err
	}

	块大小 := block.BlockSize()
	if len(密文) < 块大小 {
		return nil, errors.New("密文长度不足")
	}

	iv := 密文[:块大小]
	密文 = 密文[块大小:]

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(密文, 密文)

	return 密文, nil
}

// ============================================================
// DES 加解密
// ============================================================

// J加解密_DES_CBC加密 使用 DES-CBC 模式加密数据。
// DES 密钥长度固定为 8 字节（56 位有效密钥）。
// 采用 PKCS7 填充方案，返回 Base64 编码的密文。
//
// 参数:
//   - 明文: 待加密的原始数据
//   - 密钥: DES 密钥，长度必须为 8 字节
//   - IV: 初始化向量，长度必须为 8 字节
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func J加解密_DES_CBC加密(明文, 密钥, IV []byte) (string, error) {
	block, err := des.NewCipher(密钥)
	if err != nil {
		return "", err
	}
	if len(IV) != block.BlockSize() {
		return "", errors.New("IV长度必须为8字节")
	}

	明文 = pkcs7Padding(明文, block.BlockSize())
	密文 := make([]byte, len(明文))
	mode := cipher.NewCBCEncrypter(block, IV)
	mode.CryptBlocks(密文, 明文)

	return base64.StdEncoding.EncodeToString(密文), nil
}

// J加解密_DES_CBC解密 DES-CBC 模式解密。
//
// 参数:
//   - 密文Base64: Base64 编码的密文
//   - 密钥: DES 密钥
//   - IV: 初始化向量
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func J加解密_DES_CBC解密(密文Base64 string, 密钥, IV []byte) ([]byte, error) {
	密文, err := base64.StdEncoding.DecodeString(密文Base64)
	if err != nil {
		return nil, err
	}

	block, err := des.NewCipher(密钥)
	if err != nil {
		return nil, err
	}
	if len(IV) != block.BlockSize() {
		return nil, errors.New("IV长度必须为8字节")
	}
	if len(密文)%block.BlockSize() != 0 {
		return nil, errors.New("密文长度不是块大小的整数倍")
	}

	明文 := make([]byte, len(密文))
	mode := cipher.NewCBCDecrypter(block, IV)
	mode.CryptBlocks(明文, 密文)

	明文, err = pkcs7Unpadding(明文)
	if err != nil {
		return nil, err
	}

	return 明文, nil
}

// J加解密_DES_ECB加密 使用 DES-ECB 模式加密数据。
// ECB 模式不使用 IV，安全性较低，不建议在生产环境使用。
//
// 参数:
//   - 明文: 待加密的原始数据
//   - 密钥: DES 密钥（8 字节）
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func J加解密_DES_ECB加密(明文, 密钥 []byte) (string, error) {
	block, err := des.NewCipher(密钥)
	if err != nil {
		return "", err
	}

	明文 = pkcs7Padding(明文, block.BlockSize())
	密文 := make([]byte, len(明文))
	块大小 := block.BlockSize()

	for i := 0; i < len(明文); i += 块大小 {
		block.Encrypt(密文[i:i+块大小], 明文[i:i+块大小])
	}

	return base64.StdEncoding.EncodeToString(密文), nil
}

// J加解密_DES_ECB解密 DES-ECB 模式解密。
//
// 参数:
//   - 密文Base64: Base64 编码的密文
//   - 密钥: DES 密钥
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func J加解密_DES_ECB解密(密文Base64 string, 密钥 []byte) ([]byte, error) {
	密文, err := base64.StdEncoding.DecodeString(密文Base64)
	if err != nil {
		return nil, err
	}

	block, err := des.NewCipher(密钥)
	if err != nil {
		return nil, err
	}
	if len(密文)%block.BlockSize() != 0 {
		return nil, errors.New("密文长度不是块大小的整数倍")
	}

	明文 := make([]byte, len(密文))
	块大小 := block.BlockSize()

	for i := 0; i < len(密文); i += 块大小 {
		block.Decrypt(明文[i:i+块大小], 密文[i:i+块大小])
	}

	明文, err = pkcs7Unpadding(明文)
	if err != nil {
		return nil, err
	}

	return 明文, nil
}

// ============================================================
// 3DES（Triple DES）加解密
// ============================================================

// J加解密_3DES_CBC加密 使用 3DES-CBC 模式加密数据。
// 3DES 对每个数据块应用三次 DES 加密，密钥长度固定为 24 字节（168 位有效密钥）。
// 采用 PKCS7 填充方案，返回 Base64 编码的密文。
//
// 参数:
//   - 明文: 待加密的原始数据
//   - 密钥: 3DES 密钥，长度必须为 24 字节
//   - IV: 初始化向量，长度必须为 8 字节
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func J加解密_3DES_CBC加密(明文, 密钥, IV []byte) (string, error) {
	block, err := des.NewTripleDESCipher(密钥)
	if err != nil {
		return "", err
	}
	if len(IV) != block.BlockSize() {
		return "", errors.New("IV长度必须为8字节")
	}

	明文 = pkcs7Padding(明文, block.BlockSize())
	密文 := make([]byte, len(明文))
	mode := cipher.NewCBCEncrypter(block, IV)
	mode.CryptBlocks(密文, 明文)

	return base64.StdEncoding.EncodeToString(密文), nil
}

// J加解密_3DES_CBC解密 3DES-CBC 模式解密。
//
// 参数:
//   - 密文Base64: Base64 编码的密文
//   - 密钥: 3DES 密钥（24 字节）
//   - IV: 初始化向量
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func J加解密_3DES_CBC解密(密文Base64 string, 密钥, IV []byte) ([]byte, error) {
	密文, err := base64.StdEncoding.DecodeString(密文Base64)
	if err != nil {
		return nil, err
	}

	block, err := des.NewTripleDESCipher(密钥)
	if err != nil {
		return nil, err
	}
	if len(IV) != block.BlockSize() {
		return nil, errors.New("IV长度必须为8字节")
	}
	if len(密文)%block.BlockSize() != 0 {
		return nil, errors.New("密文长度不是块大小的整数倍")
	}

	明文 := make([]byte, len(密文))
	mode := cipher.NewCBCDecrypter(block, IV)
	mode.CryptBlocks(明文, 密文)

	明文, err = pkcs7Unpadding(明文)
	if err != nil {
		return nil, err
	}

	return 明文, nil
}

// J加解密_3DES_ECB加密 使用 3DES-ECB 模式加密数据。
// ECB 模式不使用 IV，安全性较低。
//
// 参数:
//   - 明文: 待加密的原始数据
//   - 密钥: 3DES 密钥（24 字节）
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 加密失败时返回错误信息
func J加解密_3DES_ECB加密(明文, 密钥 []byte) (string, error) {
	block, err := des.NewTripleDESCipher(密钥)
	if err != nil {
		return "", err
	}

	明文 = pkcs7Padding(明文, block.BlockSize())
	密文 := make([]byte, len(明文))
	块大小 := block.BlockSize()

	for i := 0; i < len(明文); i += 块大小 {
		block.Encrypt(密文[i:i+块大小], 明文[i:i+块大小])
	}

	return base64.StdEncoding.EncodeToString(密文), nil
}

// J加解密_3DES_ECB解密 3DES-ECB 模式解密。
//
// 参数:
//   - 密文Base64: Base64 编码的密文
//   - 密钥: 3DES 密钥
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func J加解密_3DES_ECB解密(密文Base64 string, 密钥 []byte) ([]byte, error) {
	密文, err := base64.StdEncoding.DecodeString(密文Base64)
	if err != nil {
		return nil, err
	}

	block, err := des.NewTripleDESCipher(密钥)
	if err != nil {
		return nil, err
	}
	if len(密文)%block.BlockSize() != 0 {
		return nil, errors.New("密文长度不是块大小的整数倍")
	}

	明文 := make([]byte, len(密文))
	块大小 := block.BlockSize()

	for i := 0; i < len(密文); i += 块大小 {
		block.Decrypt(明文[i:i+块大小], 密文[i:i+块大小])
	}

	明文, err = pkcs7Unpadding(明文)
	if err != nil {
		return nil, err
	}

	return 明文, nil
}

// ============================================================
// RC4 加解密（对称流加密，加密和解密使用相同的函数）
// ============================================================

// J加解密_RC4 使用 RC4 算法对数据进行加解密。
// RC4 为对称流加密算法，加密和解密调用同一函数。
// 返回值经过 Base64 编码方便存储。
//
// 参数:
//   - 数据: 待加解密的原始数据
//   - 密钥: RC4 密钥，建议长度 5-16 字节
//
// 返回:
//   - string: Base64 编码的加解密结果
func J加解密_RC4(数据, 密钥 []byte) string {
	s := make([]byte, 256)
	for i := 0; i < 256; i++ {
		s[i] = byte(i)
	}

	// KSA（密钥调度算法）
	var j byte = 0
	for i := 0; i < 256; i++ {
		j = j + s[i] + 密钥[i%len(密钥)]
		s[i], s[j] = s[j], s[i]
	}

	// PRGA（伪随机生成算法）
	result := make([]byte, len(数据))
	var x, y byte = 0, 0
	for k := 0; k < len(数据); k++ {
		x = x + 1
		y = y + s[x]
		s[x], s[y] = s[y], s[x]
		result[k] = 数据[k] ^ s[s[x]+s[y]]
	}

	return base64.StdEncoding.EncodeToString(result)
}

// J加解密_RC4字节集 RC4 加解密，返回原始字节集（不经过 Base64 编码）。
//
// 参数:
//   - 数据: 待加解密的原始数据
//   - 密钥: RC4 密钥
//
// 返回:
//   - []byte: 加解密结果的原始字节集
func J加解密_RC4字节集(数据, 密钥 []byte) []byte {
	s := make([]byte, 256)
	for i := 0; i < 256; i++ {
		s[i] = byte(i)
	}

	var j byte = 0
	for i := 0; i < 256; i++ {
		j = j + s[i] + 密钥[i%len(密钥)]
		s[i], s[j] = s[j], s[i]
	}

	result := make([]byte, len(数据))
	var x, y byte = 0, 0
	for k := 0; k < len(数据); k++ {
		x = x + 1
		y = y + s[x]
		s[x], s[y] = s[y], s[x]
		result[k] = 数据[k] ^ s[s[x]+s[y]]
	}

	return result
}

// ============================================================
// XOR 加解密（简单的异或加解密）
// ============================================================

// J加解密_XOR 使用 XOR 算法对数据进行加解密。
// XOR 是最简单的对称加密算法，加密和解密调用同一函数。
// 密钥会循环使用以匹配数据长度。
//
// 参数:
//   - 数据: 待加解密的原始数据
//   - 密钥: XOR 密钥
//
// 返回:
//   - []byte: 加解密结果的原始字节集
func J加解密_XOR(数据, 密钥 []byte) []byte {
	result := make([]byte, len(数据))
	密钥长度 := len(密钥)
	if 密钥长度 == 0 {
		// 无密钥时原样返回
		copy(result, 数据)
		return result
	}

	for i := range 数据 {
		result[i] = 数据[i] ^ 密钥[i%密钥长度]
	}
	return result
}

// J加解密_XOR文本 XOR 加解密，返回 Base64 编码字符串。
// 适用于字符串加密场景，方便存储和传输。
//
// 参数:
//   - 数据: 待加解密的原始数据
//   - 密钥: XOR 密钥
//
// 返回:
//   - string: Base64 编码的加解密结果
func J加解密_XOR文本(数据, 密钥 []byte) string {
	return base64.StdEncoding.EncodeToString(J加解密_XOR(数据, 密钥))
}

// ============================================================
// TEA（Tiny Encryption Algorithm）加解密
// ============================================================

// J加解密_TEA加密 TEA 加密算法。
// TEA 是一种轻量级分组密码，使用 128 位密钥，性能优秀且实现简单。
// 明文长度必须是 8 的整数倍（不足会自动零填充）。
// 迭代轮数默认 32 轮。
//
// 参数:
//   - 明文: 待加密的原始数据
//   - 密钥: TEA 密钥，长度必须为 16 字节
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 密钥长度不足时返回错误
func J加解密_TEA加密(明文, 密钥 []byte) (string, error) {
	if len(密钥) < 16 {
		return "", errors.New("TEA密钥长度必须至少16字节")
	}

	// 将密钥解析为 4 个 uint32
	k := make([]uint32, 4)
	for i := 0; i < 4; i++ {
		k[i] = uint32(密钥[i*4]) | uint32(密钥[i*4+1])<<8 | uint32(密钥[i*4+2])<<16 | uint32(密钥[i*4+3])<<24
	}

	// 零填充到 8 的整数倍
	data := make([]byte, len(明文))
	copy(data, 明文)
	for len(data)%8 != 0 {
		data = append(data, 0)
	}

	密文 := make([]byte, len(data))
	for i := 0; i < len(data); i += 8 {
		v0 := uint32(data[i]) | uint32(data[i+1])<<8 | uint32(data[i+2])<<16 | uint32(data[i+3])<<24
		v1 := uint32(data[i+4]) | uint32(data[i+5])<<8 | uint32(data[i+6])<<16 | uint32(data[i+7])<<24

		teaEncrypt(&v0, &v1, k)

		// 将加密结果写回字节数组
		putU32LE(密文[i:], v0)
		putU32LE(密文[i+4:], v1)
	}

	return base64.StdEncoding.EncodeToString(密文), nil
}

// J加解密_TEA解密 TEA 解密算法。
//
// 参数:
//   - 密文Base64: Base64 编码的密文
//   - 密钥: TEA 密钥（16 字节）
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func J加解密_TEA解密(密文Base64 string, 密钥 []byte) ([]byte, error) {
	密文, err := base64.StdEncoding.DecodeString(密文Base64)
	if err != nil {
		return nil, err
	}

	if len(密钥) < 16 {
		return nil, errors.New("TEA密钥长度必须至少16字节")
	}
	if len(密文)%8 != 0 {
		return nil, errors.New("密文长度必须是8的整数倍")
	}

	k := make([]uint32, 4)
	for i := 0; i < 4; i++ {
		k[i] = uint32(密钥[i*4]) | uint32(密钥[i*4+1])<<8 | uint32(密钥[i*4+2])<<16 | uint32(密钥[i*4+3])<<24
	}

	// 去除零填充
	data := make([]byte, len(密文))
	for i := 0; i < len(密文); i += 8 {
		v0 := uint32(密文[i]) | uint32(密文[i+1])<<8 | uint32(密文[i+2])<<16 | uint32(密文[i+3])<<24
		v1 := uint32(密文[i+4]) | uint32(密文[i+5])<<8 | uint32(密文[i+6])<<16 | uint32(密文[i+7])<<24

		teaDecrypt(&v0, &v1, k)

		putU32LE(data[i:], v0)
		putU32LE(data[i+4:], v1)
	}

	// 去除尾部零填充
	for len(data) > 0 && data[len(data)-1] == 0 {
		data = data[:len(data)-1]
	}

	return data, nil
}

// ============================================================
// XXTEA（Corrected Block TEA）加解密
// ============================================================

// J加解密_XXTEA加密 XXTEA 加密算法。
// XXTEA 是 TEA 的改进版本，支持任意长度数据，无需填充。
// 密钥长度为 16 字节，加密结果经过 Base64 编码。
//
// 参数:
//   - 明文: 待加密的任意长度数据
//   - 密钥: XXTEA 密钥，长度必须至少 16 字节
//
// 返回:
//   - string: Base64 编码的密文
//   - error: 密钥长度不足时返回错误
func J加解密_XXTEA加密(明文, 密钥 []byte) (string, error) {
	if len(密钥) < 16 {
		return "", errors.New("XXTEA密钥长度必须至少16字节")
	}

	// 将字节集转换为 uint32 数组
	v := bytesToU32LE(明文)

	// 将密钥解析为 uint32 数组
	k := make([]uint32, 4)
	for i := 0; i < 4; i++ {
		k[i] = uint32(密钥[i*4]) | uint32(密钥[i*4+1])<<8 | uint32(密钥[i*4+2])<<16 | uint32(密钥[i*4+3])<<24
	}

	// 调用 XXTEA 核心加密
	v = xxteaEncrypt(v, [4]uint32{k[0], k[1], k[2], k[3]})

	// 将 uint32 数组转换回字节集
	result := u32LEToBytes(v)

	return base64.StdEncoding.EncodeToString(result), nil
}

// J加解密_XXTEA解密 XXTEA 解密算法。
//
// 参数:
//   - 密文Base64: Base64 编码的密文
//   - 密钥: XXTEA 密钥（至少 16 字节）
//
// 返回:
//   - []byte: 解密后的明文字节集
//   - error: 解密失败时返回错误信息
func J加解密_XXTEA解密(密文Base64 string, 密钥 []byte) ([]byte, error) {
	密文, err := base64.StdEncoding.DecodeString(密文Base64)
	if err != nil {
		return nil, err
	}

	if len(密钥) < 16 {
		return nil, errors.New("XXTEA密钥长度必须至少16字节")
	}

	v := bytesToU32LE(密文)

	k := make([]uint32, 4)
	for i := 0; i < 4; i++ {
		k[i] = uint32(密钥[i*4]) | uint32(密钥[i*4+1])<<8 | uint32(密钥[i*4+2])<<16 | uint32(密钥[i*4+3])<<24
	}

	v = xxteaDecrypt(v, [4]uint32{k[0], k[1], k[2], k[3]})
	return u32LEToBytes(v), nil
}

// ============================================================
// AES 密钥/IV 生成工具
// ============================================================

// J加解密_生成AES密钥 生成指定长度的 AES 密钥（16/24/32 字节）。
// 密钥由 crypto/rand 生成，确保安全随机。
//
// 参数:
//   - 长度: 密钥长度，可选 16（AES-128）、24（AES-192）、32（AES-256）
//
// 返回:
//   - []byte: 随机生成的 AES 密钥
//   - error: 生成失败或长度非法时返回错误
func J加解密_生成AES密钥(长度 int) ([]byte, error) {
	if 长度 != 16 && 长度 != 24 && 长度 != 32 {
		return nil, errors.New("AES密钥长度必须为16、24或32字节")
	}

	密钥 := make([]byte, 长度)
	if _, err := rand.Read(密钥); err != nil {
		return nil, err
	}
	return 密钥, nil
}

// J加解密_生成IV 生成指定块大小的随机初始化向量（IV）。
// AES/DES 的 IV 长度等于块大小（AES=16, DES/3DES=8）。
//
// 参数:
//   - 块大小: IV 长度，通常为 8（DES/3DES）或 16（AES）
//
// 返回:
//   - []byte: 随机生成的 IV
//   - error: 生成失败时返回错误
func J加解密_生成IV(块大小 int) ([]byte, error) {
	if 块大小 <= 0 {
		return nil, errors.New("块大小必须大于0")
	}

	iv := make([]byte, 块大小)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	return iv, nil
}

// ============================================================
// 辅助函数（PKCS7 填充 / TEA / XXTEA 内部实现）
// ============================================================

// pkcs7Padding 对数据进行 PKCS7 填充。
// 加密前调用，使数据长度对齐到块大小的整数倍。
// 填充内容为所需填充的字节数，如缺少 3 字节则填充 0x03 0x03 0x03。
//
// 参数:
//   - 数据: 原始数据
//   - 块大小: 加密算法块大小（AES=16, DES=8）
//
// 返回:
//   - []byte: 填充后的数据
func pkcs7Padding(数据 []byte, 块大小 int) []byte {
	填充长度 := 块大小 - len(数据)%块大小
	填充数据 := make([]byte, 填充长度)
	for i := range 填充数据 {
		填充数据[i] = byte(填充长度)
	}
	return append(数据, 填充数据...)
}

// pkcs7Unpadding 去除 PKCS7 填充。
// 解密后调用，自动检测并移除填充字节。
//
// 参数:
//   - 数据: 带填充的数据
//
// 返回:
//   - []byte: 去除填充后的原始数据
//   - error: 填充格式非法时返回错误
func pkcs7Unpadding(数据 []byte) ([]byte, error) {
	if len(数据) == 0 {
		return nil, errors.New("数据为空")
	}

	填充长度 := int(数据[len(数据)-1])
	if 填充长度 > len(数据) || 填充长度 == 0 {
		return nil, errors.New("非法PKCS7填充")
	}

	// 验证填充完整性
	for i := len(数据) - 填充长度; i < len(数据); i++ {
		if 数据[i] != byte(填充长度) {
			return nil, errors.New("PKCS7填充验证失败")
		}
	}

	return 数据[:len(数据)-填充长度], nil
}

// teaEncrypt 执行 TEA 加密的 32 轮 Feistel 网络迭代。
// 内部使用 delta = 0x9E3779B9（黄金分割率派生常量）。
//
// 参数:
//   - v0, v1: 指向两个 32 位数据块的指针（加密后原地修改）
//   - k: 4 个 32 位密钥
func teaEncrypt(v0, v1 *uint32, k []uint32) {
	var sum uint32 = 0
	delta := uint32(0x9E3779B9)

	// 32 轮迭代
	for i := 0; i < 32; i++ {
		sum += delta
		*v0 += ((*v1 << 4) + k[0]) ^ (*v1 + sum) ^ ((*v1 >> 5) + k[1])
		*v1 += ((*v0 << 4) + k[2]) ^ (*v0 + sum) ^ ((*v0 >> 5) + k[3])
	}
}

// teaDecrypt 执行 TEA 解密的 32 轮 Feistel 网络逆迭代。
//
// 参数:
//   - v0, v1: 指向两个 32 位数据块的指针（解密后原地修改）
//   - k: 4 个 32 位密钥
func teaDecrypt(v0, v1 *uint32, k []uint32) {
	delta := uint32(0x9E3779B9)
	sum := delta * 32

	for i := 0; i < 32; i++ {
		*v1 -= ((*v0 << 4) + k[2]) ^ (*v0 + sum) ^ ((*v0 >> 5) + k[3])
		*v0 -= ((*v1 << 4) + k[0]) ^ (*v1 + sum) ^ ((*v1 >> 5) + k[1])
		sum -= delta
	}
}

// xxteaMX XXTEA 核心混合函数。
// 将密钥绑定到数据变换中，确保密钥每位都参与加密。
//
// 参数:
//   - v: 32 位值（会被修改）
//   - k: 密钥数组索引
//   - e, y, z, p: XXTEA 状态变量
//   - sum: 累加器
//
// 返回 32 位混合值
func xxteaMX(v []uint32, k [4]uint32, e, y, z, p uint32, sum uint32) uint32 {
	return ((z>>5 ^ y<<2) + (y>>3 ^ z<<4)) ^ ((sum ^ y) + (k[p&3^e] ^ z))
}

// xxteaEncrypt 执行 XXTEA 加密。
func xxteaEncrypt(v []uint32, k [4]uint32) []uint32 {
	n := uint32(len(v))
	if n < 2 {
		return v
	}

	// XXTEA 使用 n-1 作为变量
	z := v[n-1]
	var sum uint32 = 0
	delta := uint32(0x9E3779B9)

	rounds := uint32(6 + 52/n)
	for q := uint32(0); q < rounds; q++ {
		sum += delta
		e := (sum >> 2) & 3

		var y uint32
		for p := uint32(0); p < n-1; p++ {
			y = v[p+1]
			v[p] += xxteaMX(v, k, e, y, z, p, sum)
			z = v[p]
		}

		y = v[0]
		v[n-1] += xxteaMX(v, k, e, y, z, n-1, sum)
		z = v[n-1]
	}

	return v
}

// xxteaDecrypt 执行 XXTEA 解密。
func xxteaDecrypt(v []uint32, k [4]uint32) []uint32 {
	n := uint32(len(v))
	if n < 2 {
		return v
	}

	z := v[n-1]
	delta := uint32(0x9E3779B9)

	rounds := uint32(6 + 52/n)
	sum := rounds * delta

	for q := uint32(0); q < rounds; q++ {
		e := (sum >> 2) & 3

		var y uint32
		for p := n - 1; p > 0; p-- {
			z = v[p-1]
			y = v[0]
			if p > 1 {
				y = v[p-1]
			} else {
				y = v[n-1]
			}
			v[p] -= xxteaMX(v, k, e, y, z, p, sum)
			z = v[p]
		}

		y = v[n-1]
		v[0] -= xxteaMX(v, k, e, y, z, 0, sum)
		z = v[0]

		sum -= delta
	}

	return v
}

// ============================================================
// 字节集与 uint32 数组转换辅助函数
// ============================================================

// putU32LE 将 uint32 按小端序写入字节切片。
//
// 参数:
//   - buf: 目标字节切片（长度至少 4）
//   - v: 要写入的 uint32 值
func putU32LE(buf []byte, v uint32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
}

// u32LE 从字节切片中按小端序读取 uint32。
//
// 参数:
//   - buf: 源字节切片（长度至少 4）
//
// 返回:
//   - uint32: 读取的值
func u32LE(buf []byte) uint32 {
	return uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24
}

// bytesToU32LE 将字节切片按小端序转换为 uint32 数组。
// 如果长度不是 4 字节对齐，自动末尾补零。
//
// 参数:
//   - data: 原始字节切片
//
// 返回:
//   - []uint32: 转换后的 uint32 数组
func bytesToU32LE(data []byte) []uint32 {
	// 确保是 4 的整数倍
	for len(data)%4 != 0 {
		data = append(data, 0)
	}

	n := len(data) / 4
	result := make([]uint32, n)
	for i := 0; i < n; i++ {
		result[i] = u32LE(data[i*4:])
	}
	return result
}

// u32LEToBytes 将 uint32 数组按小端序转换回字节切片。
//
// 参数:
//   - v: uint32 数组
//
// 返回:
//   - []byte: 转换后的字节切片
func u32LEToBytes(v []uint32) []byte {
	result := make([]byte, len(v)*4)
	for i, val := range v {
		putU32LE(result[i*4:], val)
	}
	return result
}