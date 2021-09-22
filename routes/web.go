package routes

import (
	"github.com/gorilla/mux"
	"goblong/app/http/controllers"
	"net/http"
)

// Register web routes
func RegisterWebRoutes(r *mux.Router) {
	pc := new(controllers.PagesController)

	// Static page
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	// Article pages
	ac := new(controllers.ArticlesController)
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")

}
