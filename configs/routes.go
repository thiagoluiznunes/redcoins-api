package configs

import (
	"rv-api/components/user"

	"github.com/go-chi/chi"
)

// InitRoutes : init all routes from API
// func InitRoutes() *mux.Router {
func InitRoutes() *chi.Mux {
	// router := mux.NewRouter()
	router := chi.NewRouter()

	user.InitUserRoutes(router)

	return router
}
