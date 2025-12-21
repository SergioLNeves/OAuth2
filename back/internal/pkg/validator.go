package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     sync.Once
)

// GetValidator returns a singleton instance of the validator.
func GetValidator() *validator.Validate {
	if validate == nil {
		once.Do(func() {
			validate = validator.New()
		})
	}
	return validate
}
