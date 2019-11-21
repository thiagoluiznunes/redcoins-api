package user

import (
	"encoding/json"
	"net/http"
	hp "redcoins-api/internal"
	"regexp"
)

var nameRegex = regexp.MustCompile("^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$")
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var passwordRegex = regexp.MustCompile("[a-zA-Z0-9]{6,}")

// GetUser : get user handler
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("res")
	return
}

// SingUp : singup user handler
func SingUp(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	hashPassword, _ := hp.HashPassword(password)

	if !nameRegex.MatchString(name) {
		hp.ResponseHandler(w, r, 406, "Invalid name.")
		return
	}
	if !emailRegex.MatchString(email) {
		hp.ResponseHandler(w, r, 406, "Invalid email")
		return
	}
	if !passwordRegex.MatchString(password) {
		hp.ResponseHandler(w, r, 406, "Invalid password. Must have minimum 6 characters.")
		return
	}

	err := CreateUser(User{
		uuid:     ``,
		name:     name,
		email:    email,
		password: hashPassword})

	if err != nil {
		hp.ResponseHandler(w, r, 406, err.Error())
		return
	}
	hp.ResponseHandler(w, r, 201, "User registered with success.")
	return
}

// Login : login user handler
func Login(w http.ResponseWriter, r *http.Request) {
	user := User{}
	r.ParseForm()
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, _ = FindUserByEmail(email)
	if !hp.CheckPasswordHash(password, user.password) {
		hp.ResponseHandler(w, r, 406, "Invalid email/password.")
		return
	}

	token, err := hp.GenerateToken(user.uuid, user.name, user.email)
	if err != nil {
		hp.ResponseHandler(w, r, 400, err.Error())
		return
	}
	res := hp.JSONJwtResponse{Code: 200, JWT: token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
	return
}
