package main

import (
	"embed"
	"fmt"
	"goblong/app/http/middlewares"
	"goblong/bootstrap"
	"goblong/config"
	c "goblong/pkg/config"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//go:embed resources/views/articles/*
//go:embed resources/views/auth/*
//go:embed resources/views/categories/*
//go:embed resources/views/layouts/*
var tplFS embed.FS

//go:embed public/*
var staticFS embed.FS

var router *mux.Router

func init() {
	// Initialization config
	config.Initialize()
}

func main() {

	// Setup ORM
	bootstrap.SetupDB()

	// Initialization template
	bootstrap.SetupTemplate(tplFS)

	// initialize router
	router = bootstrap.SetupRoute(staticFS)

	// Get router name URL example
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "23")
	fmt.Println("articleURL: ", articleURL)

	router.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprint(w, "From [GET] request ")
		case "POST":
			fmt.Fprint(w, "From [POST] request")
		}
	})

	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
