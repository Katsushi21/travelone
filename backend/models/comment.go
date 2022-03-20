package models

import "time"

type Comment struct {
	ID        int       `json:"id" gorm:"column:id"`
	User      *int      `json:"user"`
	Post      *int      `json:"post"`
	Body      *string   `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CommentInput struct {
	Body string `json:"body"`
}
