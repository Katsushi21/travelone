package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Uid   uint
	Title string
	Body  string
	Img   string
	Liked []uint
}
