package models

import "time"

type Marker struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	Title     string    `json:"title"`
	Lat       string    `json:"lat"`
	Lng       string    `json:"lng"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MarkerInput struct {
	ID     *int   `json:"id"`
	PostID *int   `json:"post_id"`
	Title  string `json:"title"`
	Lat    string `json:"lat"`
	Lng    string `json:"lng"`
}
