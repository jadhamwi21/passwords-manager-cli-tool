package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/jadhamwi21/passwords-manager-cli-tool/utils"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

type encryption_service_interface interface {
	Encrypt(string) string
	Decrypt(string) string
}

type encryption_service struct {
	secret string
}

func (EncryptionService encryption_service) Encrypt(text string) string {
	block, err := aes.NewCipher([]byte(EncryptionService.secret))
	if err != nil {
		panic(err)
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return utils.Encode(cipherText)
}

func (EncryptionService encryption_service) Decrypt(text string) string {
	block, err := aes.NewCipher([]byte(EncryptionService.secret))
	if err != nil {
		panic(err)
	}
	cipherText := utils.Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText)
}

func generateSecretKey() string {
	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	return hex.EncodeToString(key)
}

func createSecretFile(path string) {
	secret := generateSecretKey()
	newSecretFile, newSecretFileError := os.Create(path)
	if newSecretFileError != nil {
		panic(newSecretFileError)
	}
	newSecretFile.WriteString(secret)
}

func getSecret(path string) string {
	secret, secretError := os.ReadFile(path)
	if secretError != nil {
		panic(secretError)
	}
	return string(secret)
}

func newEncryptionService() encryption_service_interface {
	workingDir, workingDirErr := os.Getwd()
	if workingDirErr != nil {
		panic(workingDirErr)
	}
	path := filepath.Clean(filepath.Join(workingDir, "secret"))
	_, secretFileError := os.Stat(path)
	if os.IsNotExist(secretFileError) {
		createSecretFile(path)
		secret := getSecret(path)
		return &encryption_service{secret: secret}
	} else {
		secret := getSecret(path)
		return &encryption_service{secret: secret}
	}

}

var EncryptionService encryption_service_interface = newEncryptionService()
