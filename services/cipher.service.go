package services

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/jadhamwi21/passwords-manager-cli-tool/utils"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

type cipher_service_interface interface {
	Encrypt(string) string
	Decrypt(string) string
}

type cipher_service struct {
	secret string
}

func (CipherService cipher_service) Encrypt(text string) string {
	block, err := aes.NewCipher([]byte(CipherService.secret))
	if err != nil {
		panic(err)
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return utils.Encode(cipherText)
}

func (CipherService cipher_service) Decrypt(text string) string {
	block, err := aes.NewCipher([]byte(CipherService.secret))
	if err != nil {
		panic(err)
	}
	cipherText := utils.Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText)
}

func newCipherService() cipher_service_interface {
	SecretService := newSecretService()
	secret := SecretService.ResolveSecret()

	return &cipher_service{secret: secret}

}

var CipherService cipher_service_interface = newCipherService()
