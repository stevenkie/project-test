package login

import "github.com/go-playground/validator"

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (p Login) Validate() error {
	return validate.Struct(p)
}
