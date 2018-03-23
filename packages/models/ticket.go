package models

import (
	"time"
)

const TicketStatusClusteredOpen = "Open"
const TicketStatusClusteredClosed = "Closed"

type Ticket struct {
	Id int `json:"id"`
	Key string `json:"key"`
	ProjectId int `json:"project_id`
	Summary string `json:"summary"`
	Components string `json:"components"`
	Labels string `json:"labels"`
	StatusByJira string `json:"status_by_jira"`
	StatusClustered string `json:"status_clustered"`
	Priority string `json:"priority"`
	Issuetype string `json:"issuetype"`
	CreatedAtByJira time.Time `json:"created_at_by_jira"`
	LastUpdatedByJira time.Time `json:"last_updated_by_jira"`
	CreatedAt time.Time `json:"created_at"`
	Expired time.Time `json:"expired"`
}

func NewTicket() *Ticket {
	t := new(Ticket)
	t.StatusByJira = "Open"
	t.StatusClustered = TicketStatusClusteredOpen
	t.CreatedAt = time.Now()
	return t
}

func (t *Ticket) Expire() {
	t.Expired = time.Now()
}
