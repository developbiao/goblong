package routes

import (
	"goblong/app/http/controllers"
	"goblong/app/http/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Register web routes
func RegisterWebRoutes(r *mux.Router) {
	pc := new(controllers.PagesController)

	// Static page
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	// Article pages
	ac := new(controllers.ArticlesController)
	// Home page
	r.HandleFunc("/", ac.Index).Methods("GET").Name("home")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")

	// Articles
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")

	// Create article page
	r.HandleFunc("/articles/create", middlewares.Auth(ac.Create)).
		Methods("GET").
		Name("articles.create")

	// Save article
	r.HandleFunc("/articles", middlewares.Auth(ac.Store)).Methods("POST").Name("articles.store")

	// edit
	r.HandleFunc("/articles/{id:[0-9]+}/edit", middlewares.Auth(ac.Edit)).
		Methods("GET").
		Name("articles.edit")
	// update
	r.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(ac.Update)).
		Methods("POST").
		Name("articles.update")

	// Delete
	r.HandleFunc("/articles/{id:[0-9]+}/delete", middlewares.Auth(ac.Delete)).
		Methods("POST").
		Name("articles.delete")

	// Categories
	cc := new(controllers.CategoriesController)
	r.HandleFunc("/categories/create", middlewares.Auth(cc.Create)).
		Methods("GET").Name("categories.create")
	r.HandleFunc("/categories", middlewares.Auth(cc.Store)).
		Methods("POST").Name("categories.store")
	r.HandleFunc("/categories/{id:[0-9+]}/show", middlewares.Auth(cc.Show)).
		Methods("GET").Name("categories.show")

	// User authorization
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", middlewares.Guest(auc.Register)).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", middlewares.Guest(auc.DoRegister)).Methods("POST").Name("auth.doregister")

	// Login
	r.HandleFunc("/auth/login", middlewares.Guest(auc.Login)).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", middlewares.Guest(auc.DoLogin)).Methods("POST").Name("auth.dologin")

	// Logout
	r.HandleFunc("/auth/logout", middlewares.Auth(auc.Logout)).Methods("post").Name("auth.logout")

	// Middleware force content is HTML
	//r.Use(middlewares.ForceHTML)
	r.Use(middlewares.StartSession)

	// User authoriation
	uc := new(controllers.UserController)
	r.HandleFunc("/users/{id:[0-9]+}", uc.Show).Methods("GET").Name("users.show")

}
