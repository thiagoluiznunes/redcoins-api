package operation

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	hp "redcoins-api/internal"
)

// Create : get user handler
func Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userUUID := r.Context().Value("uuid")
	operationType := r.Form.Get("operation_type")
	amount, err := strconv.ParseFloat(r.Form.Get("amount"), 64)

	if operationType != "purchase" && operationType != "sale" {
		hp.ResponseHandler(w, r, 406, "Invalid operation")
		return
	}

	if err != nil || amount <= 0 {
		hp.ResponseHandler(w, r, 406, "Invalid amount")
		return
	}

	price, err := hp.RequestBitCoinPrice()
	if err != nil {
		hp.ResponseHandler(w, r, 406, err.Error())
		return
	}

	err = CreateOperation(Operation{
		UUID:          ``,
		OperationType: operationType,
		Amount:        amount,
		Price:         price,
		CreatedAt:     ``,
		UserUUID:      userUUID.(string)})

	if err != nil {
		hp.ResponseHandler(w, r, 406, err.Error())
		return
	}

	hp.ResponseHandler(w, r, 200, "Operation successfully performed.")
	return
}

// Get : get all operations handler
func Get(w http.ResponseWriter, r *http.Request) {
	userUUID := r.Context().Value("uuid")

	operations, err := GetOperations(userUUID.(string))
	if err != nil {
		hp.ResponseHandler(w, r, 406, err.Error())
		return
	}

	res := JSONOperationsResponse{Code: 200, Operations: operations}
	var jsonData []byte
	jsonData, err = json.Marshal(res)

	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonData)
	return
}
