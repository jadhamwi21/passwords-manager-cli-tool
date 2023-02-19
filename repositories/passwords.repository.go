package repositories

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jadhamwi21/passwords-manager-cli-tool/services"
)

type i_passwords_repository interface {
	AddPassword(id string, password string)
	UpdatePassword(id string, password string)
	RemovePassword(id string)
	RetrieveAllPasswords()
	RetrievePasswordById(id string)
}

type passwords_repository struct {
	passwordsPath string
}

func (passwordsRepository passwords_repository) AddPassword(id string, password string) {
	passwordFilePath := fmt.Sprintf("%v/%v.txt", passwordsRepository.passwordsPath, id)
	_, err := os.Stat(passwordFilePath)
	if errors.Is(err, os.ErrNotExist) {
		newPasswordFile, _ := os.OpenFile(passwordFilePath, os.O_CREATE, 0664)
		newPasswordFile.WriteString(services.EncryptionService.Encrypt(password))
	} else {
		log.Fatal("password with this id already exists")
	}
}

func (passwordsRepository passwords_repository) UpdatePassword(id string, password string) {
	passwordFilePath := fmt.Sprintf("%v/%v.txt", passwordsRepository.passwordsPath, id)
	_, err := os.Stat(passwordFilePath)
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("password with this id already exists")
	} else {
		passwordFile, _ := os.OpenFile(passwordFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		passwordFile.WriteString(services.EncryptionService.Encrypt(password))
	}
}

func (passwordsRepository passwords_repository) RemovePassword(id string) {
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

func (passwordsRepository passwords_repository) RetrieveAllPasswords() {
	files, err := ioutil.ReadDir(passwordsRepository.passwordsPath)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		log.Fatal("passwords directory is empty")
	}
	passwordsMap := make(map[string]string)
	for _, file := range files {
		passwordFileNameSplitted := strings.Split(file.Name(), ".")
		PASSWORD_ID := passwordFileNameSplitted[0]
		PASSWORD_VALUE, err := os.ReadFile(filepath.Clean(filepath.Join(passwordsRepository.passwordsPath, "/", file.Name())))
		if err != nil {
			log.Fatal(err)
		}
		passwordsMap[PASSWORD_ID] = string(services.EncryptionService.Decrypt(string(PASSWORD_VALUE)))
	}
	for password_id, password_value := range passwordsMap {
		fmt.Printf("%v : %v\n", password_id, password_value)
	}
}

func (passwordsRepository passwords_repository) RetrievePasswordById(id string) {

	PASSWORD_VALUE, err := os.ReadFile(filepath.Clean(filepath.Join(passwordsRepository.passwordsPath, fmt.Sprintf("%v.txt", id))))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(services.EncryptionService.Decrypt(string(PASSWORD_VALUE))))
}

func newPasswordsRepository() i_passwords_repository {

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	passwordsDirectoryPath := filepath.Clean(filepath.Join(workingDir, "/passwords"))

	_, findDirectoryError := os.Stat(passwordsDirectoryPath)

	if errors.Is(findDirectoryError, os.ErrNotExist) {
		os.Mkdir(passwordsDirectoryPath, os.ModePerm)
	}

	return &passwords_repository{
		passwordsPath: passwordsDirectoryPath,
	}
}

var PasswordsRepository i_passwords_repository = newPasswordsRepository()
