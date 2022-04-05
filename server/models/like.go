package models

type Like struct {
	PostID *int  `json:"post_id"`
	UID    *int  `json:"uid"`
	Post   *Post `json:"post"`
	User   *User `json:"user"`
}

type LikeInput struct {
	PostID int `json:"post_id"`
	UID    int `json:"uid"`
}
