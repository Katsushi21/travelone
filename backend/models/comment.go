package models

type Comment struct {
	ID        string  `json:"id"`
	User      *User   `json:"user"`
	Post      *Post   `json:"post"`
	Body      *string `json:"body"`
	CreatedAt *string `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt"`
}

type CommentInput struct {
	Body string `json:"body"`
}
