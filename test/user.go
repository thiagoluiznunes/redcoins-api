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

// UserSignUp : describe
func UserSignUp(t *testing.T) {
	secret := "Red_C0ins123"
	user, err := json.Marshal(map[string]string{
		"name":             "Admin test",
		"email":            "admin.test@email.com",
		"password":         "adminusertest123",
		"confirm_password": "adminusertest123",
		"secret":           secret,
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

	if res.Code != 201 {
		t.Error(fmt.Sprintf(`Failed to create an user. Message: %s. Code: %d`, res.Message, res.Code))
		return
	}
	t.Log(res.Message)
}

// UserLogin : describe
func UserLogin(t *testing.T) {
	user, err := json.Marshal(map[string]string{
		"email":    "admin.test@email.com",
		"password": "adminusertest123",
	})

	if err != nil {
		t.Fatalf("Failed to parse user in JSON object")
		return
	}

	resp, err := http.Post("http://localhost:8000/api/v1/users/login", "application/json", bytes.NewBuffer(user))
	if err != nil {
		t.Fatalf("Failed to request login")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to retrieve response body from login request")
		return
	}
	bytes := []byte(body)
	var res hp.JSONStandardResponse
	json.Unmarshal(bytes, &res)

	if res.Code != 200 {
		t.Fatalf(fmt.Sprintf(`Failed to login. Message: %s. Code: %d`, res.Message, res.Code))
		return
	}
	var jwtRes hp.JSONJwtResponse
	json.Unmarshal(bytes, &jwtRes)
	UserJWT = jwtRes.JWT
}

// DeleteUser : describe
func DeleteUser(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8000/api/v1/users/test", nil)
	if err != nil {
		t.Error("Failed to create user delete request")
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Failed to request user delete")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Failed to retrieve response body from user delete request")
		return
	}
	bytes := []byte(body)
	var res hp.JSONStandardResponse
	json.Unmarshal(bytes, &res)

	if res.Code != 200 {
		t.Error(fmt.Sprintf(`Failed to delete user. Message: %s. Code: %d`, res.Message, res.Code))
		return
	}
}
