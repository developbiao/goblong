package user

import "goblong/app/models"

type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;unique"`
	Email    string `gorm:"column:email;type:varchar(255);default:NULL;unique"`
	Password string `gorm:"column:passwrod:type:varchar(255)"`
	// gorm: "-" set GORM read and write ignore this field
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}
