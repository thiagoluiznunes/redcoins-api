package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

var nameRegex = regexp.MustCompile("^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$")
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var passwordRegex = regexp.MustCompile("[a-zA-Z0-9]{6,}")

// JSONResponseError : structure to classify JSON response error
type JSONResponseError struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

func init() {
	fmt.Println("User: Init Handler.")
}

// GetUser : describe what this function does
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("res")
	return
}

// SingUp : describe what this function does
func SingUp(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	hashPassword, _ := HashPassword(password)

	if !nameRegex.MatchString(name) {
		json.NewEncoder(w).Encode("Invalid name!")
		return
	}
	if !emailRegex.MatchString(email) {
		json.NewEncoder(w).Encode("Invalid email!")
		return
	}
	if !passwordRegex.MatchString(password) {
		json.NewEncoder(w).Encode("Invalid password. Must have minimum 6 characters.")
		return
	}

	err := CreateUserService(User{
		uuid:     ``,
		name:     name,
		email:    email,
		password: hashPassword})

	if err != nil {
		res := JSONResponseError{406, err.Error()}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}
	res := JSONResponseError{201, "User registered with success."}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	return
}

// Login : describe function
func Login(w http.ResponseWriter, r *http.Request) {
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"uuid":     user.Username,
	// 	"password": user.Password,
	// })
	// os.Getenv("JWT_SECRET")
}
