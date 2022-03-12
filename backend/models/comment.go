package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostId uint
	Uid    uint
	Body   string
}
