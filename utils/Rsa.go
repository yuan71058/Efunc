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
	"errors"
	"fmt"
	"math/big"
	"strings"
)

// Rsa_私钥签名 使用 RSA 私钥对数据进行 MD5 签名。
// 签名流程：对明文取 MD5 哈希 → 使用 PKCS1v15 填充方案签名 → 返回大写十六进制字符串。
//
// 参数:
//   - base64后明文: 待签名的明文数据（Base64 编码前的原始文本）
//   - RSA私钥: PEM 格式的 RSA 私钥字符串
//
// 返回:
//   - string: 大写十六进制格式的签名值；签名失败返回空串
func Rsa_私钥签名(base64后明文 string, RSA私钥 string) string {

	pemKey := []byte(RSA私钥)

	data := []byte(base64后明文)
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

// Rsa_GetKey 生成 1024 位的 RSA 公钥/私钥对。
// 返回 PEM 格式的公钥和私钥字符串。
//
// 返回:
//   - error: 生成失败时的错误信息
//   - PublicKey: PEM 格式的 PKCS8 公钥
//   - PrivateKey: PEM 格式的 PKCS1 私钥
func Rsa_GetKey() (err error, PublicKey string, PrivateKey string) {
	bits := 1024
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err, "", ""
	}

	privateKeyStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyStream,
	}
	privateKeyStr := string(pem.EncodeToMemory(block))
	publicKey := &privateKey.PublicKey
	publicKeyStream, _ := x509.MarshalPKIXPublicKey(publicKey)
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyStream,
	}
	publicKeyStr := string(pem.EncodeToMemory(block))
	return err, publicKeyStr, privateKeyStr
}

// Rsa_私钥解密 使用 RSA 私钥解密数据，返回解密后的字符串。
// 内部调用 Rsa_私钥解密2 并将结果转换为字符串。
//
// 参数:
//   - Rsa私钥: PEM 格式的 RSA 私钥字节集
//   - 加密数据: 待解密的密文字节集
//
// 返回:
//   - string: 解密后的明文字符串；解密失败返回空串
func Rsa_私钥解密(Rsa私钥 []byte, 加密数据 []byte) string {
	return string(Rsa_私钥解密2(Rsa私钥, 加密数据))
}

// Rsa_私钥解密2 使用 RSA 私钥解密数据，返回解密后的字节集。
// 使用 PKCS1v15 填充方案进行解密。
//
// 参数:
//   - Rsa私钥: PEM 格式的 RSA 私钥字节集
//   - 加密数据: 待解密的密文字节集
//
// 返回:
//   - []byte: 解密后的明文字节集；解密失败返回空字节集
func Rsa_私钥解密2(Rsa私钥 []byte, 加密数据 []byte) []byte {

	if len(加密数据) == 0 || len(Rsa私钥) == 0 {
		return []byte{}
	}

	block, _ := pem.Decode(Rsa私钥)
	if block == nil {
		fmt.Printf("私钥载入失败可能格式不正确:%s\n", string(Rsa私钥))
		return []byte{}
	}
	derText := block.Bytes
	privateKey, err := x509.ParsePKCS1PrivateKey(derText)
	if err != nil {
		return []byte{}
	}
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, 加密数据)

	if err != nil {
		return []byte{}
	}
	return plainText

}

// Rsa_公钥加密 使用 RSA 公钥加密数据，返回 Base64 编码的密文。
// 支持解析 PKCS8 格式的公钥（最常见的公钥格式）。
//
// 参数:
//   - 公钥: PEM 格式的 RSA 公钥字符串
//   - 加密内容: 待加密的明文字节集
//
// 返回:
//   - base64密文: Base64 编码的加密结果；加密失败返回空串
func Rsa_公钥加密(公钥 string, 加密内容 []byte) (base64密文 string) {
	block, _ := pem.Decode([]byte(公钥))
	if block == nil {
		fmt.Printf("Rsa公钥加密公钥载入失败可能格式不正确:%s\n", 公钥)
		return ""
	}
	derText := block.Bytes

	publicKeyInterface, err := x509.ParsePKIXPublicKey(derText)
	if err != nil {
		return ""
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	cipherData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, 加密内容)
	if err != nil {
		return
	}
	base64密文 = base64.StdEncoding.EncodeToString(cipherData)
	return base64密文

}

// RSA_私钥加密 使用 RSA 私钥进行"加密"（实际是签名操作）。
// 使用 PKCS1v15 填充方案，crypto.Hash(0) 表示不进行哈希预处理。
// 注意：私钥加密通常用于签名场景，而非保密场景。
//
// 参数:
//   - Rsa私钥: PEM 格式的 RSA 私钥字节集
//   - 明文: 待加密的明文字节集
//
// 返回:
//   - base64密文: Base64 编码的加密结果；加密失败返回空串
func RSA_私钥加密(Rsa私钥 []byte, 明文 []byte) (base64密文 string) {
	if len(明文) == 0 || len(Rsa私钥) == 0 {
		return ""
	}

	block, _ := pem.Decode(Rsa私钥)
	if block == nil {
		fmt.Printf("RSA私钥加密Rsa私钥载入失败可能格式不正确:%s\n", string(Rsa私钥))
		return ""
	}
	derText := block.Bytes
	privateKey, err := x509.ParsePKCS1PrivateKey(derText)

	signData, err := rsa.SignPKCS1v15(nil, privateKey, crypto.Hash(0), 明文)
	if err != nil {
		return ""
	}
	base64密文 = base64.StdEncoding.EncodeToString(signData)
	return base64密文
}

