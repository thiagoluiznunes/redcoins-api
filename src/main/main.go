package main
import (
	"github.com/gorilla/mux"
  "net/http"
	"encoding/json"
)

type User struct {
  _id string `json:"id"`
  full_name string `json:"title"`
	email string `json:"email"`
	birthday string `json:birthday`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(`user`)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", createUser).Methods("POST")
	http.ListenAndServe(":8000", router)
}
