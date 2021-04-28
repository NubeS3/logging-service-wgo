package common

import "time"

type Event struct {
	Type string    `json:"type"`
	Date time.Time `json:"at"`
}
