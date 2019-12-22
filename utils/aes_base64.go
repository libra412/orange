package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

const (
	//偏移量设置
	iv = "0123456789abcdef"
)

func AesBase64Encrypt(origData string, key []byte) string {
	data, _ := AesEncrypt(origData, []byte(key))
	return data
}

func AesBase64Decrypt(origData string, key []byte) string {
	data, _ := AesDecrypt(origData, []byte(key))
	return string(data)
}

func AesEncrypt(encodeStr string, key []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func AesDecrypt(decodeStr string, key []byte) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(decodeBytes))
	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
