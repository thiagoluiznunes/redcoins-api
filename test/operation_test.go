package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	hp "redcoins-api/internal"
	pkg "redcoins-api/internal/operation"
)

// TestCreateOperation : test case
func TestCreateOperation(t *testing.T) {
	opt := pkg.Operation{
		UUID:          ``,
		OperationType: `sale`,
		Amount:        150,
		Price:         3000.000,
		CreatedAt:     ``,
		UserUUID:      `asdf-asddsa-dsada-dasdasd`}

	requestBody, err := json.Marshal(opt)
	if err != nil {
		t.Error("Failed to parse operation in JSON object")
		return
	}

	resp, err := http.Post("http://localhost:8000/api/v1/operations", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Error("Failed to create an operation")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Failed to retrieve response body from create operation")
		return
	}
	bytes := []byte(body)
	var res hp.JSONStandardResponse
	json.Unmarshal(bytes, &res)

	if res.Code != 200 {
		t.Error(`Failed to create an operation. Response code:`, res.Code)
		return
	}
	t.Log("Operation successfully created")
}
