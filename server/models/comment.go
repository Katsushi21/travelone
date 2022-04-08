package models

import "time"

type Comment struct {
	ID        int        `json:"id"`
	PostID    *int       `json:"post_id"`
	UserID    *int       `json:"user_id"`
	Body      string     `json:"body"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Post      *Post      `json:"post"`
	User      *User      `json:"user"`
}

type CommentInput struct {
	PostID int    `json:"post_id"`
	UserID int    `json:"user_id"`
	Body   string `json:"body"`
}
