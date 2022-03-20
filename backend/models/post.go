package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	UID       *int      `json:"user"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Img       *string   `json:"img"`
	Liked     []*int    `json:"liked"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PostInput struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Img   string `json:"img"`
}

type LikedInput struct {
	UID int `json:"uid"`
}
