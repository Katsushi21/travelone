package models

type Friend struct {
	UserID   int   `json:"user_id"`
	FriendID int   `json:"friend_id"`
	Friend   *User `json:"friend"`
}

type FriendInput struct {
	UserID   int `json:"user_id"`
	FriendID int `json:"friend_id"`
}
