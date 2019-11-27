package user

import (
	"encoding/json"
	"net/http"
	"os"
	hp "redcoins-api/internal"
	"regexp"
)

var nameRegex = regexp.MustCompile("^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$")
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var passwordRegex = regexp.MustCompile("[a-zA-Z0-9]{6,}")

// SingUp : singup user handler
func SingUp(w http.ResponseWriter, r *http.Request) {
	var body BodyRequest
	json.NewDecoder(r.Body).Decode(&body)

	if !nameRegex.MatchString(body.Name) {
		hp.ResponseHandler(w, r, 406, "Invalid name.")
		return
	}
	if !emailRegex.MatchString(body.Email) {
		hp.ResponseHandler(w, r, 406, "Invalid email.")
		return
	}
	if !passwordRegex.MatchString(body.Password) {
		hp.ResponseHandler(w, r, 406, "Invalid password. Must have minimum 6 characters.")
		return
	}
	if body.ConfirmPassword != body.Password {
		hp.ResponseHandler(w, r, 406, "Passwords diverges.")
		return
	}

	hashPassword, err := hp.HashPassword(body.Password)
	role := "user"
	if body.Secret == os.Getenv("JWT_SECRET") {
		role = "admin"
	}
	err = CreateUser(User{
		UUID:     ``,
		Name:     body.Name,
		Email:    body.Email,
		Password: hashPassword,
		Role:     role})

	if err != nil {
		hp.ResponseHandler(w, r, 406, err.Error())
		return
	}
	hp.ResponseHandler(w, r, 201, "User registered with success.")
	return
}

// Login : login user handler
func Login(w http.ResponseWriter, r *http.Request) {
	var body BodyRequest
	json.NewDecoder(r.Body).Decode(&body)

	user, err := FindUserByEmail(body.Email)
	if err != nil {
		hp.ResponseHandler(w, r, 406, "Invalid email/password.")
		return
	}

	if !hp.CheckPasswordHash(body.Password, user.Password) {
		hp.ResponseHandler(w, r, 406, "Invalid email/password.")
		return
	}

	token, err := hp.GenerateToken(user.UUID, user.Name, user.Email, user.Role)
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
