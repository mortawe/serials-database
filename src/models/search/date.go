package search

import "time"

type BetweenDate struct {
	Before time.Time `json:"before"`
	After time.Time `json:"after"`
}

