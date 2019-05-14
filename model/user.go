package model

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	ID        string
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
}

func (u *User) SaltAndSetPassword() {
	u.Salt = uuid.NewV4().String()
	password := fmt.Sprint(u.Password, u.Salt)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	u.Password = string(hash)
}

func (u *User) CheckPassword(password string) bool {
	password = fmt.Sprint(password, u.Salt)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return u.Password == string(hash)
}
