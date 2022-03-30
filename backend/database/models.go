package database

import (
	"time"
)

type Comment struct {
	ID        int       `gorm:"primaryKey;type:autoIncrement"`
	PostID    int       `gorm:"not null;default:0"`
	UID       int       `gorm:"not null;default:0"`
	Body      string    `gorm:"not null;default:''"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Marker struct {
	ID        int       `gorm:"primaryKey;type:autoIncrement"`
	PostID    int       `gorm:"not null;default:0"`
	Title     string    `gorm:"not null;default:''"`
	Lat       string    `gorm:"not null;default:''"`
	Lng       string    `gorm:"not null;default:''"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Post struct {
	ID        int       `gorm:"primaryKey;type:autoIncrement"`
	UID       int       `gorm:"not null;default:0"`
	Title     string    `gorm:"not null;default:''"`
	Body      string    `gorm:"not null;default:''"`
	Img       string    `gorm:"not null;default:''"`
	Liked     []int     `gorm:"not null;default:[]"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Request struct {
	ID           int       `gorm:"primaryKey;type:autoIncrement"`
	RequestUID   int       `gorm:"not null;default:0"`
	RequestedUID int       `gorm:"not null;default:0"`
	Status       string    `gorm:"not null;default:''"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type User struct {
	ID           int       `gorm:"primaryKey;type:autoIncrement"`
	Email        string    `gorm:"unique;"`
	Password     string    `gorm:"not null;default:''"`
	Type         string    `gorm:"not null;default:''"`
	Session      string    `gorm:"not null;default:''"`
	Name         string    `gorm:"not null;default:''"`
	Age          int       `gorm:"not null;default:0"`
	Gender       string    `gorm:"not null;default:''"`
	Avatar       string    `gorm:"not null;default:''"`
	Introduction string    `gorm:"not null;default:''"`
	Friends      []int     `gorm:"not null;default:[]"`
	Mute         []int     `gorm:"not null;default:[]"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
