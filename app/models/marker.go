package models

import "time"

type Marker struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	Title     string    `json:"title"`
	Lat       string    `json:"lat"`
	Lng       string    `json:"lng"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Post      *Post     `json:"post"`
}

type MarkerInput struct {
	PostID int    `json:"post_id"`
	Title  string `json:"title"`
	Lat    string `json:"lat"`
	Lng    string `json:"lng"`
}
