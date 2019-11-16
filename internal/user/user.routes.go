package user

import (
	"database/sql"
	"fmt"

	"github.com/go-chi/chi"
)

// DB : database instance
var DB *sql.DB

// InitUserRoutes : init all routes from user component
func InitUserRoutes(db *sql.DB, router *chi.Mux) {
	DB = db
	fmt.Println("Init Users routes.")

	router.Get("/user", GetUser)
	router.Post("/user", CreateUser)
}
