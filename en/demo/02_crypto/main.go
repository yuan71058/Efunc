package main

import (
	"fmt"

	"github.com/yuan71058/Efunc/en/utils"
)

func main() {
	fmt.Println("=== Crypto Demo ===")

	key := []byte("1234567890abcdef")
	iv := []byte("1234567890abcdef")

	// AES CBC
	plaintext := []byte("Hello, Efunc Crypto!")
	encrypted, err := utils.Crypto_AES_CBC_Encrypt(plaintext, key, iv)
	if err != nil {
		fmt.Println("AES encrypt error:", err)
		return
	}
	fmt.Printf("AES encrypted (hex): %x\n", encrypted)

	decrypted, err := utils.Crypto_AES_CBC_Decrypt(encrypted, key, iv)
	if err != nil {
		fmt.Println("AES decrypt error:", err)
		return
	}
	fmt.Printf("AES decrypted: %s\n", decrypted)

	// DES
	desKey := []byte("12345678")
	desIv := []byte("12345678")
	desEnc, _ := utils.Crypto_DES_CBC_Encrypt(plaintext, desKey, desIv)
	desDec, _ := utils.Crypto_DES_CBC_Decrypt(desEnc, desKey, desIv)
	fmt.Printf("DES decrypted: %s\n", desDec)

	// 3DES
	tripleKey := []byte("123456789012345678901234")
	tripleEnc, _ := utils.Crypto_3DES_CBC_Encrypt(plaintext, tripleKey, desIv)
	tripleDec, _ := utils.Crypto_3DES_CBC_Decrypt(tripleEnc, tripleKey, desIv)
	fmt.Printf("3DES decrypted: %s\n", tripleDec)

	// HMAC
	hmac := utils.Crypto_HMAC_SHA256([]byte("message"), key)
	fmt.Printf("HMAC-SHA256 (hex): %x\n", hmac)

	// RSA
	fmt.Println("\n--- RSA ---")
	priv, pub, err := utils.RSA_GenerateKey(2048)
	if err != nil {
		fmt.Println("RSA key gen error:", err)
		return
	}

	cipher, err := utils.RSA_Encrypt([]byte("RSA test message"), pub)
	if err != nil {
		fmt.Println("RSA encrypt error:", err)
		return
	}
	fmt.Printf("RSA encrypted (hex): %x\n", cipher)

	clear, err := utils.RSA_Decrypt(cipher, priv)
	if err != nil {
		fmt.Println("RSA decrypt error:", err)
		return
	}
	fmt.Printf("RSA decrypted: %s\n", clear)

	// Sign & verify
	sig, err := utils.RSA_Sign([]byte("sign this"), priv)
	if err != nil {
		fmt.Println("RSA sign error:", err)
		return
	}
	err = utils.RSA_Verify([]byte("sign this"), sig, pub)
	fmt.Printf("RSA verify: %v\n", err == nil)

	fmt.Println("\nDone!")
}