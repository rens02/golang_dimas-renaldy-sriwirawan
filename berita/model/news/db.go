package news

import "gorm.io/gorm"

type News struct {
	*gorm.Model
	Title   string
	Author  string
	content string
}
