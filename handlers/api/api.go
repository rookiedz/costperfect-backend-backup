package api

import "gopkg.in/go-playground/validator.v9"

var validate *validator.Validate

//Initial ...
func Initial() {
	validate = validator.New()
}
