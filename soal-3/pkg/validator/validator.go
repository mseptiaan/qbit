package validator

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

type ErrorResponseValidator struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()
	validate.RegisterValidation("is-password", ValidatePassword)
	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) []*ErrorResponseValidator {
	var errors []*ErrorResponseValidator
	for _, err := range err.(validator.ValidationErrors) {
		var element ErrorResponseValidator
		element.Param = err.Field()
		element.Message = msgForTag(err)
		errors = append(errors, &element)
	}
	return errors
}

func ValidatePassword(fl validator.FieldLevel) bool {
	return validPassword(fl.Field().String())
}

func validPassword(s string) bool {
next:
	for _, classes := range map[string][]*unicode.RangeTable{
		"upper case": {unicode.Upper, unicode.Title},
		"lower case": {unicode.Lower},
		"numeric":    {unicode.Number, unicode.Digit},
		"special":    {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
	} {
		for _, r := range s {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return false
	}
	return true
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "This field does not fulfill the minimum character"
	case "max":
		return "This field exceeds the maximum character"
	case "is-password":
		return "This field does not containing at least 1 capital characters, 1 number or 1 special characters."
	}
	return fe.Error() // default error
}
