package validation

import "github.com/go-playground/validator/v10"

var Validator = validator.New()

func CollectErrors(err error) []map[string]string {
	errors := []map[string]string{}
	for _, err := range err.(validator.ValidationErrors) {
		error := map[string]string{"field": err.Field(), "tag": err.Tag(), "value": err.Param()}
		errors = append(errors, error)
	}
	
	return errors
}
