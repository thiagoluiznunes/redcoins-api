package user

import (
	"fmt"

	"github.com/go-chi/chi"
)

// InitUserRoutes : init all routes from user component
// func InitUserRoutes(router *mux.Router) {
func InitUserRoutes(router *chi.Mux) {
	fmt.Println("Init Users routes.")

	router.Get("/user", GetUser)
	router.Post("/user", CreateUser)
	// router.HandleFunc("/user", GetUser).Methods("GET")
	// router.HandleFunc("/user", CreateUser).Methods("POST")
}
