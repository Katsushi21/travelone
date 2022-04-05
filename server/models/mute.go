package models

type Mute struct {
	UID        *int  `json:"uid"`
	TargetUID  *int  `json:"target_uid"`
	User       *User `json:"user"`
	TargetUser *User `json:"target_user"`
}

type MuteInput struct {
	UID       int `json:"uid"`
	TargetUID int `json:"target_uid"`
}
