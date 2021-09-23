package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

var route *mux.Router

// Set route
func SetRoute(r *mux.Router) {
	route = r
}

// Convert route name to URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		// checkError(err)
		return ""
	}

	return url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
