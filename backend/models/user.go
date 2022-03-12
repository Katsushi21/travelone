package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string
	Password     string
	Type         string
	Session      string
	Name         string
	Age          uint8
	Gender       string
	Avater       string
	Introduction string
	Friends      []uint
	Mute         []uint
}
