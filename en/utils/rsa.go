package utils

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
	"strings"
)

// RSA_SignWithPrivateKey 使用 RSA 私钥对数据进行 MD5 签名。
// 签名流程：对明文取 MD5 哈希 → 使用 PKCS1v15 填充方案签名 → 返回大写十六进制字符串。
//
// 参数:
//   - plaintext: 待签名的明文数据（Base64 编码前的原始文本）
//   - privateKeyPEM: PEM 格式的 RSA 私钥字符串
//
// 返回:
//   - string: 大写十六进制格式的签名值；签名失败返回空串
func RSA_SignWithPrivateKey(plaintext string, privateKeyPEM string) string {
	pemKey := []byte(privateKeyPEM)

	data := []byte(plaintext)
	hashMd5 := md5.Sum(data)
	hashed := hashMd5[:]
	block, _ := pem.Decode(pemKey)
	if block == nil {
		return ""
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return ""
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.MD5, hashed)
	if err != nil {
		return ""
	}
	return strings.ToUpper(hex.EncodeToString(signature))
}

// RSA_GenerateKey 生成 1024 位的 RSA 公钥/私钥对。
// 返回 PEM 格式的公钥和私钥字符串。
//
// 返回:
//   - err: 生成失败时的错误信息
//   - publicKey: PEM 格式的 PKCS8 公钥
//   - privateKey: PEM 格式的 PKCS1 私钥
func RSA_GenerateKey() (err error, publicKey string, privateKey string) {
	bits := 1024
	privKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err, "", ""
	}

	privateKeyStream := x509.MarshalPKCS1PrivateKey(privKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyStream,
	}
	privateKeyStr := string(pem.EncodeToMemory(block))
	pubKey := &privKey.PublicKey
	publicKeyStream, _ := x509.MarshalPKIXPublicKey(pubKey)
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyStream,
	}
	publicKeyStr := string(pem.EncodeToMemory(block))
	return err, publicKeyStr, privateKeyStr
}

// RSA_DecryptWithPrivateKey 使用 RSA 私钥解密数据，返回解密后的字符串。
// 内部调用 RSA_DecryptWithPrivateKey2 并将结果转换为字符串。
//
// 参数:
//   - privateKeyBytes: PEM 格式的 RSA 私钥字节集
//   - encryptedData: 待解密的密文字节集
//
// 返回:
//   - string: 解密后的明文字符串；解密失败返回空串
func RSA_DecryptWithPrivateKey(privateKeyBytes []byte, encryptedData []byte) string {
	return string(RSA_DecryptWithPrivateKey2(privateKeyBytes, encryptedData))
}

// RSA_DecryptWithPrivateKey2 使用 RSA 私钥解密数据，返回解密后的字节集。
// 使用 PKCS1v15 填充方案进行解密。
//
// 参数:
//   - privateKeyBytes: PEM 格式的 RSA 私钥字节集
//   - encryptedData: 待解密的密文字节集
//
// 返回:
//   - []byte: 解密后的明文字节集；解密失败返回空字节集
func RSA_DecryptWithPrivateKey2(privateKeyBytes []byte, encryptedData []byte) []byte {
	if len(encryptedData) == 0 || len(privateKeyBytes) == 0 {
		return []byte{}
	}

	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		fmt.Printf("私钥载入失败可能格式不正确:%s\n", string(privateKeyBytes))
		return []byte{}
	}
	derText := block.Bytes
	privateKey, err := x509.ParsePKCS1PrivateKey(derText)
	if err != nil {
		return []byte{}
	}
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedData)

	if err != nil {
		return []byte{}
	}
	return plainText
}

// RSA_EncryptWithPublicKey 使用 RSA 公钥加密数据，返回 Base64 编码的密文。
// 支持解析 PKCS8 格式的公钥（最常见的公钥格式）。
//
// 参数:
//   - publicKeyPEM: PEM 格式的 RSA 公钥字符串
//   - plaintext: 待加密的明文字节集
//
// 返回:
//   - base64Cipher: Base64 编码的加密结果；加密失败返回空串
func RSA_EncryptWithPublicKey(publicKeyPEM string, plaintext []byte) (base64Cipher string) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		fmt.Printf("Rsa公钥加密公钥载入失败可能格式不正确:%s\n", publicKeyPEM)
		return ""
	}
	derText := block.Bytes

	publicKeyInterface, err := x509.ParsePKIXPublicKey(derText)
	if err != nil {
		return ""
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	cipherData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	if err != nil {
		return
	}
	base64Cipher = base64.StdEncoding.EncodeToString(cipherData)
	return base64Cipher
}

