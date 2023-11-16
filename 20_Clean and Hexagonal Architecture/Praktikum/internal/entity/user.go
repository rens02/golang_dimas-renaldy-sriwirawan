package entity

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Name        string `json:"name" form:"name" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email" `
	Password    string `json:"password" form:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
}
