package database

import (
	"time"
)

type Account struct {
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

type Comment struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	PostID    int       `gorm:"not null;default:0"`
	AccountID int       `gorm:"not null;default:0"`
	Body      string    `gorm:"not null;default:''"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Post      Post
	Account   Account
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
	AccountID int       `gorm:"not null;default:0"`
	Title     string    `gorm:"not null;default:''"`
	Body      string    `gorm:"not null;default:''"`
	Img       string    `gorm:"not null;default:''"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Account   Account
}

type Like struct {
	PostID    int       `gorm:"primaryKey"`
	AccountID int       `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Post      Post
	Account   Account
}

type Request struct {
	AccountID       int       `gorm:"primaryKey"`
	TargetAccountID int       `gorm:"primaryKey"`
	Status          string    `gorm:"not null;default:''"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	Account         Account
	TargetAccount   Account
}

type Friend struct {
	AccountID int       `gorm:"primaryKey;foreignKey:ID"`
	FriendID  int       `gorm:"primaryKey;foreignKey:ID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Friend    Account
}

type Mute struct {
	AccountID int       `gorm:"primaryKey;foreignKey:ID"`
	MuteID    int       `gorm:"primaryKey;foreignKey:ID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Mute      Account
}

type Session struct {
	AccountID int       `gorm:"primaryKey"`
	Session   string    `gorm:"primaryKey"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Account   Account
}
