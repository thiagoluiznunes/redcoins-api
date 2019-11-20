package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))
var nameRegex = regexp.MustCompile("^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$")
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var passwordRegex = regexp.MustCompile("[a-zA-Z0-9]{6,}")

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

// Claims : declares structure
type Claims struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func init() {
	fmt.Println("User: Init Handler.")
}

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

	err := CreateUser(User{
		uuid:     ``,
		name:     name,
		email:    email,
		password: hashPassword})

	if err != nil {
		res := JSONStandardResponse{406, err.Error()}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}
	res := JSONStandardResponse{201, "User registered with success."}
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

	if !CheckPasswordHash(password, user.password) {
		res := JSONStandardResponse{406, "Invalid email/password."}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UUID:  user.uuid,
		Name:  user.name,
		Email: user.email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		res := JSONStandardResponse{500, "Internal error."}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}
	res := JSONJwtResponse{200, tokenString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	return
}
