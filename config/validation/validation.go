package validation

import (
	"grouper/adapter/input/model/request"
	"grouper/config/rest_errors"

	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateUserRequest(user *request.UserRequest) *rest_errors.RestErr {
	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, fieldErr := range validationErrors {
			errorMessages = append(errorMessages, fieldErr.Error())
		}
		return rest_errors.NewBadRequestError("Invalid request parameters: " + strings.Join(errorMessages, ", "))
	}
	return nil
}
