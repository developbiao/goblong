package user

import "goblong/app/models"

type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);default:NULL;unique" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`
	// gorm: "-" set GORM read and write ignore this field
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}
