package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title    string
	Text     string
	AuthorId uint
}
