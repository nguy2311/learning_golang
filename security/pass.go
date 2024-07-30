package security

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) string {
	bytes, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}
	return string(bytes)
}

func CheckPasswordHash(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
