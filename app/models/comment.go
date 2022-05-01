package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	AccountID int       `json:"account_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Post      *Post     `json:"post"`
	Account   *Account  `json:"account"`
}

type CommentInput struct {
	PostID    int    `json:"post_id"`
	AccountID int    `json:"account_id"`
	Body      string `json:"body"`
}
