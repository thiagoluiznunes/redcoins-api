package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	hp "redcoins-api/internal"
)

var userJwt string

// TestUserSignUp : test case
func TestUserSignUp(t *testing.T) {
	user, err := json.Marshal(map[string]string{
		"name":             "User test",
		"email":            "user.test@email.com",
		"password":         "usertest123",
		"confirm_password": "usertest123",
	})

	if err != nil {
		t.Error("Failed to parse user in JSON object")
		return
	}

	resp, err := http.Post("http://localhost:8000/api/v1/users/signup", "application/json", bytes.NewBuffer(user))
	if err != nil {
		t.Error("Failed to request signup")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Failed to retrieve response body from signup request")
		return
	}
	bytes := []byte(body)
	var res hp.JSONStandardResponse
	json.Unmarshal(bytes, &res)

	if res.Code != 200 {
		t.Error(fmt.Sprintf(`Failed to create an user. Message: %s. Code: %d`, res.Message, res.Code))
		return
	}
	t.Log("User successfully created")
}

// TestUserLogin : test case
func TestUserLogin(t *testing.T) {
	user, err := json.Marshal(map[string]string{
		"email":    "user.test@email.com",
		"password": "usertest123",
	})

	if err != nil {
		t.Error("Failed to parse user in JSON object")
		return
	}

	resp, err := http.Post("http://localhost:8000/api/v1/users/login", "application/json", bytes.NewBuffer(user))
	if err != nil {
		t.Error("Failed to request login")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Failed to retrieve response body from login request")
		return
	}
	bytes := []byte(body)
	var res hp.JSONStandardResponse
	json.Unmarshal(bytes, &res)

	if res.Code != 200 {
		t.Error(fmt.Sprintf(`Failed to login. Message: %s. Code: %d`, res.Message, res.Code))
		return
	}
	var data hp.JSONJwtResponse
	json.Unmarshal(bytes, &data)
	userJwt = data.JWT
	t.Log("User successfully logged")
}
