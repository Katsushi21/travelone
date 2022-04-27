package models

type Mute struct {
	AccountID int      `json:"account_id"`
	MuteID    int      `json:"mute_id"`
	Mute      *Account `json:"mute"`
}

type MuteInput struct {
	AccountID int `json:"account_id"`
	MuteID    int `json:"mute_id"`
}
