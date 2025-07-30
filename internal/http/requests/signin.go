package requests

import (
	"errors"

	validator "github.com/go-playground/validator/v10"
)

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

var validate = validator.New()

// Validate checks the SignInRequest for required fields and valid email format
func (r *SignInRequest) Validate() error {
	err := validate.Struct(r)
	if err != nil {
		return errors.New("invalid request: " + err.Error())
	}
	return nil
}
