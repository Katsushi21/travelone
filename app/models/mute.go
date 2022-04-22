package models

type Mute struct {
	OwnID  int   `json:"own_id"`
	UserID int   `json:"user_id"`
	User   *User `json:"user"`
}

type MuteInput struct {
	OwnID  int `json:"own_id"`
	UserID int `json:"user_id"`
}
