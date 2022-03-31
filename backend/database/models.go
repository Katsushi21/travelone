package database

import (
	"time"
)

type Comments struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	PostID    int       `gorm:"not null;default:0"`
	UID       int       `gorm:"not null;default:0"`
	Body      string    `gorm:"not null;default:''"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Markers struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	PostID    int       `gorm:"not null;default:0"`
	Title     string    `gorm:"not null;default:''"`
	Lat       string    `gorm:"not null;default:''"`
	Lng       string    `gorm:"not null;default:''"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Posts struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UID       int       `gorm:"not null;default:0"`
	Title     string    `gorm:"not null;default:''"`
	Body      string    `gorm:"not null;default:''"`
	Img       string    `gorm:"not null;default:''"`
	Liked     int       `gorm:"type:integer [] NOT NULL DEFAULT '{}'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Requests struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	RequestUID   int       `gorm:"not null;default:0"`
	RequestedUID int       `gorm:"not null;default:0"`
	Status       string    `gorm:"not null;default:''"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type Users struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	Email        string    `gorm:"unique;not null;default:''"`
	Password     string    `gorm:"not null;default:''"`
	Type         string    `gorm:"not null;default:''"`
	Session      string    `gorm:"not null;default:''"`
	Name         string    `gorm:"not null;default:''"`
	Age          int       `gorm:"not null;default:0"`
	Gender       string    `gorm:"not null;default:''"`
	Avatar       string    `gorm:"not null;default:''"`
	Introduction string    `gorm:"not null;default:''"`
	Friends      int       `gorm:"type:integer [] NOT NULL DEFAULT '{}'"`
	Mute         int       `gorm:"type:integer [] NOT NULL DEFAULT '{}'"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
