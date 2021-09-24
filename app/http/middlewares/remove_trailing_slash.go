package middlewares

import (
	"net/http"
	"strings"
)

// Remove trailing slash
func RemoveTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		// continue serve
		next.ServeHTTP(w, r)
	})
}
