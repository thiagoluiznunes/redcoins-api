package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("Init controller.")
}

// GetUser : describe what this function does
func GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(`GET USER`)
}

// CreateUser : describe what this function does
func CreateUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(`POST USER`)
}
