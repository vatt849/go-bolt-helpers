package helpers

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var ErrNotImplemented = errors.New(ERR_MSG_NOT_IMPLEMENTED)

func IsErrNotImplemented(err error) bool {
	return errors.Is(err, ErrNotImplemented)
}

func IsErrDbNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func IsErrValidation(err error) bool {
	_, check := err.(validator.ValidationErrors)

	return check
}

func ErrToValidation(err error) validator.ValidationErrors {
	return err.(validator.ValidationErrors)
}
