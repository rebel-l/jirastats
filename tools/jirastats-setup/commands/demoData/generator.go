package demoData

import (
	"database/sql"
	"fmt"
	"time"
	j "github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/jira"
	"github.com/rebel-l/jirastats/packages/process"
	"github.com/rebel-l/jirastats/packages/models"
	"github.com/rebel-l/jirastats/packages/utils/random"
	log "github.com/sirupsen/logrus"
)

var components = [5]string{"OrderService", "CustomerService", "ProductService", "StockService", "WishlistService"}
var labels = [4]string{"TechDebt", "Budget", "OKR", "QA"}
var priorities = [5]string{"Trivial", "Minor", "Major", "Critical", "Blocker"}
var status = [5]string{"Open", "In Progress", "Reopened", "Closed", "Resolved"}
var finalStatus = [2]string{"Closed", "Resolved"}
var types = [5]string{"SubTask", "Task", "Story", "Bug", "Improvement"}

type Generator struct {
	db *sql.DB
	project *models.Project
	historicIssues []j.Issue
	ticketCounter int
}

func NewGenerator(db *sql.DB) *Generator {
	g := new(Generator)
	g.db = db
	g.historicIssues = make([]j.Issue, 0)
	return g
}

func (g *Generator) CreateProject(name string, keys string) error {
	p := models.NewProject()
	p.Name = name
	p.Keys = keys
	pm := database.NewProjectMapper(g.db)
	err := pm.Save(p)
	if err != nil {
		return err
	}
	g.project = p
	return nil
}

func (g *Generator) GenerateDay(actualDate time.Time) {
	diff := time.Now().Sub(actualDate)
	interval := diff.Hours() / 24.0
	log.Debugf("Actual date: %s, Now: %s, Duration: %s, Days: %d", actualDate.Format("02.01.2006"), time.Now().Format("02.01.2006"), diff.String(), int(interval))

	cs := jira.NewClientStub()
	log.Debugf("Iterval: %d", int(interval))
	pp := process.NewProject(g.project, cs, g.db, int(interval))

	max := random.Int(1, 10) // TODO: must be possible to have also 0 here but not for project init
	for i := 0; i < max; i++ {
		issue := g.newIssue(actualDate)
		
		if len(g.historicIssues) == 0 {
			if g.isInFinalState(issue.Fields.Status.Name) {
				cs.AddIssue(issue, pp.GetJqlForClosedTickets())
			} else {
				cs.AddIssue(issue, pp.GetJqlForOpenTickets())
			}
		} else {
			cs.AddIssue(issue, pp.GetJqlForUpdatedTickets()) // avoid to create new tickets in final state
		}
		
		log.Debugf("New Issue: %#v", issue)
		g.historicIssues = append(g.historicIssues, issue)
	}

	pp.Process()
}

func (g *Generator) newIssue(actualDate time.Time) j.Issue {
	key := g.newKey()

	created := random.DateTimeBefore(actualDate, 0, 100)
	updated := random.TimeBefore(actualDate)

	jPriority := new(j.Priority)
	jPriority.Name = g.getPriority()

	jStatus := new(j.Status)
	jStatus.Name = g.getStatus()

	issueFields := new(j.IssueFields)
	issueFields.Summary = "Summary"
	issueFields.Type = j.IssueType{Name: g.getType()}
	issueFields.Status = jStatus
	issueFields.Priority = jPriority
	issueFields.Components = g.getComponents()
	issueFields.Labels = g.getLabels()
	issueFields.Created = created.Format(jira.JiraDateTimeFormat)
	issueFields.Updated = updated.Format(jira.JiraDateTimeFormat)

	return j.Issue{ID: key,	Key: key, Fields: issueFields}
}

func (g *Generator) newKey() string {
	g.ticketCounter++
	keys := g.project.GetKeys()
	return fmt.Sprintf("%s-%d", keys[random.Int(0, len(keys))], g.ticketCounter)
}

func (g *Generator) getPriority() string {
	r := random.Int(1, 100)
	i := 0
	switch {
	case r > 5 && r <= 15:
		i = 1
		break
	case r > 15 && r <= 85:
		i = 2;
		break
	case r > 85 && r <= 95:
		i = 3
		break
	case r > 95 && r <= 100:
		i = 4
		break
	}
	return priorities[i]
}

func (g *Generator) getType() string {
	r := random.Int(1, 100)
	i := 0
	switch {
	case r > 5 && r <= 15:
		i = 1
		break
	case r > 15 && r <= 85:
		i = 2;
		break
	case r > 85 && r <= 95:
		i = 3
		break
	case r > 95 && r <= 100:
		i = 4
		break
	}
	return types[i]
}

func (g *Generator) getLabels() []string {
	l := make([]string, 0)
	r := random.Int(1, 100)
	max := 0
	switch {
	case r > 50 && r <= 80:
		max = 1
		break
	case r > 80 && r <= 95:
		max = 2;
		break
	case r > 95 && r <= 100:
		max = 3
		break
	}

	pool := labels[0:len(labels)]
	for i := 0; i < max; i++ {
		r := random.Int(0, len(pool))
		l = append(l, pool[r])

		left := pool[0:r]
		right := pool[r:len(pool)]
		pool = left
		pool = append(pool, right...)
	} 

	return l
}

func (g *Generator) getStatus() string {
	return status[random.Int(0, len(status) - 1)]
}

func (g *Generator) getComponents() []*j.Component {
	c := make([]*j.Component, 0)
	r := random.Int(1, 100)
	max := 0
	switch {
	case r > 10 && r <= 85:
		max = 1
		break
	case r > 85 && r <= 100:
		max = 2;
		break
	}

	pool := components[0:len(components)]
	for i := 0; i < max; i++ {
		r := random.Int(0, len(pool))

		nc := new(j.Component)
		nc.Name = pool[r]
		c = append(c, nc)
		
		left := pool[0:r]
		right := pool[r:len(pool)]
		pool = left
		pool = append(pool, right...)
	} 

	return c
}

func (g *Generator) isInFinalState(status string) bool {
	for i := range finalStatus {
		if finalStatus[i] == status {
			return true
		}
	}

	return false
}
