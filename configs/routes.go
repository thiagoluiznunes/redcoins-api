package configs

import (
	"database/sql"
	"redcoins-api/internal/operation"
	"redcoins-api/internal/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// InitRoutes : init all routes from API
func InitRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Use(Cors().Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Mount("/api/v1/users", user.Routes(db))
	router.Mount("/api/v1/operations", operation.Routes(db))

	return router
}
