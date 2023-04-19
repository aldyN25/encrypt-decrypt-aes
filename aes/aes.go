package generalhelper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
)

func GcmEncode(text string) (result string) {
	secretKey := "JayaFgyklmnNbmj2"
	nonce := "JayaViVerCds"

	textByte := []byte(text)

	secretKeyByte := []byte(secretKey)
	block, _ := aes.NewCipher(secretKeyByte)

	nonceByte := []byte(nonce)

	log.Println("PLAIN TEXT : ", text)
	log.Println("SECRET KEY : ", secretKey)
	log.Println("NONCE : ", nonce)
	log.Println("PLAIN TEXT BYTE : ", textByte)
	log.Println("SECRET KEY BYTE : ", secretKeyByte)
	log.Println("NONCE BYTE : ", nonceByte)

	aes, _ := cipher.NewGCM(block)

	cipherText := aes.Seal(nil, nonceByte, textByte, nil)
	clearText := append(nonceByte, cipherText...)

	result = base64.StdEncoding.EncodeToString(clearText)

	return result
}

func GcmDecode(text string) (result string, status int) {
	secretKey := "JayaFgyklmnNbmj2"
	nonce := "JayaViVerCds"

	cipherText, errDec := base64.StdEncoding.DecodeString(text)
	fmt.Println("NNNNN", errDec)
	if errDec != nil {
		return result, 400
	}

	// remove 0 - 11 index (nonce)
	if len(cipherText) < 12 {
		return result, 400
	}
	clearText := cipherText[12:len(cipherText)]

	secretKeyByte := []byte(secretKey)
	block, errBlock := aes.NewCipher(secretKeyByte)

	if errBlock != nil {
		return result, 400
	}

	nonceByte := []byte(nonce)

	log.Println("CIPHER TEXT : ", text)
	log.Println("CIPHER BYTE : ", cipherText)
	log.Println("SECRET KEY : ", secretKey)
	log.Println("SECRET KEY BYTE : ", secretKeyByte)
	log.Println("NONCE : ", nonce)
	log.Println("NONCE BYTE : ", nonceByte)

	aes, errAes := cipher.NewGCM(block)
	if errAes != nil {
		return result, 400
	}

	plainText, errText := aes.Open(nil, nonceByte, clearText, nil)
	if errText != nil {
		return result, 400
	}

	result = string(plainText)

	return result, 200
}
