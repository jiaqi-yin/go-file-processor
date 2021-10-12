package domain

import "time"

type Item struct {
	UUID      string
	Firstname string
	Lastname  string
	Created   time.Time
}