// RSA_公钥解密 使用 RSA 公钥解密由私钥"加密"的数据。
// 这是 RSA_私钥加密 的逆操作，使用自定义的 publicDecrypt 实现。
// 支持解析 PKCS8 格式的公钥。
//
// 参数:
//   - 公钥: PEM 格式的 RSA 公钥字符串
//   - 密文: 待解密的密文字节集
//
// 返回:
//   - 明文字节集: 解密后的明文字节集；解密失败返回空字节集
func RSA_公钥解密(公钥 string, 密文 []byte) (明文字节集 []byte) {

	block, _ := pem.Decode([]byte(公钥))
	if block == nil {
		fmt.Printf("RSA公钥解密Rsa公钥载入失败可能格式不正确:%s\n", 公钥)
		return []byte{}
	}
	derText := block.Bytes

	publicKey, err := x509.ParsePKCS1PublicKey(derText)

	publicKeyInterface, err := x509.ParsePKIXPublicKey(derText)
	if err != nil {
		return 明文字节集
	}

	publicKey = publicKeyInterface.(*rsa.PublicKey)

	明文字节集, err = publicDecrypt(publicKey, crypto.Hash(0), nil, 密文)
	if err != nil {
		return 明文字节集
	}
	return 明文字节集
}

// hashPrefixes 各哈希算法的 ASN.1 前缀，用于 PKCS1v15 签名验证。
// 复制自 Go 标准库 crypto/rsa/pkcs1v5.go。
var hashPrefixes = map[crypto.Hash][]byte{
	crypto.MD5:       {0x30, 0x20, 0x30, 0x0c, 0x06, 0x08, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x02, 0x05, 0x05, 0x00, 0x04, 0x10},
	crypto.SHA1:      {0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14},
	crypto.SHA224:    {0x30, 0x2d, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x04, 0x05, 0x00, 0x04, 0x1c},
	crypto.SHA256:    {0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20},
	crypto.SHA384:    {0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30},
	crypto.SHA512:    {0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40},
	crypto.MD5SHA1:   {},
	crypto.RIPEMD160: {0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14},
}

// encrypt RSA 公钥加密的底层实现，执行模幂运算 c = m^e mod n。
// 复制自 Go 标准库 crypto/rsa/pkcs1v5.go。
func encrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

// pkcs1v15HashInfo 获取指定哈希算法的哈希长度和 ASN.1 前缀。
// 复制自 Go 标准库 crypto/rsa/pkcs1v5.go。
func pkcs1v15HashInfo(hash crypto.Hash, inLen int) (hashLen int, prefix []byte, err error) {
	if hash == 0 {
		return inLen, nil, nil
	}

	hashLen = hash.Size()
	if inLen != hashLen {
		return 0, nil, errors.New("crypto/rsa: input must be hashed message")
	}
	prefix, ok := hashPrefixes[hash]
	if !ok {
		return 0, nil, errors.New("crypto/rsa: unsupported hash function")
	}
	return
}

// leftPad 将输入字节集左侧填充零至指定长度。
// 复制自 Go 标准库 crypto/rsa/pkcs1v5.go。
func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

// unLeftPad 去除 PKCS1v15 填充，提取实际数据。
// 复制自 Go 标准库 crypto/rsa/pkcs1v5.go（已修改）。
func unLeftPad(input []byte) (out []byte) {
	n := len(input)
	t := 2
	for i := 2; i < n; i++ {
		if input[i] == 0xff {
			t = t + 1
		} else {
			if input[i] == input[0] {
				t = t + int(input[1])
			}
			break
		}
	}
	out = make([]byte, n-t)
	copy(out, input[t:])
	return
}

// publicDecrypt 使用 RSA 公钥进行"解密"操作（私钥加密的逆操作）。
// 复制并修改自 Go 标准库 crypto/rsa/pkcs1v5.go。
// 此函数实现了非标准的 RSA 公钥解密流程，用于配合 RSA_私钥加密 使用。
func publicDecrypt(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte) (out []byte, err error) {
	hashLen, prefix, err := pkcs1v15HashInfo(hash, len(hashed))
	if err != nil {
		return nil, err
	}

	tLen := len(prefix) + hashLen
	k := (pub.N.BitLen() + 7) / 8
	if k < tLen+11 {
		return nil, fmt.Errorf("length illegal")
	}

	c := new(big.Int).SetBytes(sig)
	m := encrypt(new(big.Int), pub, c)
	em := leftPad(m.Bytes(), k)
	out = unLeftPad(em)

	err = nil
	return
}
