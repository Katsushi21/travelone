package models

type Marker struct {
	ID    string  `json:"id"`
	Post  *Post   `json:"post"`
	Title *string `json:"title"`
	Lat   *string `json:"lat"`
	Lng   *string `json:"lng"`
}

type MarkerInput struct {
	Title string `json:"title"`
	Lat   string `json:"lat"`
	Lng   string `json:"lng"`
}
