package database

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(s string) string {
	b := []byte(s)
	hashedString, err := bcrypt.GenerateFromPassword(b, 10)
	if err != nil {
		panic(err)
	}
	return string(hashedString)
}

func EncryptMatch(hashedString, s string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(s))
	if err != nil {
		return false
	}
	return true
}
