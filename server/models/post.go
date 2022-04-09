package models

import "time"

type Post struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Img       string     `json:"img"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	User      *User      `json:"user"`
	Marker    *Marker    `json:"marker"`
	Like      []*Like    `json:"like"`
	Comment   []*Comment `json:"comment"`
}

type PostInput struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Img    string `json:"img"`
}
