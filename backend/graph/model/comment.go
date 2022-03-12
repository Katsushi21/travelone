package model

type Comment struct {
	ID           string  `json:"id"`
	UID          string  `json:"user"`
	PostID       string  `json:"post_id"`
	Body         *string `json:"body"`
	Registration *string `json:"registration"`
	Modification *string `json:"modification"`
}
