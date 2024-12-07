package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type PresensiRequest struct {
	Nim string `json:"nim" validate:"required,nim"`
}

// Custom validation function for Nim
func nimValidation(fl validator.FieldLevel) bool {
	nim := fl.Field().String()
	regex := `^\d{2}\.\d{2}\.\d{4}$`
	match, _ := regexp.MatchString(regex, nim)
	return match
}

func PresensiValidator(PresensiRequest PresensiRequest) error {
	validate := validator.New()

	validate.RegisterValidation("nim", nimValidation)

	err := validate.Struct(PresensiRequest)
	if err != nil {
		return err
	}
	return nil
}
