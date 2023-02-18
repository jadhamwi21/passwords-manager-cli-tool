package repositories

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
)

type IPasswordsRepository interface {
	AddPassword(id string, password string)
	UpdatePassword(id string, password string)
	RemovePassword(id string)
}

type PasswordsRepository struct {
	passwordsPath string
}

func (passwordsRepository PasswordsRepository) AddPassword(id string, password string) {
	passwordFilePath := fmt.Sprintf("%v/%v.txt", passwordsRepository.passwordsPath, id)
	encryptedPassword, encryptPasswordError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if encryptPasswordError != nil {
		log.Fatal(encryptPasswordError)
	}
	_, err := os.Stat(passwordFilePath)
	if errors.Is(err, os.ErrNotExist) {
		newPasswordFile, _ := os.OpenFile(passwordFilePath, os.O_CREATE, 0664)
		newPasswordFile.WriteString(string(encryptedPassword))
	} else {
		log.Fatal("password with this id already exists")
	}
}

func (passwordsRepository PasswordsRepository) UpdatePassword(id string, password string) {
	passwordFilePath := fmt.Sprintf("%v/%v.txt", passwordsRepository.passwordsPath, id)
	encryptedPassword, encryptPasswordError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if encryptPasswordError != nil {
		log.Fatal(encryptPasswordError)
	}
	_, err := os.Stat(passwordFilePath)
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("password with this id already exists")
	} else {
		passwordFile, _ := os.OpenFile(passwordFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		passwordFile.WriteString(string(encryptedPassword))
	}
}

func (passwordsRepository PasswordsRepository) RemovePassword(id string) {
	passwordFilePath := fmt.Sprintf("%v/%v.txt", passwordsRepository.passwordsPath, id)
	_, err := os.Stat(passwordFilePath)
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("password with this id doesn't exist")
	} else {
		err := os.Remove(passwordFilePath)
		if err != nil {
			log.Fatal("error deleting this password")
		}
	}
}

func NewPasswordsRepository() IPasswordsRepository {

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	passwordsDirectoryPath := filepath.Clean(filepath.Join(workingDir, "/passwords"))

	_, findDirectoryError := os.Stat(passwordsDirectoryPath)

	if errors.Is(findDirectoryError, os.ErrNotExist) {
		os.Mkdir(passwordsDirectoryPath, os.ModePerm)
	}

	return &PasswordsRepository{
		passwordsPath: passwordsDirectoryPath,
	}
}
