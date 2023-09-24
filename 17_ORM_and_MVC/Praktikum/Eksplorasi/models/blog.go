package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	UsersID uint   `json:"users_id" form:"users_id"`
	User    User   `json:"user" gorm:"foreignkey:UsersID"` // Explicitly define the foreign key
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}
