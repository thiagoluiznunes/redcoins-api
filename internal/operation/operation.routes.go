package operation

import (
	"database/sql"
	"log"

	"github.com/go-chi/chi"
)

// InitOperationRoutes : init all routes from user component
func InitOperationRoutes(db *sql.DB, router *chi.Mux) {
	DB = db
	InitOperationSchema()

	router.Post("/api/v1/operations", Create)

	log.Println("operations: routes registered")
}
