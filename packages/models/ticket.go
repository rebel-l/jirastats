package models

import (
	"time"
)

const TicketStatusClusteredOpen = "Open"
const TicketStatusClusteredClosed = "Closed"
const dateFormat = "2006-01-02"
const dateTimeFormat = "2006-01-02T15:04:05Z"

type Ticket struct {
	Id                int       `json:"id"`
	Key               string    `json:"key"`
	ProjectId         int       `json:"project_id"`
	Summary           string    `json:"summary"`
	Components        []string  `json:"components"`
	Labels            []string  `json:"labels"`
	IsNew             bool      `json:"is_new"`
	StatusByJira      string    `json:"status_by_jira"`
	StatusClustered   string    `json:"status_clustered"`
	Priority          string    `json:"priority"`
	Issuetype         string    `json:"issuetype"`
	Removed           bool      `json:"removed"`
	CreatedAtByJira   time.Time `json:"created_at_by_jira"`
	LastUpdatedByJira time.Time `json:"last_updated_by_jira"`
	CreatedAt         time.Time `json:"created_at"`
	Expired           time.Time `json:"expired"`
}

func NewTicket() *Ticket {
	t := new(Ticket)
	t.StatusByJira = "Open"
	t.StatusClustered = TicketStatusClusteredOpen
	t.IsNew = false
	t.Removed = false
	t.SetCreatedAt(time.Now())
	return t
}

func (t *Ticket) ExpireNow(actualRun time.Time) {
	t.Expired, _ = time.Parse(dateTimeFormat,actualRun.Format(dateFormat)+"T23:59:59Z")
}

func (t *Ticket) ExpireEndOfDayBefore(actualRun time.Time) {
	t.Expired = actualRun
	t.Expired.AddDate(0, 0, -1)
	t.Expired, _ = time.Parse(dateTimeFormat, t.Expired.Format(dateFormat)+"T23:59:59Z")
}

func (t *Ticket) SetStatusClustered(status string, actualRun time.Time) {
	t.StatusClustered = status
	if t.StatusClustered == TicketStatusClusteredClosed && t.Expired.IsZero() {
		t.ExpireNow(actualRun)
	}
}

func (t *Ticket) SetCreatedAt(actualRun time.Time) {
	t.CreatedAt, _ = time.Parse(dateTimeFormat, actualRun.Format(dateFormat)+"T00:00:00Z")
}
