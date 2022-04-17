package bootstrap

import (
	"embed"
	"goblong/pkg/route"
	"goblong/routes"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoute(staticFS embed.FS) *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)

	// Static resources
	sub, _ := fs.Sub(staticFS, "public")
	router.PathPrefix("/").Handler(http.FileServer(http.FS(sub)))

	return router
}
