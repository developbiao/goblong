package auth

import (
	"errors"
	"fmt"
	"goblong/app/models/user"
	"goblong/pkg/session"
	"gorm.io/gorm"
)

// Get uid from session
func _getUID() string {
	_uid := session.Get("uid")
	if _uid == nil {
		return ""
	}
	uid := fmt.Sprintf("%v", _uid)
	if len(uid) > 0 {
		return uid
	}
	return ""
}

func User() user.User {
	uid := _getUID()
	if len(uid) > 0 {
		_user, err := user.Get(uid)
		if err == nil {
			return _user
		}
	}
	return user.User{}
}

func Attempt(email string, password string) error {
	// 1. Get user by email
	_user, err := user.GetByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Account or password incorrect")
		} else {
			return errors.New("Interal server error")
		}
	}
	// 2. Compare password with database
	if !_user.ComparePassword(password) {
		return errors.New("Account or password error")
	}

	// 3. Save session status
	session.Put("uid", _user.ID)

	return nil
}

// Login
func Login(_user user.User) {

	session.Put("uid", _user.GetStringID())
}

// Logout
func Logout() {
	session.Forget("uid")
}

func Check() bool {
	return len(_getUID()) > 0
}
