package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// HashPassword : return hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPasswordHash : check if password math to hash password
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ErrorsHandler : catch panic throwed
func ErrorsHandler(w http.ResponseWriter, r *http.Request) {
	if rec := recover(); rec != nil {
		message := fmt.Sprintf("autorize: %v", rec)
		ResponseHandler(w, r, 406, message)
		return
	}
}

// GenerateToken : create token
func GenerateToken(uuid string, name string, email string, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UUID:  uuid,
		Name:  name,
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return err.Error(), err
	}
	return tokenString, nil
}

// ValidateToken : verify token validate
func ValidateToken(token string) (bool, UserSignature) {
	signature := UserSignature{}
	decode, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	claims := decode.Claims.(jwt.MapClaims)
	if err != nil {
		return false, signature
	}
	if !decode.Valid {
		return false, signature
	}
	signature.UUID = claims["uuid"].(string)
	signature.Role = claims["role"].(string)
	return true, signature
}

// ResponseHandler : handler
func ResponseHandler(w http.ResponseWriter, r *http.Request, code int, message string) {
	res := JSONStandardResponse{Code: code, Message: message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}

// RequestBitCoinPrice : function responsible for request to 3third api
func RequestBitCoinPrice() (float64, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	coinKey := os.Getenv("COIN_MARKET_KEY")

	q := url.Values{}
	q.Add("slug", "bitcoin")
	q.Add("convert", "BRL")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", coinKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		os.Exit(1)
		return 0, err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	bytes := []byte(respBody)
	var res JSONRequestCurrencyQuote
	json.Unmarshal(bytes, &res)

	if err != nil {
		return 0, err
	}
	return res.Data.Num1.Quote.BRL.Price, nil
}
