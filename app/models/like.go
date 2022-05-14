package models

import "time"

type Like struct {
	PostID    int       `json:"postId"`
	AccountID int       `json:"accountId"`
	CreatedAt time.Time `json:"createdAt"`
	Post      *Post     `json:"post"`
	Account   *Account  `json:"account"`
}

type LikeInput struct {
	PostID    int `json:"postId"`
	AccountID int `json:"accountId"`
}
