package process

import (
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/models"
	"github.com/rebel-l/jirastats/packages/utils"
	jp "github.com/rebel-l/jirastats/packages/jira"
	log "github.com/sirupsen/logrus"
	"time"
)

type Ticket struct {
	projectId int
	issue jira.Issue
	tm *database.TicketMapper
	IsNew bool
	StatusClustered string
	statusMap map[string][]string
}

func NewTicket(projectId int, issue jira.Issue, tm *database.TicketMapper, openStatusMap []string, closedStatus []string) *Ticket {
	t := new(Ticket)
	t.projectId = projectId
	t.issue = issue
	t.tm = tm
	t.statusMap = make(map[string][]string, 2)
	t.statusMap[models.TicketStatusClusteredOpen] = openStatusMap
	t.statusMap[models.TicketStatusClusteredClosed] = closedStatus
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
			t.processStats(oldTicket)
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
	err = t.tm.Save(newTicket)
	if err != nil {
		log.Errorf("New ticket couldn't be saved: %s, error: %s", newTicket.Key, err.Error())
		return
	}

	t.processStats(newTicket)
}

func (t *Ticket) changed(newTicket *models.Ticket, oldTicket *models.Ticket) bool {
	// TODO implement diff new/old
	return false
}

func (t *Ticket) getNewTicket() *models.Ticket {
	newTicket := models.NewTicket()
	newTicket.CreatedAtByJira, _ = time.Parse(jp.JiraDateTimeFormat, t.issue.Fields.Created)
	newTicket.Issuetype = t.issue.Fields.Type.Name
	newTicket.Key = t.issue.Key
	newTicket.Labels = t.issue.Fields.Labels
	newTicket.LastUpdatedByJira, _ = time.Parse(jp.JiraDateTimeFormat, t.issue.Fields.Updated)
	newTicket.Priority = t.issue.Fields.Priority.Name
	newTicket.ProjectId = t.projectId
	newTicket.Summary = t.issue.Fields.Summary
	newTicket.StatusByJira = t.issue.Fields.Status.Name

	for _, component := range t.issue.Fields.Components {
		newTicket.Components = append(newTicket.Components, component.Name)
	}

	for clusteredStatus, statusMap := range t.statusMap {
		if utils.IsValueInMap(statusMap, newTicket.StatusByJira) {
			newTicket.StatusClustered = clusteredStatus
			break
		}
	}

	return newTicket
}

func (t *Ticket) setIsNew(ticket *models.Ticket) {
	// TODO classify as (not) new and return this info
	t.IsNew = true
}

func (t *Ticket) processStats(ticket *models.Ticket) {
	t.StatusClustered = ticket.StatusClustered
	t.setIsNew(ticket)
}
