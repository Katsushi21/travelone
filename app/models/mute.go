package models

type Mute struct {
	UserID int   `json:"user_id"`
	MuteID int   `json:"mute_id"`
	Mute   *User `json:"mute"`
}

type MuteInput struct {
	UserID int `json:"user_id"`
	MuteID int `json:"mute_id"`
}
