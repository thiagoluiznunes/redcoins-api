package user

import (
	"database/sql"
	"log"

	"github.com/go-chi/chi"
)

// Routes : init all routes from user component
func Routes(db *sql.DB) chi.Router {
	DB = db
	InitUserSchema()

	router := chi.NewRouter()
	router.Post("/signup", SingUp)
	router.Post("/login", Login)

	log.Println("users: routes registered")

	return router
}
