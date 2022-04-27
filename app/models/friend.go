package models

type Friend struct {
	AccountID int      `json:"account_id"`
	FriendID  int      `json:"friend_id"`
	Friend    *Account `json:"friend"`
}

type FriendInput struct {
	AccountID int `json:"account_id"`
	FriendID  int `json:"friend_id"`
}
