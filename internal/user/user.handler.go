package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var passwordRegex = regexp.MustCompile("[a-zA-Z0-9]{6,}")

func init() {
	fmt.Println("Init controller.")
}

// GetUser : describe what this function does
func GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(`GET USER`)
}

// CreateUser : describe what this function does
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	hashPassword, _ := HashPassword(password)

	if !emailRegex.MatchString(email) {
		json.NewEncoder(w).Encode("Invalid email!")
		return
	}
	if !passwordRegex.MatchString(password) {
		json.NewEncoder(w).Encode("Invalid password. Must have minimum 6 characters.")
		return
	}

	user := User{name: name, email: email, password: hashPassword}
	fmt.Println(user)

	newUser, err := CreateUserService(user)
	if err == nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(newUser)
	return
}
