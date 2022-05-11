package models

type Like struct {
	PostID    int      `json:"postId"`
	AccountID int      `json:"accountId"`
	Post      *Post    `json:"post"`
	Account   *Account `json:"account"`
}

type LikeInput struct {
	PostID    int `json:"postId"`
	AccountID int `json:"accountId"`
}