// RSA_EncryptWithPrivateKey 使用 RSA 私钥进行"加密"（实际是签名操作）。
// 使用 PKCS1v15 填充方案，crypto.Hash(0) 表示不进行哈希预处理。
// 注意：私钥加密通常用于签名场景，而非保密场景。
//
// 参数:
//   - privateKeyBytes: PEM 格式的 RSA 私钥字节集
//   - plaintext: 待加密的明文字节集
//
// 返回:
//   - base64Cipher: Base64 编码的加密结果；加密失败返回空串
func RSA_EncryptWithPrivateKey(privateKeyBytes []byte, plaintext []byte) (base64Cipher string) {
	if len(plaintext) == 0 || len(privateKeyBytes) == 0 {
		return ""
	}

	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		fmt.Printf("RSA私钥加密Rsa私钥载入失败可能格式不正确:%s\n", string(privateKeyBytes))
		return ""
	}
	derText := block.Bytes
	privateKey, err := x509.ParsePKCS1PrivateKey(derText)

	signData, err := rsa.SignPKCS1v15(nil, privateKey, crypto.Hash(0), plaintext)
	if err != nil {
		return ""
	}
	base64Cipher = base64.StdEncoding.EncodeToString(signData)
	return base64Cipher
}

// RSA_DecryptWithPublicKey 使用 RSA 公钥解密由私钥"加密"的数据。
// 这是 RSA_EncryptWithPrivateKey 的逆操作，使用自定义的 publicDecrypt 实现。
// 支持解析 PKCS8 格式的公钥。
//
// 参数:
//   - publicKeyPEM: PEM 格式的 RSA 公钥字符串
//   - ciphertext: 待解密的密文字节集
//
// 返回:
//   - plaintextBytes: 解密后的明文字节集；解密失败返回空字节集
func RSA_DecryptWithPublicKey(publicKeyPEM string, ciphertext []byte) (plaintextBytes []byte) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		fmt.Printf("RSA公钥解密Rsa公钥载入失败可能格式不正确:%s\n", publicKeyPEM)
		return []byte{}
	}
	derText := block.Bytes

	publicKeyInterface, err := x509.ParsePKIXPublicKey(derText)
	if err != nil {
		return plaintextBytes
	}

	publicKey := publicKeyInterface.(*rsa.PublicKey)

	plaintextBytes, err = publicDecrypt(publicKey, crypto.Hash(0), nil, ciphertext)
	if err != nil {
		return plaintextBytes
	}
	return plaintextBytes
}

// publicDecrypt RSA 公钥"解密"的底层实现（实际是签名验证运算的模拟）。
// 复制自 Go 标准库 crypto/rsa/pkcs1v15.go 的公开验证逻辑。
func publicDecrypt(pub *rsa.PublicKey, hash crypto.Hash, hashed, sig []byte) ([]byte, error) {
	hashLen, prefix, err := pkcs1v15HashInfo(hash, len(hashed))
	if err != nil {
		return nil, err
	}

	tLen := len(prefix) + hashLen
	k := pub.Size()
	if k < tLen+11 {
		return nil, rsa.ErrVerification
	}

	c := new(big.Int).SetBytes(sig)
	m := new(big.Int)
	e := big.NewInt(int64(pub.E))
	m.Exp(c, e, pub.N)

	if m.Sign() <= 0 {
		return nil, rsa.ErrDecryption
	}

	em := m.Bytes()
	if len(em) > k {
		return nil, rsa.ErrDecryption
	}

	return em, nil
}

// ============================================================
// 以下函数复制自 Go 标准库 crypto/rsa/pkcs1v15.go
// 由于 Go 标准库未公开这些函数，需要手动复制
// ============================================================

// hashPrefixes 各哈希算法的 ASN.1 前缀，用于 PKCS1v15 签名验证。
var hashPrefixes = map[crypto.Hash][]byte{
	crypto.MD5:    {0x30, 0x20, 0x30, 0x0c, 0x06, 0x08, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x02, 0x05, 0x05, 0x00, 0x04, 0x10},
	crypto.SHA1:   {0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14},
	crypto.SHA224: {0x30, 0x2d, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x04, 0x05, 0x00, 0x04, 0x1c},
	crypto.SHA256: {0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20},
	crypto.SHA384: {0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30},
	crypto.SHA512: {0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40},
}

// encrypt RSA 公钥加密的底层实现，执行模幂运算 c = m^e mod n。
func rsaEncrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

// pkcs1v15HashInfo 获取指定哈希算法的哈希长度和 ASN.1 前缀。
func pkcs1v15HashInfo(hash crypto.Hash, inLen int) (hashLen int, prefix []byte, err error) {
	if hash == 0 {
		return inLen, nil, nil
	}

	prefix, ok := hashPrefixes[hash]
	if !ok {
		return 0, nil, fmt.Errorf("未知的哈希算法: %v", hash)
	}
	return hash.Size(), prefix, nil
}