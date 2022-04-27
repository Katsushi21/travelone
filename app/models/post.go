package models

import "time"

type Post struct {
	ID        int        `json:"id"`
	AccountID int        `json:"account_id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Img       string     `json:"img"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Account   *Account   `json:"account"`
	Marker    *Marker    `json:"marker"`
	Like      []*Like    `json:"like"`
	Comment   []*Comment `json:"comment"`
}

type PostInput struct {
	AccountID int    `json:"account_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Img       string `json:"img"`
}
