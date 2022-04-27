package models

type Like struct {
	PostID    int      `json:"post_id"`
	AccountID int      `json:"account_id"`
	Post      *Post    `json:"post"`
	Account   *Account `json:"account"`
}

type LikeInput struct {
	PostID    int `json:"post_id"`
	AccountID int `json:"account_id"`
}
