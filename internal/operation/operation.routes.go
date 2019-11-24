package operation

import (
	"database/sql"
	"log"

	hp "redcoins-api/internal"

	"github.com/go-chi/chi"
)

// Routes : init all routes from user component
func Routes(db *sql.DB) chi.Router {
	DB = db
	InitOperationSchema()

	router := chi.NewRouter()
	router.Use(hp.AutorizeMiddleware)
	router.Post("/", Create)
	router.Get("/", GetByUser)
	router.Get("/{date}", GetByDate)

	log.Println("operations: routes registered")

	return router
}
