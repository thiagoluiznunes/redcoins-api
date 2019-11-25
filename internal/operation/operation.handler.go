package operation

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	hp "redcoins-api/internal"

	"github.com/go-chi/chi"
)

// ResponseOperationsHandler : handler
func ResponseOperationsHandler(w http.ResponseWriter, r *http.Request, opt JSONOperationsResponse) {
	var jsonData []byte
	jsonData, err := json.Marshal(opt)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(opt.Code)
	w.Write(jsonData)
}

// Create : get user handler
func Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	operationType := r.Form.Get("operation_type")
	amount, err := strconv.ParseFloat(r.Form.Get("amount"), 64)
	ctx := r.Context()
	signature, ok := ctx.Value("signature").(hp.UserSignature)

	if !ok {
		hp.ResponseHandler(w, r, 403, "Restrict Access")
		return
	}

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
		UserUUID:      signature.UUID})

	if err != nil {
		hp.ResponseHandler(w, r, 406, err.Error())
		return
	}

	hp.ResponseHandler(w, r, 200, "Operation successfully performed.")
	return
}

// GetByUser :  get operationsByUser handler
func GetByUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	signature, ok := ctx.Value("signature").(hp.UserSignature)

	if !ok {
		hp.ResponseHandler(w, r, 403, "Restrict Access")
		return
	}

	operations, err := GetOperationsByID(signature.UUID)
	if err != nil {
		hp.ResponseHandler(w, r, 406, err.Error())
		return
	}

	if len(operations) <= 0 {
		hp.ResponseHandler(w, r, 404, "Operations not found")
		return
	}

	res := JSONOperationsResponse{Code: 200, Operations: operations}
	ResponseOperationsHandler(w, r, res)
}

// GetByDate :  get operationsByDate handler
func GetByDate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	date := chi.URLParam(r, "date")
	signature, ok := ctx.Value("signature").(hp.UserSignature)

	if !ok || signature.Role == "user" {
		hp.ResponseHandler(w, r, 403, "Restrict Access")
		return
	}

	operations, err := GetOperationsByDate(date)
	if err != nil {
		hp.ResponseHandler(w, r, 406, err.Error())
		return
	}

	if len(operations) <= 0 {
		hp.ResponseHandler(w, r, 404, "Operations not found")
		return
	}

	res := JSONOperationsResponse{Code: 200, Operations: operations}
	ResponseOperationsHandler(w, r, res)
}

// GetByEmail :  get operationsByEmail handler
func GetByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := chi.URLParam(r, "email")
	signature, ok := ctx.Value("signature").(hp.UserSignature)

	if !ok || signature.Role == "user" {
		hp.ResponseHandler(w, r, 403, "Restrict Access")
		return
	}

	operations, err := GetOperationsByParam("email", email)
	if err != nil {
		hp.ResponseHandler(w, r, 406, err.Error())
		return
	}

	if len(operations) <= 0 {
		hp.ResponseHandler(w, r, 404, "Operations not found")
		return
	}

	res := JSONOperationsResponse{Code: 200, Operations: operations}
	ResponseOperationsHandler(w, r, res)
	return
}

// GetByName :  get operationsByName handler
func GetByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := chi.URLParam(r, "name")
	signature, ok := ctx.Value("signature").(hp.UserSignature)

	if !ok || signature.Role == "user" {
		hp.ResponseHandler(w, r, 403, "Restrict Access")
		return
	}

	operations, err := GetOperationsByParam("name", name)
	if err != nil {
		hp.ResponseHandler(w, r, 406, err.Error())
		return
	}

	if len(operations) <= 0 {
		hp.ResponseHandler(w, r, 404, "Operations not found")
		return
	}

	res := JSONOperationsResponse{Code: 200, Operations: operations}
	ResponseOperationsHandler(w, r, res)
}
