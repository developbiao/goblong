package middlewares

import (
	"goblong/pkg/session"
	"net/http"
)

// Start session open session
func StartSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Start session
		session.StartSession(w, r)

		// 2. Next
		next.ServeHTTP(w, r)
	})
}
