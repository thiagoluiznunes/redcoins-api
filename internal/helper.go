package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// Claims : declares structure
type Claims struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

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

// ErrorsHandler : catch panic throwed
func ErrorsHandler(w http.ResponseWriter, r *http.Request) {
	if r := recover(); r != nil {
		message := fmt.Sprintf("autorize: %v", r)
		res := JSONStandardResponse{Code: 406, Message: message}
		fmt.Println(r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res)
		return
	}
}

// AutorizeMiddleware : middleware to filter unauthorized requests
func AutorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer ErrorsHandler(w, r)

		header := r.Header.Get("Authorization")
		if header == "" {
			res := JSONStandardResponse{Code: 401, Message: "Unauthorized access."}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(res.Code)
			json.NewEncoder(w).Encode(res)
			return
		}

		token := strings.Split(header, "Bearer")[1]
		isValid := ValidateToken(strings.TrimSpace(token))
		if isValid {
			next.ServeHTTP(w, r)
			return
		}
		res := JSONStandardResponse{Code: 401, Message: "Unauthorized access."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res)
		return
	})
}

// GenerateToken : create token
func GenerateToken(uuid string, name string, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UUID:  uuid,
		Name:  name,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken : verify token validate
func ValidateToken(token string) bool {
	decode, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if !decode.Valid {
		return false
	}
	return true
}
