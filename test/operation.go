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

// UserJWT : describe
var UserJWT string

// CreateOperation : test case
func CreateOperation(t *testing.T) {
	var bearer = "Bearer " + UserJWT
	client := &http.Client{}
	opt, err := json.Marshal(map[string]string{
		"operation_type": "purchase",
		"amount":         "150",
	})

	if err != nil {
		t.Error("operation: failed to parse operation in JSON object")
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8000/api/v1/operations", bytes.NewBuffer(opt))
	req.Header.Add("Authorization", bearer)
	if err != nil {
		t.Error("operation: failed to create delete request")
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error("operation: failed to request delete")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("operation: retrieve response body from create operation")
		return
	}

	bytes := []byte(body)
	var res hp.JSONStandardResponse
	json.Unmarshal(bytes, &res)

	if res.Code != 200 {
		msg := fmt.Sprintf(`operation: to create an operation. Response code: %d`, res.Code)
		t.Error(msg)
		return
	}
	t.Log(res.Message)
}

// DeleteOperation : describe
func DeleteOperation(t *testing.T) {
	var bearer = "Bearer " + UserJWT
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8000/api/v1/operations/test", nil)
	req.Header.Add("Authorization", bearer)
	if err != nil {
		t.Error("Failed to create operation delete request")
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Failed to request operation delete")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Failed to retrieve response body from operation delete request")
		return
	}
	bytes := []byte(body)
	var res hp.JSONStandardResponse
	json.Unmarshal(bytes, &res)

	if res.Code != 200 {
		t.Error(fmt.Sprintf(`Failed to delete operation. Message: %s. Code: %d`, res.Message, res.Code))
		return
	}
}
