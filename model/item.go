package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Item struct {
	ID                 string
	Parent             string
	Title              string
	Body               string
	CreateDate         time.Time
	CreatedBy          string
	WaitingFor         string
	OrgID              string
	WaitingForDone     bool
	WaitingForDoneDate time.Time
	CreatedByDone      bool
	CreatedByDoneDate  time.Time
}

func (i *Item) GenerateID() {
	i.ID = uuid.NewV4().String()
}
