package process

import (
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/database"
	log "github.com/sirupsen/logrus"
	"github.com/rebel-l/jirastats/packages/models"
)

type Ticket struct {
	issue jira.Issue
	tm *database.TicketMapper
	IsNew bool
	StatusClustered string
}

func NewTicket(issue jira.Issue, tm *database.TicketMapper) *Ticket {
	t := new(Ticket)
	t.issue = issue
	t.tm = tm
	return t
}

func (t *Ticket) Process() {
	log.Infof("Process ticket: %s", t.issue.Key)
	newTicket := t.getNewTicket()

	// 1st find not expired old ticket and process if not changed
	oldTicket, err := t.tm.LoadByKey(t.issue.Key)
	if err != nil {
		log.Error("Couldn't load not expired old ticket: %s", err.Error())
	}

	if oldTicket.Id != 0 {
		log.Debugf("Unexpired old ticket (%s) found: %d", t.issue.Key, oldTicket.Id)
		if t.changed(newTicket, oldTicket) == false {
			log.Debugf("No changes found for ticket: %d (%s)", oldTicket.Id, oldTicket.Key)
			t.StatusClustered = oldTicket.StatusClustered
			t.setIsNew(oldTicket)
			return
		}

		oldTicket.Expire()
		err = t.tm.Save(oldTicket)
		if err != nil {
			log.Errorf("Old ticket could not be expired: %d (%s), error: %s", oldTicket.Id, oldTicket.Key, err.Error())
			return
		}
	} else {
		log.Debugf("No unexpired old ticket (%s) found", t.issue.Key)
	}

	// 2nd process new ticket
	// TODO: Cluster Status by Project information
	err = t.tm.Save(newTicket)
	if err != nil {
		log.Errorf("New ticket couldn't be saved: %s, error: %s", newTicket.Key, err.Error())
		return
	}

	t.StatusClustered = newTicket.StatusClustered
	t.setIsNew(newTicket)
}

func (t *Ticket) changed(newTicket *models.Ticket, oldTicket *models.Ticket) bool {
	// TODO implement diff new/old
	return false
}

func (t *Ticket) getNewTicket() *models.Ticket {
	newTicket := models.NewTicket()
	// TODO map jira issue to ticket
	return newTicket
}

func (t *Ticket) setIsNew(ticket *models.Ticket) {
	// TODO classify as (not) new and return this info
	t.IsNew = true
}
