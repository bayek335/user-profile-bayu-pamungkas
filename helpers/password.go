package helpers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassoword(password string) (string, error) {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 5)

	if err != nil {
		err = errors.New("Internal server error!")
		return password, err
	}

	password = string(hashPass)
	return password, nil
}
