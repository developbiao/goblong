package bootstrap

import (
	"github.com/gorilla/mux"
	"goblong/routes"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	return router
}
