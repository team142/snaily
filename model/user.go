package model

import (
	"encoding/base64"
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
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	u.Password = base64.StdEncoding.EncodeToString(hash)
}

func (u *User) CheckPassword(password string) (ok bool) {
	hashB, err := base64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		panic(err)
	}
	hashP := []byte(password)
	err = bcrypt.CompareHashAndPassword(hashB, hashP)
	if err != nil {
		return false
	}
	return true
}
