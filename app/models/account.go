package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Account struct {
	ID           int           `json:"id"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Type         AccountType   `json:"type"`
	Name         string        `json:"name"`
	Age          int           `json:"age"`
	Gender       AccountGender `json:"gender"`
	Avatar       string        `json:"avatar"`
	Introduction string        `json:"introduction"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	Friend       []*Friend     `json:"friend"`
	Mute         []*Mute       `json:"mute"`
	Post         []*Post       `json:"post"`
	Like         []*Like       `json:"like"`
	Comment      []*Comment    `json:"comment"`
}

type AccountInput struct {
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Type         AccountType   `json:"type"`
	Name         string        `json:"name"`
	Age          int           `json:"age"`
	Gender       AccountGender `json:"gender"`
	Avatar       string        `json:"avatar"`
	Introduction string        `json:"introduction"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountGender string

const (
	AccountGenderMale   AccountGender = "male"
	AccountGenderFemale AccountGender = "female"
	AccountGenderNone   AccountGender = "none"
)

var AllAccountGender = []AccountGender{
	AccountGenderMale,
	AccountGenderFemale,
	AccountGenderNone,
}

func (e AccountGender) IsValid() bool {
	switch e {
	case AccountGenderMale, AccountGenderFemale, AccountGenderNone:
		return true
	}
	return false
}

func (e AccountGender) String() string {
	return string(e)
}

func (e *AccountGender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AccountGender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AccountGender", str)
	}
	return nil
}

func (e AccountGender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AccountType string

const (
	AccountTypeActive   AccountType = "active"
	AccountTypeInactive AccountType = "inactive"
	AccountTypeAdmin    AccountType = "admin"
)

var AllAccountType = []AccountType{
	AccountTypeActive,
	AccountTypeInactive,
	AccountTypeAdmin,
}

func (e AccountType) IsValid() bool {
	switch e {
	case AccountTypeActive, AccountTypeInactive, AccountTypeAdmin:
		return true
	}
	return false
}

func (e AccountType) String() string {
	return string(e)
}

func (e *AccountType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AccountType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AccountType", str)
	}
	return nil
}

func (e AccountType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
