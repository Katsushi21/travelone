package model

type Request struct {
	ID        string        `json:"id"`
	Request   string        `json:"request"`
	Requested string        `json:"requested"`
	Status    RequestStatus `json:"status"`
}
