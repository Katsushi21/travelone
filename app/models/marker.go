package models

import "time"

type Marker struct {
	ID        int       `json:"id"`
	PostID    int       `json:"postId"`
	Title     string    `json:"title"`
	Lat       string    `json:"lat"`
	Lng       string    `json:"lng"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Post      *Post     `json:"post"`
}

type MarkerInput struct {
	PostID int    `json:"postId"`
	Title  string `json:"title"`
	Lat    string `json:"lat"`
	Lng    string `json:"lng"`
}
