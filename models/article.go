// models/article.go
package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title   string `json:"title" gorm:"type:varchar(100)"`
	Content string `json:"content" gorm:"type:text"`
}
