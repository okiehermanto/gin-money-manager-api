package helper

import "github.com/go-playground/validator/v10"

func GetValidationError(err error) string {

	validationErrors, ok := err.(validator.ValidationErrors)

	if !ok {
		return err.Error()
	}

	for _, e := range validationErrors {

		switch e.Tag() {

		case "required":
			return e.Field() + " is required"

		case "min":
			return e.Field() + " is too short"

		case "max":
			return e.Field() + " is too long"

		default:
			return e.Field() + " is invalid"
		}
	}

	return err.Error()
}
