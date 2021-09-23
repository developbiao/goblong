package bootstrap

import (
	"github.com/gorilla/mux"
	"goblong/pkg/route"
	"goblong/routes"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)
	return router
}
