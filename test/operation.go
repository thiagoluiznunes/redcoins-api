package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

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
	t.Log("operation: " + res.Message)
}

// GetOperationByUser : describe
func GetOperationByUser(t *testing.T) {
	var bearer = "Bearer " + UserJWT
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8000/api/v1/operations", nil)
	req.Header.Add("Authorization", bearer)
	if err != nil {
		t.Error("operation: failed to create operation get request")
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("operation: failed to request operation get")
		return
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error("operation: failed to retrieve response body from operation get request")
		return
	}
	if resp.StatusCode == 404 {
		t.Log("operation: operations not found")
		return
	} else if resp.StatusCode != 200 {
		t.Error("operation: failed to get operations")
	}
	t.Log("operation: success in getting operations")
}

// GetOperationByDate : describe
func GetOperationByDate(t *testing.T) {
	var bearer = "Bearer " + UserJWT
	client := &http.Client{}
	time := time.Now()
	date := time.Format("2006-01-02")

	url := fmt.Sprintf(`http://localhost:8000/api/v1/operations/date/%s`, date)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	if err != nil {
		t.Error("operation: failed to create operation get request")
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("operation: failed to request operation get")
		return
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error("operation: failed to retrieve response body from operation get by date request")
		return
	}
	if resp.StatusCode == 404 {
		t.Log("operation: operations not found")
		return
	} else if resp.StatusCode == 403 {
		t.Error("operation: restrict access")
		return
	} else if resp.StatusCode != 200 {
		t.Error("operation: failed to get operations by date")
		return
	}
	t.Log("operation: success in getting operations by date")
}

// GetOperationByUserEmail : describe
func GetOperationByUserEmail(t *testing.T) {
	var bearer = "Bearer " + UserJWT
	client := &http.Client{}
	email := "admin.test@email.com"

	url := fmt.Sprintf(`http://localhost:8000/api/v1/operations/email/%s`, email)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	if err != nil {
		t.Error("operation: failed to create operation get request")
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("operation: failed to request operation get")
		return
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error("operation: failed to retrieve response body from operation get by email request")
		return
	}
	if resp.StatusCode == 404 {
		t.Log("operation: operations not found")
		return
	} else if resp.StatusCode == 403 {
		t.Error("operation: restrict access")
		return
	} else if resp.StatusCode != 200 {
		t.Error("operation: failed to get operations by email")
		return
	}
	t.Log("operation: success in getting operations by email")
}

// DeleteOperation : describe
func DeleteOperation(t *testing.T) {
	var bearer = "Bearer " + UserJWT
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "http://localhost:8000/api/v1/operations/test", nil)
	req.Header.Add("Authorization", bearer)
	if err != nil {
		t.Error("operation: failed to create operation delete request")
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("operation: failed to request operation delete")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("operation: failed to retrieve response body from operation delete request")
		return
	}
	bytes := []byte(body)
	var res hp.JSONStandardResponse
	json.Unmarshal(bytes, &res)

	if res.Code != 200 {
		t.Error(fmt.Sprintf(`Failed to delete operation. Message: %s. Code: %d`, res.Message, res.Code))
		return
	}
	t.Log(res.Message)
}
