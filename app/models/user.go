package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type User struct {
	ID           int        `json:"id"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	Type         UserType   `json:"type"`
	Name         string     `json:"name"`
	Age          int        `json:"age"`
	Gender       UserGender `json:"gender"`
	Avatar       string     `json:"avatar"`
	Introduction string     `json:"introduction"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	Friend       []*Friend  `json:"friend"`
	Mute         []*Mute    `json:"mute"`
	Post         []*Post    `json:"post"`
	Like         []*Like    `json:"like"`
	Comment      []*Comment `json:"comment"`
}

type UserInput struct {
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	Type         UserType   `json:"type"`
	Name         string     `json:"name"`
	Age          int        `json:"age"`
	Gender       UserGender `json:"gender"`
	Avatar       string     `json:"avatar"`
	Introduction string     `json:"introduction"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserGender string

const (
	UserGenderMale   UserGender = "male"
	UserGenderFemale UserGender = "female"
	UserGenderNone   UserGender = "none"
)

var AllUserGender = []UserGender{
	UserGenderMale,
	UserGenderFemale,
	UserGenderNone,
}

func (e UserGender) IsValid() bool {
	switch e {
	case UserGenderMale, UserGenderFemale, UserGenderNone:
		return true
	}
	return false
}

func (e UserGender) String() string {
	return string(e)
}

func (e *UserGender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserGender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserGender", str)
	}
	return nil
}

func (e UserGender) MarshalGQL(w io.Writer) {
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
