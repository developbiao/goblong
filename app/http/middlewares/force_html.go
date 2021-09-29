package middlewares

import (
	"net/http"
)

// force add html header
func ForceHTML(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Set header
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// 2. Continue request
		next.ServeHTTP(w, r)
	})
}
