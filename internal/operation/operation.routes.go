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
	router.Use(hp.AuthorizeMiddleware)
	router.Post("/", Create)
	router.Get("/", GetByUser)
	router.Post("/email", GetByEmail)
	router.Get("/date/{date}", GetByDate)
	router.Get("/name/{name}", GetByName)
	router.Delete("/test", DeleteTestOperations)

	log.Println("operations: routes registered")

	return router
}
