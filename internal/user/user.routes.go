package user

import (
	"database/sql"
	"log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// InitUserRoutes : init all routes from user component
func InitUserRoutes(db *sql.DB, router *chi.Mux) {
	DB = db
	InitUserSchema()

	router.Use(middleware.Logger)
	router.Get("/api/v1/user", GetUser)
	router.Post("/api/v1/signup", SingUp)
	router.Post("/api/v1/login", Login)

	log.Println("users: routes registered")
}
