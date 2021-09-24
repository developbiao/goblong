package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblong/app/http/middlewares"
	"goblong/bootstrap"
	"net/http"
)

var router *mux.Router

func main() {

	// Setup ORM
	bootstrap.SetupDB()

	// initialize router
	router = bootstrap.SetupRoute()

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

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
