package models

type Post struct {
	ID           string    `json:"id"`
	UID          string    `json:"user"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	Img          *string   `json:"img"`
	Liked        []*string `json:"liked"`
	Registration *string   `json:"registration"`
	Modification *string   `json:"modification"`
}
