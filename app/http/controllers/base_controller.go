package controllers

import (
	"fmt"
	"goblong/pkg/flash"
	"goblong/pkg/logger"
	"gorm.io/gorm"
	"net/http"
)

// BaseController
type BaseController struct {
}

func (bc BaseController) ResponseFromSQLError(w http.ResponseWriter, err error) {
	if err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 file not found")
	} else {
		// database error
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 server internal server error")
	}
}

// Response unauthorized
func (bc BaseController) ResponseFromUnauthorized(w http.ResponseWriter, r *http.Request) {
	flash.Warning("UnAuthorized")
	http.Redirect(w, r, "/", http.StatusFound)
}
