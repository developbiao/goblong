package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1 style=\"color:pink\">Hello, here is my goblog</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Reqeust not found page :(</h1>"+
			"<p>If you have any doubts, please contact us. </p>")

	}
}

func aboutHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>This blog just for record leanring golang, if you have any quetion please contact "+
		"<a href=\"mailto:developbiao@gmail.com\">developbiao@gmail.com</a></h1>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.HandleFunc("/about", aboutHandle)
	http.ListenAndServe(":3000", nil)
}
