package model

import "time"

type Item struct {
	ID                  string
	Parent              string
	Title               string
	Body                string
	CreateDate          time.Time
	Thrower             string
	Holder              string
	OrganizationID      string
	HolderSaysDone      bool
	HolderSaysDoneTime  time.Time
	ThrowerSaysDone     bool
	ThrowerSaysDoneTime time.Time
}
