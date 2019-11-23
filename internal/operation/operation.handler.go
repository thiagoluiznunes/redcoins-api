package operation

import (
	"encoding/json"
	"net/http"
	"strconv"

	hp "redcoins-api/internal"
)

// Create : get user handler
func Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uuid := r.Context().Value("uuid")
	operationType := r.Form.Get("operation_type")
	amount, err := strconv.ParseFloat(r.Form.Get("amount"), 64)

	if operationType != "purchase" && operationType != "sale" {
		hp.ResponseHandler(w, r, 406, "Invalid operation")
		return
	}

	if err != nil {
		hp.ResponseHandler(w, r, 406, "Invalid amount")
		return
	}

	price, err := hp.RequestBitCoinPrice()
	if err != nil {
		hp.ResponseHandler(w, r, 406, err.Error())
		return
	}

	err = CreateOperation(Operation{
		uuid:          ``,
		opertaionType: operationType,
		amount:        amount,
		price:         price,
		userUUID:      uuid.(string)})

	if err != nil {
		res := hp.JSONStandardResponse{Code: 406, Message: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("res")
	return
}
