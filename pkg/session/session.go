package session

import (
	"github.com/gorilla/sessions"
	"goblong/pkg/logger"
	"net/http"
)

// Store gorilla sessions
var Store = sessions.NewCookieStore([]byte("33446a9dcf9ea060a0a6532b166da32f304af0de"))

// Current session
var Session *sessions.Session

// Request for get session
var Request *http.Request

// Response for set session
var Response http.ResponseWriter

// StartSession init session using with middleware
func StartSession(w http.ResponseWriter, r *http.Request) {
	var err error
	Session, err = Store.Get(r, "goblog-session")
	logger.LogError(err)

	Request = r
	Response = w
}

// Put write session
func Put(key string, value interface{}) {
	Session.Values[key] = value
	Save()
}

// Get session
func Get(key string) interface{} {
	return Session.Values[key]
}

// Forget session
func Forget(key string) {
	delete(Session.Values, key)
	Save()
}

// Destroy all session
func Flush() {
	Session.Options.MaxAge = -1
	Save()
}

// Save session
func Save() {
	// Not HTTPS link can't use Secure and HttpOnly
	// Session.Options.Secure = true
	// Session.Options.HttpOnly = true
	err := Session.Save(Request, Response)
	logger.LogError(err)
}
