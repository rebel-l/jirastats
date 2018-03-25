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

func NewStats(projectId int) *Stats {
	s := new(Stats)
	s.CreatedAt = time.Now()
	s.ProjectId = projectId
	return s
}
