package models

type Friend struct {
	OwnID  int   `json:"own_id"`
	UserID int   `json:"user_id"`
	User   *User `json:"user"`
}

type FriendInput struct {
	OwnID  int `json:"own_id"`
	UserID int `json:"user_id"`
}
