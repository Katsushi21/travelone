package models

type Like struct {
	PostID int   `json:"post_id"`
	UserID int   `json:"user_id"`
	Post   *Post `json:"post"`
	User   *User `json:"user"`
}

type LikeInput struct {
	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
}
