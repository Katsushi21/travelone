package models

import (
	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	Request   uint
	Requested uint
	Status    string
}
