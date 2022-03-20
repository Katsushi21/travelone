package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Request struct {
	ID           int           `json:"id"`
	RequestUID   *User         `json:"request_uid"`
	RequestedUID *User         `json:"requested_uid"`
	Status       RequestStatus `json:"status"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
}

type RequestInput struct {
	RequestedUID *int          `json:"requested_uid"`
	Status       RequestStatus `json:"status"`
}

type RequestStatus string

const (
	RequestStatusInProcess      RequestStatus = "in_process"
	RequestStatusAccept         RequestStatus = "accept"
	RequestStatusDeny           RequestStatus = "deny"
	RequestStatusBreakInProcess RequestStatus = "break_in_process"
	RequestStatusBreakAccept    RequestStatus = "break_accept"
	RequestStatusBreakDeny      RequestStatus = "break_deny"
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
