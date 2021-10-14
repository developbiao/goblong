package user

import (
	"goblong/pkg/password"
	"gorm.io/gorm"
)

// Before create hook
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = password.Hash(u.Password)
	return
}

// Update hook
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if !password.IsHashed(u.Password) {
		u.Password = password.Hash(u.Password)
	}
	return
}

// BeforeSave hook
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if !password.IsHashed(u.Password) {
		u.Password = password.Hash(u.Password)
	}
	return
}
