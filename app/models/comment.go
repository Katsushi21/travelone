package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	AccountID int       `json:"account_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Post      *Post     `json:"post"`
	Account   *Account  `json:"account"`
}

type CommentInput struct {
	PostID    int    `json:"post_id"`
	AccountID int    `json:"account_id"`
	Body      string `json:"body"`
}
