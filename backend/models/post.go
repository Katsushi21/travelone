package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	UID       int       `json:"uid"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Img       string    `json:"img"`
	Liked     []*int    `json:"liked"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PostInput struct {
	ID    *int   `json:"id"`
	UID   *int   `json:"uid"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Img   string `json:"img"`
}

type LikedInput struct {
	PostID int `json:"post_id"`
	UID    int `json:"uid"`
}
