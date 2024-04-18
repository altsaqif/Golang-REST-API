package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string `gorm:"type:varchar(300);not null" json:"name"`
	Email    string `gorm:"type:varchar(300);not null;unique" json:"email"`
	Password string `gorm:"type:varchar(300);not null" json:"password"`
}

type Register struct {
	Name            string `gorm:"not null" json:"name"`
	Email           string `gorm:"not null;unique" json:"email"`
	Password        string `gorm:"not null" json:"password"`
	PasswordConfirm string `gorm:"not null" json:"password_confirm"`
}

type Login struct {
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
