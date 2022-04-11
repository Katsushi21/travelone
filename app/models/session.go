package models

import "time"

type Session struct {
	UserID    int       `json:"user_id"`
	Session   *string   `json:"session"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SessionInput struct {
	UserID  int    `json:"user_id"`
	Session string `json:"session"`
}
