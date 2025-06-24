package validate

import (
	"github.com/Prizze/GitHub-monitor/gh-monitor/domain"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *validator.Validate
}

func NewValidator() *Validator {
	validator := validator.New()
	validator.RegisterValidation("lang", langValidate)

	return &Validator{
		Validator: validator,
	}
}

func langValidate(fl validator.FieldLevel) bool {
	lang := fl.Field().String()
	_, ok := domain.AllowedLanguages[lang]
	return ok
}
