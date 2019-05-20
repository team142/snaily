package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Item struct {
	ID                 string    `json:"id"`
	Parent             string    `json:"parent"`
	Title              string    `json:"title"`
	Body               string    `json:"body"`
	CreateDate         time.Time `json:"createdDate"`
	CreatedBy          string    `json:"createdBy"`
	WaitingFor         string    `json:"waitingFor"`
	OrgID              string    `json:"orgID"`
	WaitingForDone     bool      `json:"waitingForDone"`
	WaitingForDoneDate time.Time `json:"waitingForDoneDate"`
	CreatedByDone      bool      `json:"createdByDone"`
	CreatedByDoneDate  time.Time `json:"createdByDoneDate"`
}

func (i *Item) GenerateID() {
	i.ID = uuid.NewV4().String()
}
