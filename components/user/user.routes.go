package user

import (
	"fmt"

	"github.com/gorilla/mux"
)

// InitUserRoutes : init all routes from user component
func InitUserRoutes(router *mux.Router) {
	fmt.Println("Init Users routes.")

	router.HandleFunc("/user", GetUser).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")

}
