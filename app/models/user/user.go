package user

import (
	"fmt"
	"goblong/app/models"
	"goblong/pkg/password"
)

type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);default:NULL;unique" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`
	// gorm: "-" set GORM read and write ignore this field
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

// Compare user password from client
func (u User) ComparePassword(_password string) bool {
	fmt.Println("Plaintext password:", _password)
	return password.CheckHash(_password, u.Password)
}
