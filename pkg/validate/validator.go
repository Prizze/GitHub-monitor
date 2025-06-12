package validate

import "github.com/go-playground/validator/v10"

type Validator struct {
	Validator *validator.Validate
}

func NewValidator()*Validator {
	validator := validator.New()
	validator.RegisterValidation("lang", langValidate)

	return &Validator{
		Validator: validator,
	}
}

// Поддерживаемые языки
var allowedLanguages = map[string]struct{} {
	"go": {},
	"python": {},
	"rust": {},
	"c++": {},
	"java": {},
	"javascript": {},
	"swift": {},
}

func langValidate(fl validator.FieldLevel) bool {
	lang := fl.Field().String()
	_, ok := allowedLanguages[lang]
	return ok
}
