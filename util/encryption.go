package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// Encrypt 对数据进行 AES 加密
func Encrypt(plaintext []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	plaintext = pkcs7Padding(plaintext, blockSize)

	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, key[:blockSize])
	mode.CryptBlocks(ciphertext, plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 对加密数据进行 AES 解密
func Decrypt(ciphertext string, key []byte) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(decoded)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, key[:blockSize])
	mode.CryptBlocks(decoded, decoded)

	return pkcs7UnPadding(decoded), nil
}

// pkcs7Padding 填充数据以满足块大小要求
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// pkcs7UnPadding 去除填充数据
func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
