package berita

import (
	"gorm.io/gorm"
)

type News struct {
	*gorm.Model
	Title   string `json:"title" form:"title" `
	Author  string `json:"author" form:"author" `
	Content string `json:"content" form:"content" `
}

type Berita2 struct {
	*gorm.Model

	Title   string `json:"title" form:"title" `
	Author  string `json:"author" form:"author" `
	Content string `json:"content" form:"content" `
}
