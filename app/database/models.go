package database

import (
	"time"
)

type Comment struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	PostID    int       `gorm:"not null;default:0"`
	UserID    int       `gorm:"not null;default:0"`
	Body      string    `gorm:"not null;default:''"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Post      Post
	User      User
}

type Marker struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	PostID    int       `gorm:"unique;not null;default:0"`
	Title     string    `gorm:"not null;default:''"`
	Lat       string    `gorm:"not null;default:''"`
	Lng       string    `gorm:"not null;default:''"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Post      Post
}

type Post struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"not null;default:0"`
	Title     string    `gorm:"not null;default:''"`
	Body      string    `gorm:"not null;default:''"`
	Img       string    `gorm:"not null;default:''"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	User      User
}

type Like struct {
	PostID    int       `gorm:"primaryKey"`
	UserID    int       `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Post      Post
	User      User
}

type Request struct {
	UserID       int       `gorm:"primaryKey"`
	TargetUserID int       `gorm:"primaryKey"`
	Status       string    `gorm:"not null;default:''"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	User         User
	TargetUser   User
}

type User struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	Email        string    `gorm:"unique;not null;default:''"`
	Password     string    `gorm:"not null;default:''"`
	Type         string    `gorm:"not null;default:''"`
	Name         string    `gorm:"not null;default:''"`
	Age          int       `gorm:"not null;default:0"`
	Gender       string    `gorm:"not null;default:''"`
	Avatar       string    `gorm:"not null;default:''"`
	Introduction string    `gorm:"not null;default:''"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type Friend struct {
	OwnID     int       `gorm:"primaryKey;foreignKey:ID"`
	UserID    int       `gorm:"primaryKey;foreignKey:ID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	User      User
}

type Mute struct {
	OwnID     int       `gorm:"primaryKey;foreignKey:ID"`
	UserID    int       `gorm:"primaryKey;foreignKey:ID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	User      User
}

type Session struct {
	UserID    int       `gorm:"primaryKey"`
	Session   string    `gorm:"primaryKey"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	User      User
}
