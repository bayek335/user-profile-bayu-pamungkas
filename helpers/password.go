package helpers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Hash(str string) (string, error) {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(str), 10)

	if err != nil {
		err = errors.New("Internal server error!")
		return str, err
	}

	str = string(hashPass)
	return str, nil
}
