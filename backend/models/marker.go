package models

import "time"

type Marker struct {
	ID        int       `json:"id"`
	Post      *Post     `json:"post"`
	Title     *string   `json:"title"`
	Lat       *string   `json:"lat"`
	Lng       *string   `json:"lng"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MarkerInput struct {
	Title string `json:"title"`
	Lat   string `json:"lat"`
	Lng   string `json:"lng"`
}
