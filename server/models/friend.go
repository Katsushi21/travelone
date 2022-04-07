package models

type Friend struct {
	UID       *int    `json:"uid"`
	TargetUID *int    `json:"target_uid"`
	User      []*User `json:"user"`
}

type FriendInput struct {
	UID       int `json:"uid"`
	TargetUID int `json:"target_uid"`
}
