package internal

import (
	"fmt"
	"net/http"
	"strings"

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

func Autorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Get("Authorization"), "Bearer")[1]
		fmt.Println(token)
		next.ServeHTTP(w, r)
	})
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
