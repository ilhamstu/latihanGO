package validators

import (
	"github.com/go-playground/validator/v10"
)

// Struct untuk validasi input saat membuat post
type PostRequest struct {
	Title   string `json:"title" validate:"required,min=5,max=100"`
	Content string `json:"content" validate:"required,min=10"`
}

// Fungsi untuk melakukan validasi terhadap input PostRequest
func ValidatePostRequest(postRequest PostRequest) error {
	// Membuat instance validator
	validate := validator.New()

	// Melakukan validasi terhadap postRequest
	err := validate.Struct(postRequest)
	if err != nil {
		return err
	}
	return nil
}
