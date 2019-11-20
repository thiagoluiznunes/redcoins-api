package internal

import (
	"golang.org/x/crypto/bcrypt"
)

// JSONStandardResponse : structure to classify JSON response
type JSONStandardResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// JSONJwtResponse : structure to classify JSON JWT response
type JSONJwtResponse struct {
	Code int    `json:"code"`
	JWT  string `json:"jwt"`
}

// HashPassword : return hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash : check if password math to hash password
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
