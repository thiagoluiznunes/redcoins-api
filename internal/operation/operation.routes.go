package operation

import (
	"database/sql"
	"log"

	"github.com/go-chi/chi"
)

// Routes : init all routes from user component
func Routes(db *sql.DB) chi.Router {
	DB = db
	InitOperationSchema()

	router := chi.NewRouter()
	// router.Use(hp.Autorize)
	router.Post("/", Create)

	log.Println("operations: routes registered")

	return router
}
