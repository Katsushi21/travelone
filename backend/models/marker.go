package models

import (
	"gorm.io/gorm"
)

type Marker struct {
	gorm.Model
	email    string
	password string
}
