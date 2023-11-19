package helpers

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func UserValidation(err error) (error, int) {
	// Return the first error message that has founded
	for _, er := range err.(validator.ValidationErrors) {
		if er.ActualTag() == "required" {
			err = errors.New("Field " + er.Field() + " is " + er.ActualTag() + "!")
			return err, 400
		} else if er.ActualTag() == "email" {
			err = errors.New("Field " + er.Field() + " is a type of email!")
			return err, 400
		} else if er.ActualTag() == "min" {
			err = errors.New("Field " + er.Field() + " must have at least 6 characters!")
			return err, 400
		} else {
			err = errors.New("Internal server error!")
		}
	}
	return err, 500
}
