package model

type Marker struct {
	ID     string  `json:"id"`
	PostID string  `json:"post"`
	Title  *string `json:"title"`
	Lat    *string `json:"lat"`
	Lng    *string `json:"lng"`
}
