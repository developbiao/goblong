package user

import (
	"goblong/pkg/password"
	"gorm.io/gorm"
)

// Before Save
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if !password.IsHashed(u.Password) {
		u.Password = password.Hash(u.Password)
	}
	return
}
