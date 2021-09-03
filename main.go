package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1 style=\"color:pink\">Hello, here is my goblog</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "<h1>This blog just for record leanring golang, if you have any quetion please contact "+
			"<a href=\"mailto:developbiao@gmail.com\">developbiao@gmail.com</a></h1>")
	} else {
		fmt.Fprint(w, "<h1>Reqeust not found page :(</h1>"+
			"<p>If you have any doubts, please contact us. </p>")

	}

}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", nil)
}
