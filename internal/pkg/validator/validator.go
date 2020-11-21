package validator

import (
	validate "github.com/go-playground/validator/v10"
	"sync"
)

var (
	once sync.Once
	validator *validate.Validate
)

func New() *validate.Validate {
	once.Do(func() {
		validator = validate.New()
	})
	return validator
}

func Validate(v interface{}) error {
	return New().Struct(v)
}