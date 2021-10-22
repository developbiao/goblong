package middlewares

import (
	"goblong/pkg/auth"
	"goblong/pkg/flash"
	"net/http"
)

func Guest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if auth.Check() {
			flash.Warning("Login user can not access this page")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next(w, r)
	}
}
