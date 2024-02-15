package validation

import "github.com/go-playground/validator/v10"

func Validate() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate
}
