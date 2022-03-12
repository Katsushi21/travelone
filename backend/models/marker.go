package models

import (
	"gorm.io/gorm"
)

type Marker struct {
	gorm.Model
	PostId uint
	Title  string
	Lat    string
	Lng    string
}
