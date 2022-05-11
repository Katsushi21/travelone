package models

type Friend struct {
	AccountID int      `json:"accountId"`
	FriendID  int      `json:"friendId"`
	Friend    *Account `json:"friend"`
}

type FriendInput struct {
	AccountID int `json:"accountId"`
	FriendID  int `json:"friendId"`
}
