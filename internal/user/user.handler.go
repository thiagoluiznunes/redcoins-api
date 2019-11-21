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
		res := hp.JSONStandardResponse{Code: 406, Message: "Invalid name."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res)
		return
	}
	if !emailRegex.MatchString(email) {
		res := hp.JSONStandardResponse{Code: 406, Message: "Invalid email."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res)
		return
	}
	if !passwordRegex.MatchString(password) {
		res := hp.JSONStandardResponse{Code: 406, Message: "Invalid password. Must have minimum 6 characters."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res)
		return
	}

	err := CreateUser(User{
		uuid:     ``,
		name:     name,
		email:    email,
		password: hashPassword})

	if err != nil {
		res := hp.JSONStandardResponse{Code: 406, Message: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res)
		return
	}
	res := hp.JSONStandardResponse{Code: 201, Message: "User registered with success."}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
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
		res := hp.JSONStandardResponse{Code: 406, Message: "Invalid email/password."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res)
		return
	}

	token, err := hp.GenerateToken(user.uuid, user.name, user.email)
	if err != nil {
		res := hp.JSONStandardResponse{Code: 500, Message: "Internal error."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res)
		return
	}
	res := hp.JSONJwtResponse{Code: 200, JWT: token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
	return
}
