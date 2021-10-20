package flash

import (
	"encoding/gob"
	"goblong/pkg/session"
)

// Flashes flash message struct
type Flashes map[string]interface{}

// Flash key
var flashKey = "_flashes"

func init() {
	// save on gorilla/sessions map and struct
	// before need serialization
	gob.Register(Flashes{})
}

// Info message
func Info(message string) {
	addFlash("info", message)
}

// Warning
func Warning(message string) {
	addFlash("warning", message)
}

// Success
func Success(message string) {
	addFlash("success", message)
}

// Danger
func Danger(message string) {
	addFlash("danger", message)
}

// Fetch all function
func All() Flashes {
	val := session.Get(flashKey)

	// Before read need check type
	flashMessage, ok := val.(Flashes)
	if !ok {
		return nil
	}

	// Destroy
	session.Forget(flashKey)

	return flashMessage

}

// Add new flash
func addFlash(key string, message string) {
	flashes := Flashes{}
	flashes[key] = message
	session.Put(flashKey, flashes)
	session.Save()
}
