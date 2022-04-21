package models

type Mute struct {
	UserID int   `json:"user_id"`
	MuteID int   `json:"mute_id"`
	User   *User `json:"user"`
}

type MuteInput struct {
	UserID int `json:"user_id"`
	MuteID int `json:"mute_id"`
}
