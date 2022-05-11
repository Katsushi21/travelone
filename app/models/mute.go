package models

type Mute struct {
	AccountID int      `json:"accountId"`
	MuteID    int      `json:"muteId"`
	Mute      *Account `json:"mute"`
}

type MuteInput struct {
	AccountID int `json:"accountId"`
	MuteID    int `json:"muteId"`
}
