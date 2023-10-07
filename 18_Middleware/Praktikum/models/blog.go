package models

import (
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	gorm.Model
	UsersID uint   `json:"users_id" form:"users_id"`
	User    User   `json:"user" gorm:"foreignkey:UsersID"` // Explicitly define the foreign key
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

type BlogResponse struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User      `json:"user" ` // Explicitly define the foreign key
	Title     string    `json:"title" `
	Content   string    `json:"content"`
}

func (blogDB Blog) ResponseConvert() BlogResponse {
	var Response BlogResponse
	Response.ID = blogDB.ID
	Response.Title = blogDB.Title
	Response.Content = blogDB.Content

	return Response

}
