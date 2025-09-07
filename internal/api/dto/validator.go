package dto

import (
	"mime"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate = initValidator()

func initValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("mimetype", func(fl validator.FieldLevel) bool {
		field := fl.Field()

		// If pointer and nil → let "required" handle it
		if field.Kind() == reflect.Pointer {
			if field.IsNil() {
				return true
			}
			field = field.Elem()
		}

		value := field.String()
		if value == "" {
			return false
		}

		_, _, err := mime.ParseMediaType(value)
		return err == nil
	})
	return v
}
