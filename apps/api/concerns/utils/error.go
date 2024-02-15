package utils

import "github.com/go-playground/validator/v10"

func ValidationErrorToStrList(err error) []string {
	validationErrors := err.(validator.ValidationErrors)

	var errors []string

	for _, validationError := range validationErrors {
		errors = append(errors, validationError.Error())
	}

	return errors
}
