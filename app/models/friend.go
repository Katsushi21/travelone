package models

type Friend struct {
	UserID   int   `json:"user_id"`
	FriendID int   `json:"friend_id"`
	User     *User `json:"user"`
}

type FriendInput struct {
	UserID   int `json:"user_id"`
	FriendID int `json:"friend_id"`
}
