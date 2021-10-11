package user

import (
	"goblong/pkg/logger"
	"goblong/pkg/model"
)

// Create user  User.id exists  is create success otherwise is failed
func (user User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
