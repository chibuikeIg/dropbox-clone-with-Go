package rv

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Err validator.ValidationErrors
}

func (v Validator) Errors() map[string]string {

	validationErrors := make(map[string]string)

	for _, err := range v.Err {

		validationErrors[err.Field()] = "The " + err.Field() + " field validation failed, reason: " + err.ActualTag()
	}

	return validationErrors
}
