package database

import (
	"time"
)

type Comments struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	PostID    int       `gorm:"not null;default:0"`
	UID       int       `gorm:"not null;default:0;foreignKey:Users"`
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
	UID       int       `gorm:"not null;default:0;foreignKey:Users"`
	Title     string    `gorm:"not null;default:''"`
	Body      string    `gorm:"not null;default:''"`
	Img       string    `gorm:"not null;default:''"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Likes struct {
	PostID    int       `gorm:"primaryKey"`
	UID       int       `gorm:"primaryKey;foreignKey:Users"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Requests struct {
	RequestUID   int       `gorm:"primaryKey"`
	RequestedUID int       `gorm:"primaryKey"`
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
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type Friends struct {
	UID       int       `gorm:"primaryKey;foreignKey:Users"`
	TargetUID int       `gorm:"primaryKey;foreignKey:Users"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Mutes struct {
	UID       int       `gorm:"primaryKey;foreignKey:Users"`
	TargetUID int       `gorm:"primaryKey;foreignKey:Users"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
