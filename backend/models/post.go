package models

type Post struct {
	ID           string     `json:"id"`
	User         *User      `json:"user"`
	Title        string     `json:"title"`
	Body         string     `json:"body"`
	Img          *string    `json:"img"`
	Liked        []*User    `json:"liked"`
	Registration *string    `json:"registration"`
	Modification *string    `json:"modification"`
	Marker       *Marker    `json:"marker"`
	Comment      []*Comment `json:"comment"`
}

type PostInput struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Img   string `json:"img"`
}
