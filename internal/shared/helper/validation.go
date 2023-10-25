package helper

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	NameSpace       string
	Field           string
	StructNamespace string
	StructField     string
	Tag             string
	ActualTag       string
	Kind            reflect.Kind
	Type            reflect.Type
	Value           any
	Param           string
}

var validate = validator.New()

func Validate(input interface{}) (validationErrors []*ValidationError) {
	if err := validate.Struct(input); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validate := ValidationError{
				NameSpace:       err.Namespace(),
				Field:           err.Field(),
				StructNamespace: err.StructNamespace(),
				StructField:     err.StructField(),
				Tag:             err.Tag(),
				ActualTag:       err.ActualTag(),
				Kind:            err.Kind(),
				Type:            err.Type(),
				Value:           err.Value(),
				Param:           err.Param(),
			}
			validationErrors = append(validationErrors, &validate)
		}
	}
	return validationErrors
}

func ValidateAlphanumericPunct(input string) bool {
	// Define a regular expression pattern for allowed characters
	pattern := "^[a-zA-Z0-9 !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~]*$"
	return regexp.MustCompile(pattern).MatchString(input)
}

func ValidateDateFormat(input string) bool {
	datePattern := `^\d{2}/\d{2}/\d{4}$`
	return regexp.MustCompile(datePattern).MatchString(input)
}
