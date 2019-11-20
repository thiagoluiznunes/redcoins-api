package configs

import (
	"database/sql"
	"redcoins-api/internal/operation"
	"redcoins-api/internal/user"

	"github.com/go-chi/chi"
)

// InitRoutes : init all routes from API
func InitRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	user.InitUserRoutes(db, router)
	operation.InitOperationRoutes(db, router)

	return router
}
