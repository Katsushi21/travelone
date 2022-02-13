package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	email    string
	password string
}
