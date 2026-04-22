package decryptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func aesEncrypt(plain string, key []byte) string {
	block, _ := aes.NewCipher(key[:32])
	iv := key[:aes.BlockSize]
	plainBytes := pkcs7Pad([]byte(plain), aes.BlockSize)
	ciphertext := make([]byte, len(plainBytes))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plainBytes)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func pkcs7Pad(b []byte, blockSize int) []byte {
	pad := blockSize - len(b)%blockSize
	return append(b, bytes.Repeat([]byte{byte(pad)}, pad)...)
}

func aesDecryptCBC(ciphertext, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)
	return plaintext
}

func asErr(msg string, a ...any) error {
	return fmt.Errorf(msg, a...)
}
