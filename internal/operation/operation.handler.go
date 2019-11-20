package operation

import (
	"encoding/json"
	"net/http"
)

// Create : get user handler
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("res")
	return
}
