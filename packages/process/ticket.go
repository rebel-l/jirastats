package process

import (
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/database"
	log "github.com/sirupsen/logrus"
)

type Ticket struct {
	issue jira.Issue
	tm *database.TicketMapper
}

func NewTicket(issue jira.Issue, tm *database.TicketMapper) *Ticket {
	t := new(Ticket)
	t.issue = issue
	t.tm = tm
	return t
}

func (t *Ticket) Process() {
	log.Infof("Process ticket: %s", t.issue.Key)
}
