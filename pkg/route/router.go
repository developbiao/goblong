package route

import (
	"github.com/gorilla/mux"
	"goblong/pkg/config"
	"goblong/pkg/logger"
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
		logger.LogError(err)
		return ""
	}

	return config.GetString("app.url") + url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
