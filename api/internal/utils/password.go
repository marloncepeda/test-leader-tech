package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"crypto/subtle"
	"fmt"

	"github.com/xdg-go/pbkdf2"
)

func HashPassword(password string) ([]byte, []byte) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		fmt.Println("Error generating salt", err)
	}

	hashedPassword := pbkdf2.Key([]byte(password), salt, 4096, 32, sha512.New)

	return hashedPassword, salt
}

func CheckPassword(password string, hashed []byte, salt []byte) bool {
	newHashed := pbkdf2.Key([]byte(password), salt, 4096, 32, sha512.New)

	return subtle.ConstantTimeCompare(hashed, newHashed) == 1
}
