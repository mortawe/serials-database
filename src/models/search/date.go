package search

import "time"

type BetweenDate struct {
	Before time.Time `json:"before" db:"before"`
	After  time.Time `json:"after" db:"after"`
}
