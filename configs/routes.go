package configs

import (
	"database/sql"
	"redcoins-api/internal/operation"
	"redcoins-api/internal/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// InitRoutes : init all routes from API
func InitRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Mount("/api/v1/users", user.Routes(db))
	router.Mount("/api/v1/operations", operation.Routes(db))

	return router
}
