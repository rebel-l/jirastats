package models

import (
	"time"
)

type Stats struct {
	Id int
	ProjectId int
	Open int
	Closed int
	New int
	CreatedAt time.Time
}

func NewStats() *Stats {
	s := new(Stats)
	s.CreatedAt = time.Now()
	return s
}
