package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	email    string
	password string
}
