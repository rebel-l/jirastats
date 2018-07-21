package models

import (
	"fmt"
)

const projectDefaultJql = "project in (%s) AND type != Epic"
const projectDefaultKnownSpeed = 1.0
const projectDefaultMapOpen = "Open, In Progress, Reopened"
const projectDefaultMapClosed = "Closed, Resolved"

// Project represents data structure of a project and it's business logic
type Project struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	Keys            string  `json:"keys"`
	Jql             string  `json:"jql"`
	KnownSpeed      float32 `json:"known_speed"`
	MapOpenStatus   string  `json:"map_open_status"`
	MapClosedStatus string  `json:"map_closed_status"`
	// TODO: needs active flag
}

// NewProject returns a Project struct
func NewProject() *Project {
	p := new(Project)
	p.Jql = projectDefaultJql
	p.KnownSpeed = projectDefaultKnownSpeed
	p.MapOpenStatus = projectDefaultMapOpen
	p.MapClosedStatus = projectDefaultMapClosed
	return p
}

// GetJql return a JQL query based on project settings
func (p *Project) GetJql() string {
	return fmt.Sprintf(p.Jql, p.Keys)
}
