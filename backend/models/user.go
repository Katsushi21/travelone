package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Email        *string   `json:"email"`
	Password     *string   `json:"password"`
	Type         *UserType `json:"type"`
	Session      *string   `json:"session"`
	Name         *string   `json:"name"`
	Age          *int      `json:"age"`
	Gender       *Gender   `json:"gender"`
	Avatar       *string   `json:"avatar"`
	Introduction *string   `json:"introduction"`
	Friends      []*int    `json:"friends"`
	Mute         []*int    `json:"mute"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserInput struct {
	Email        string   `json:"email"`
	Password     string   `json:"password"`
	Type         UserType `json:"type"`
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Gender       Gender   `json:"gender"`
	Avatar       string   `json:"avatar"`
	Introduction string   `json:"introduction"`
}

type SessionInput struct {
	Session string `json:"session"`
}
type FriendInput struct {
	Friends int `json:"friends"`
}

type MuteInput struct {
	Mute int `json:"mute"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderNone   Gender = "none"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
	GenderNone,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale, GenderNone:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserType string

const (
	UserTypeActive   UserType = "active"
	UserTypeInactive UserType = "inactive"
	UserTypeAdmin    UserType = "admin"
)

var AllUserType = []UserType{
	UserTypeActive,
	UserTypeInactive,
	UserTypeAdmin,
}

func (e UserType) IsValid() bool {
	switch e {
	case UserTypeActive, UserTypeInactive, UserTypeAdmin:
		return true
	}
	return false
}

func (e UserType) String() string {
	return string(e)
}

func (e *UserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}
	return nil
}

func (e UserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
