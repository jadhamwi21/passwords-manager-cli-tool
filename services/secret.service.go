package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
)

type secret_service_interface interface {
	ResolveSecret() string
}

type secret_service struct {
	path string
}

func generateSecretKey() string {
	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	return hex.EncodeToString(key)
}

func getSecret(path string) string {
	secret, secretError := os.ReadFile(path)
	if secretError != nil {
		panic(secretError)
	}
	return string(secret)
}

func createSecretFile(path string) {
	secret := generateSecretKey()
	newSecretFile, newSecretFileError := os.Create(path)
	if newSecretFileError != nil {
		panic(newSecretFileError)
	}
	newSecretFile.WriteString(secret)
}

func (SecretService secret_service) ResolveSecret() string {
	workingDir, workingDirErr := os.Getwd()
	if workingDirErr != nil {
		panic(workingDirErr)
	}
	path := filepath.Clean(filepath.Join(workingDir, "secret"))
	_, secretFileError := os.Stat(path)
	if errors.Is(secretFileError, os.ErrNotExist) {
		createSecretFile(path)
		return getSecret(path)
	} else {
		return getSecret(path)
	}
}

func newSecretService() secret_service_interface {
	workingDir, workingDirErr := os.Getwd()
	if workingDirErr != nil {
		panic(workingDirErr)
	}
	path := filepath.Clean(filepath.Join(workingDir, "secret"))

	return &secret_service{path: path}

}
