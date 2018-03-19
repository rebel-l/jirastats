package models

const projectDefaultJql = "project in (%s)"
const projectDefaultKnownSpeed = 1.0
const projectDefaultMapOpen = "\"Open\", \"In Progress\", \"Reopened\""
const projectDefaultMapClosed = "\"Closed\", \"Resolved\""

type Project struct {
	Id int 		`json:"id"`
	Name string `json:"name"`
	Keys string `json:"keys"`
	Jql string	`json:"jql"`
	KnownSpeed float32 `json:"known_speed"`
	MapOpenStatus string `json:"map_open_status"`
	MapClosedStatus string `json:"map_closed_status"`
}

func NewProject() *Project {
	p := new(Project)
	p.Jql = projectDefaultJql
	p.KnownSpeed = projectDefaultKnownSpeed
	p.MapOpenStatus = projectDefaultMapOpen
	p.MapClosedStatus = projectDefaultMapClosed
	return p
}
