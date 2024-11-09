package service

import (
	"regexp"

	"github.com/go-playground/validator"
	"github.com/paemuri/brdoc"
)

var (
	Validate *validator.Validate
)

func init() {
	Validate = validator.New()
	Validate.RegisterValidation("cpf", validateCPF)
}

func validateCPF(fl validator.FieldLevel) bool {
	cpf := fl.Field().String()
	re := regexp.MustCompile(`\D`)
	cpf = re.ReplaceAllString(cpf, "")

	if len(cpf) != 11 {
		return false
	}

	return brdoc.IsCPF(cpf)
}
