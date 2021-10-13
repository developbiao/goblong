package routes

import (
	"github.com/gorilla/mux"
	"goblong/app/http/controllers"
	"goblong/app/http/middlewares"
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

	// Articles
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")

	// Create article page
	r.HandleFunc("/articles/create", ac.Create).
		Methods("GET").
		Name("articles.create")

	// Save article
	r.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store")

	// edit
	r.HandleFunc("/articles/{id:[0-9]+}/edit", ac.Edit).
		Methods("GET").
		Name("articles.edit")
	// update
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Update).
		Methods("POST").
		Name("articles.update")

	// Delete
	r.HandleFunc("/articles/{id:[0-9]+}/delete", ac.Delete).
		Methods("POST").
		Name("articles.delete")

	// User authorization
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", auc.Register).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", auc.DoRegister).Methods("POST").Name("auth.doregister")

	// Login
	r.HandleFunc("/auth/login", auc.Login).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", auc.DoLogin).Methods("POST").Name("auth.dologin")

	// Static resource
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// Middleware force content is HTML
	//r.Use(middlewares.ForceHTML)
	r.Use(middlewares.StartSession)

}
