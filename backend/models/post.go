package models

import "time"

type Post struct {
	ID        int        `json:"id"`
	User      *User      `json:"user"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Img       *string    `json:"img"`
	Liked     []*User    `json:"liked"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Marker    *Marker    `json:"marker"`
	Comment   []*Comment `json:"comment"`
}

type PostInput struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Img   string `json:"img"`
}

type LikedInput struct {
	UID int `json:"uid"`
}
