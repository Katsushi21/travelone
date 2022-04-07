package models

type Mute struct {
	UID       *int    `json:"uid"`
	TargetUID *int    `json:"target_uid"`
	User      []*User `json:"user"`
}

type MuteInput struct {
	UID       int `json:"uid"`
	TargetUID int `json:"target_uid"`
}
