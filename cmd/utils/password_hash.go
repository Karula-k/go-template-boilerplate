package utils

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedBytes), err

}
