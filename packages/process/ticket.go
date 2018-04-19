package process

import (
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/models"
	"github.com/rebel-l/jirastats/packages/utils"
	jp "github.com/rebel-l/jirastats/packages/jira"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strings"
	"time"
)

type Ticket struct {
	projectId int
	issue jira.Issue
	tm *database.TicketMapper
	IsNew bool
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

	// 1st find previous ticket and process if changed
	prevTicket, err := t.tm.LoadLastByKey(t.issue.Key, t.projectId)
	if err != nil {
		log.Errorf("Project %d: Couldn't load not expired old ticket: %s",t.projectId, err.Error())
	}

	if prevTicket.Id != 0 {
		log.Debugf("Project %d: old ticket (%s) found: %d", t.projectId, t.issue.Key, prevTicket.Id)
		if t.changed(newTicket, prevTicket) == false {
			log.Debugf("Project %d: No changes found for ticket: %d (%s)", t.projectId, prevTicket.Id, prevTicket.Key)
			return
		}

		log.Debugf("Project %d: Changes found for ticket: %d (%s)", t.projectId, prevTicket.Id, prevTicket.Key)
		if prevTicket.Expired.IsZero() {
			// not expired
			prevTicket.ExpireEndOfDayBefore() // needs to be actual run at 23:59:59
			err = t.tm.Save(prevTicket)
			if err != nil {
				log.Errorf("Project %d: Old ticket could not be expired: %d (%s), error: %s", t.projectId, prevTicket.Id, prevTicket.Key, err.Error())
				return
			}
			log.Debugf("Project %d: Ticket expired: %d (%s)", t.projectId, prevTicket.Id, prevTicket.Key)
		} else {
			// expired ... only new if it was removed for some reason and appears again
			//t.IsNew = true // TODO: only if old one was removed
			//newTicket.IsNew = true
		}
	} else {
		// appeared the first time
		t.IsNew = true
		newTicket.IsNew = true
		log.Debugf("Project %d: No old ticket (%s) found", t.projectId, t.issue.Key)
	}

	// 2nd process new ticket
	for i := 0; i < retryMax; i++ {
		err = t.tm.Save(newTicket)
		if err == nil {
			break
		}

		if strings.Contains(err.Error(), "locked") == false || i == retryMax- 1 {
			log.Errorf("Project %d: New ticket couldn't be saved: %s, error: %s", t.projectId, newTicket.Key, err.Error())
			return
		}

		wt := time.Duration(rand.Intn(retryWait))
		time.Sleep(wt * time.Millisecond)
	}
}

func (t *Ticket) changed(newTicket *models.Ticket, oldTicket *models.Ticket) bool {
	switch {
		case newTicket.Issuetype != oldTicket.Issuetype:
			return true
		case newTicket.Priority != oldTicket.Priority:
			return true
		case newTicket.StatusClustered != oldTicket.StatusClustered:
			return true
		case newTicket.StatusByJira != oldTicket.StatusByJira:
			return true
		case newTicket.Summary != oldTicket.Summary:
			return true
		case utils.AreStringArrayEqual(newTicket.Components, oldTicket.Components) == false:
			return true
		case utils.AreStringArrayEqual(newTicket.Labels, oldTicket.Labels) == false:
			return true
	}
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
			newTicket.SetStatusClustered(clusteredStatus)
			break
		}
	}

	return newTicket
}
