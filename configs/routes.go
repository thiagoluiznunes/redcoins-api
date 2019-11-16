package configs

import (
	"database/sql"
	"rv-api/components/user"

	"github.com/go-chi/chi"
)

// InitRoutes : init all routes from API
func InitRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	user.InitUserRoutes(db, router)

	return router
}
