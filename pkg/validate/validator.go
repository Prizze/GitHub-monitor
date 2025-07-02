package validate

import (
	"log"

	"github.com/Prizze/GitHub-monitor/gh-monitor/domain"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *validator.Validate
}

func NewValidator() *Validator {
	validator := validator.New()
	if err := validator.RegisterValidation("lang", langValidate); err != nil {
		log.Println("add validation field failed: ", err.Error())
	}

	return &Validator{
		Validator: validator,
	}
}

func langValidate(fl validator.FieldLevel) bool {
	lang := fl.Field().String()
	_, ok := domain.AllowedLanguages[lang]
	return ok
}
