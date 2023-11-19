package helpers

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

func DatabaseErrorHandler(err error, table string) (error, int) {

	var code int

	if table == "user" {
		if strings.Contains(err.Error(), "unique") {
			err = errors.New("Email already used!")
			code = 409
		} else {
			err = errors.New("Internal service error!")
			code = 500
		}
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(table + " does not exist!")
			code = 404
		} else {
			err = errors.New("Internal service error!")
			code = 500
		}
	}
	return err, code

}
