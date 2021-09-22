package controllers

import (
	"fmt"
	"net/http"
)

// PagesController process static page
type PagesController struct {
}

// Home page
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1 style=\"color:pink\">Hello, welcome my goblog</h1>")
}

// About page
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>This blog just for record learning golang, if you have any question please contact "+
		"<a href=\"mailto:developbiao@gmail.com\">developbiao@gmail.com</a></h1>")

}

// Not found page
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Request not found page :(</h1>"+
		"<p>If you have any doubts, please contact us. </p>")

}
