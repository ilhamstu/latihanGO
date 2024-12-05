package validators

import (
	"github.com/go-playground/validator/v10"
)

// Struct untuk pesan error
type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Fungsi untuk mendapatkan pesan kesalahan yang lebih deskriptif
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Please provide a valid email address"
	case "min":
		return "Password must be at least 8 characters long"
	}
	return "Unknown error"
}

// Struct untuk User Register Request
type UserRegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
	Nik      string `json:"nik" validate:"required"`
}

// Struct untuk Mahasiswa Register Request
type MahasiswaRegisterRequest struct {
	Nim      string `json:"nim" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// Struct untuk Login Request
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// Fungsi untuk memvalidasi User Register Request
func ValidateUserRegisterRequest(userRegisterRequest UserRegisterRequest) ([]ErrorMsg, error) {
	validate := validator.New()
	err := validate.Struct(userRegisterRequest)
	if err != nil {
		var errorMsgs []ErrorMsg
		for _, err := range err.(validator.ValidationErrors) {
			errorMsgs = append(errorMsgs, ErrorMsg{
				Field:   err.Field(),
				Message: GetErrorMsg(err),
			})
		}
		return errorMsgs, err
	}
	return nil, nil
}

// Fungsi untuk memvalidasi Mahasiswa Register Request
func ValidateMahasiswaRegisterRequest(mahasiswaRegisterRequest MahasiswaRegisterRequest) ([]ErrorMsg, error) {
	validate := validator.New()
	err := validate.Struct(mahasiswaRegisterRequest)
	if err != nil {
		var errorMsgs []ErrorMsg
		for _, err := range err.(validator.ValidationErrors) {
			errorMsgs = append(errorMsgs, ErrorMsg{
				Field:   err.Field(),
				Message: GetErrorMsg(err),
			})
		}
		return errorMsgs, err
	}
	return nil, nil
}

// Fungsi untuk memvalidasi Login Request
func ValidateLoginRequest(loginRequest LoginRequest) ([]ErrorMsg, error) {
	validate := validator.New()
	err := validate.Struct(loginRequest)
	if err != nil {
		var errorMsgs []ErrorMsg
		for _, err := range err.(validator.ValidationErrors) {
			errorMsgs = append(errorMsgs, ErrorMsg{
				Field:   err.Field(),
				Message: GetErrorMsg(err),
			})
		}
		return errorMsgs, err
	}
	return nil, nil
}
