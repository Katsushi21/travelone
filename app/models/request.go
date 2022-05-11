package models

import (
	"fmt"
	"io"
	"strconv"
)

type Request struct {
	AccountID       int           `json:"accountId"`
	TargetAccountID int           `json:"targetAccountId"`
	Status          RequestStatus `json:"status"`
	Account         *Account      `json:"account"`
	TargetAccount   *Account      `json:"targetAccount"`
}

type RequestInput struct {
	AccountID       int           `json:"accountId"`
	TargetAccountID int           `json:"targetAccountId"`
	Status          RequestStatus `json:"status"`
}

type RequestStatus string

const (
	RequestStatusInProcess      RequestStatus = "inProcess"
	RequestStatusAccept         RequestStatus = "accept"
	RequestStatusDeny           RequestStatus = "deny"
	RequestStatusBreakInProcess RequestStatus = "breakInProcess"
	RequestStatusBreakAccept    RequestStatus = "breakAccept"
	RequestStatusBreakDeny      RequestStatus = "breakDeny"
	RequestStatusCancel         RequestStatus = "cancel"
)

var AllRequestStatus = []RequestStatus{
	RequestStatusInProcess,
	RequestStatusAccept,
	RequestStatusDeny,
	RequestStatusBreakInProcess,
	RequestStatusBreakAccept,
	RequestStatusBreakDeny,
	RequestStatusCancel,
}

func (e RequestStatus) IsValid() bool {
	switch e {
	case RequestStatusInProcess, RequestStatusAccept, RequestStatusDeny, RequestStatusBreakInProcess, RequestStatusBreakAccept, RequestStatusBreakDeny, RequestStatusCancel:
		return true
	}
	return false
}

func (e RequestStatus) String() string {
	return string(e)
}

func (e *RequestStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RequestStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RequestStatus", str)
	}
	return nil
}

func (e RequestStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
