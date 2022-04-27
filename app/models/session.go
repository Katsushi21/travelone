package models

import "time"

type Session struct {
	AccountID int       `json:"account_id"`
	Session   string    `json:"session"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SessionInput struct {
	AccountID int    `json:"account_id"`
	Session   string `json:"session"`
}
