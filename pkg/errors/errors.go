package errors

import (
	validation "blog/internal/provider/validation"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var errorsList = make(map[string]string)

func Init() {
	errorsList = map[string]string{}
}

func SetFromErrors(err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, filedError := range validationErrors {
			Add(filedError.Field(), getErrorMessage(filedError.Tag()))
		}

	}

}

func Add(key string, value string) {

	errorsList[strings.ToLower(key)] = value
}

func getErrorMessage(tag string) string {
	return validation.ErrorMessage()[tag]
}

func Get() map[string]string {
	return errorsList
}
