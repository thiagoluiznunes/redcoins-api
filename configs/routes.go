package configs

import (
	"rv-api/components/user"

	"github.com/gorilla/mux"
)

// InitRoutes : init all routes from API
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	user.InitUserRoutes(router)

	return router
}
