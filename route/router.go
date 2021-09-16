package route

import "github.com/gorilla/mux"

var Router *mux.Router

// Initialize
func Initialize() {
	Router = mux.NewRouter()
}

// Convert route name to URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		// checkError(err)
		return ""
	}

	return url.String()
}
