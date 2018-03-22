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

	// 1st find not expired ticket
	ticket, err := t.tm.LoadByKey(t.issue.Key)
	if err != nil {
		log.Error("Couldn't load not expired ticket: %s", err.Error())
	}

	if ticket.Id != 0 {
		// TODO: check for changes, if changed close old one
		log.Debugf("Unexpired ticket (%s) found: %d", t.issue.Key, ticket.Id)
	} else {
		log.Debugf("No unexpired ticket (%s) found", t.issue.Key)
	}
	// TODO insert new one

	// TODO: classify new, open, closed and return this info
}
