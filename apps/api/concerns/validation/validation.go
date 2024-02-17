package validation

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/go-playground/validator/v10"
)

var messages = map[string]string{
	"min":      "must be at least %s characters long;",
	"max":      "must be at most %s characters long;",
	"required": "is required;",
	"alphanum": "must be alphanumeric;",
	"boolean":  "must be a boolean;",
}

var paramIgnoreList = []string{"required", "alphanum", "bool"}

func ValidateEntity(entity interface{}) ([]error, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(entity)

	var fullMessage strings.Builder
	var errorList []error
	var fullError error

	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return errorList, err
		}

		for _, err := range err.(validator.ValidationErrors) {

			var message string
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type().Name())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()

			message = fmt.Sprintf(messages[err.Tag()], err.Param())

			if slices.Contains(paramIgnoreList, err.Kind().String()) {
				message = messages[err.Tag()]
			}

			errorMessage := fmt.Sprintf("%s %s", err.Field(), message)

			fullMessage.WriteString(errorMessage)

			fullError = errors.New(fullMessage.String())

			errorList = append(errorList, fullError)
		}

		// from here you can create your own error messages in whatever language you wish
		return errorList, fullError
	}

	return nil, nil
}
