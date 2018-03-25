package process

import (
	"database/sql"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/models"
	"github.com/rebel-l/jirastats/packages/utils"
	jp "github.com/rebel-l/jirastats/packages/jira"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type Project struct {
	db *sql.DB
	jc *jira.Client
	project *models.Project
	start time.Time
	actualRun time.Time
	stats *models.Stats
	pm *database.ProjectMapper
	sm *database.StatsMapper
}

func NewProject(project *models.Project, jc *jira.Client, db *sql.DB) *Project {
	p := new(Project)
	p.start = time.Now()
	p.actualRun = p.start.AddDate(0, 0, -1) // TODO: must be -1, just used -2 for testing
	p.project = project
	p.jc = jc
	p.pm = database.NewProjectMapper(db)
	p.sm = database.NewStatsMapper(db)
	p.db = db
	return p
}

func (p *Project) Process() {
	log.Infof("Process project ... Id: %d, Name: %s", p.project.Id, p.project.Name)

	// 1st init tickets if there are none
	if p.pm.HasTickets(p.project) == false {
		p.initStats()
	}

	// 2nd update stats
	p.updateStats()

	t := time.Now()
	elapsed := t.Sub(p.start)
	log.Infof("Project (Id: %d, Name: %s) was successful processed in %f seconds", p.project.Id, p.project.Name, elapsed.Seconds())
	return
}

func (p *Project) initStats() (err error) {
	log.Debugf("Init project stats: %d (%s)", p.project.Id, p.project.Name)
	p.stats = models.NewStats(p.project.Id)
	p.stats.CreatedAt = p.actualRun.AddDate(0, 0, -1) // Initial stats needs to be saved two days ago
	search := jp.NewSearch(p.jc, p.getJqlForOpenTickets())
	err = p.processTickets(search)
	err = p.processStats()
	return
}

func (p *Project) updateStats() (err error){
	log.Debugf("Update project stats: %d (%s)", p.project.Id, p.project.Name)
	p.stats = models.NewStats(p.project.Id)
	p.stats.CreatedAt = p.actualRun // Updated stats needs to be saved 1 day ago
	search := jp.NewSearch(p.jc, p.getJqlForUpdatedTickets())
	err = p.processTickets(search)
	if err != nil {
		return
	}

	err = p.processRemoved()
	if err != nil {
		return
	}

	err = p.processStats()
	return
}

func (p *Project) processTickets(search *jp.Search) (err error) {
	i := 0
	for {
		tickets, err := search.Do()
		if err != nil {
			log.Errorf("Project (Id: %d, Name: %s) was not processed: %s", p.project.Id, p.project.Name, err.Error())
			return err
		}

		// TODO: process in channels
		for _, t := range tickets {
			i++
			tm := database.NewTicketMapper(p.db)
			mapOpenStatus := utils.TrimMap(strings.Split(p.project.MapOpenStatus, ","))
			mapClosedStatus := utils.TrimMap(strings.Split(p.project.MapClosedStatus, ","))
			tp := NewTicket(p.project.Id, t, tm, mapOpenStatus, mapClosedStatus)
			tp.Process()
			if tp.IsNew {
				p.stats.New++
			}
		}

		if search.Next() == false {
			break
		}
	}

	log.Debugf("Tickets for project %d (%s) processed: %d", p.project.Id, p.project.Name, i)
	return
}

func (p *Project) processStats() (err error) {
	tm := database.NewTicketMapper(p.db)

	// open = expired null && status_clustered open
	p.stats.Open, err = tm.CountStatusClusteredAndNotExpired(models.TicketStatusClusteredOpen)
	if err != nil {
		return
	}

	// closed = expired today && status closed and is the first closed entry
	p.stats.Closed, err = tm.CountStatusClusteredFromDay(models.TicketStatusClusteredClosed, time.Now())
	if err != nil {
		return
	}

	err = p.sm.Save(p.stats)
	log.Infof("Stats processed: %d open, %d closed, %d new", p.stats.Open, p.stats.Closed, p.stats.New)
	return
}

func (p *Project) processRemoved() (err error) {
	// how to figure out the tickets removed from result of jql? ==> idea: not expired is not existing in overall jql
	tm := database.NewTicketMapper(p.db)
	tickets, err := tm.LoadNotExpired(p.project.Id)
	if err != nil {
		return
	}

	log.Debugf("Found %d not expired tickets", len(tickets))


	search := jp.NewSearch(p.jc, p.project.GetJql())
	search.Request.MaxResults = 20
	search.Request.Fields = make([]string, 0)
	i := 0
	for {
		res, err := search.Do()
		if err != nil {
			log.Errorf("Project (Id: %d, Name: %s) removed were not processed: %s", p.project.Id, p.project.Name, err.Error())
			return err
		}

		// TODO: parallelize in channels
		i += p.markTicketsToKeep(res, tickets)
		if search.Next() == false {
			break
		}
	}

	j, err := p.expireRemoved(tickets, tm)
	if err != nil {
		log.Errorf("Removed tickets couldn't be expired: %s", err.Error())
	}

	log.Debugf("Removed processed: %d, Kept: %d/%d", i, j, len(tickets))
	return
}

func (p *Project) markTicketsToKeep(jiraTickets []jira.Issue, tickets []*models.Ticket) (processed int) {
	for _, jt := range jiraTickets {
		processed++
		for _, t := range tickets {
			if jt.Key == t.Key {
				// mark ticket to keep
				t.Id = 0
			}
		}
	}
	return
}

func (p *Project) expireRemoved(tickets []*models.Ticket, tm *database.TicketMapper) (kept int, err error) {
	for _, t := range tickets {
		// expire all which still has an id > 0
		if t.Id == 0 {
			kept++
			continue
		}

		t.Expire()
		err = tm.Save(t)
		return
	}
	return
}



func (p *Project) getJqlForUpdatedTickets() string {
	startDate := p.actualRun.Format(jp.JiraJqlDateFormat)
	endDate := p.start.Format(jp.JiraJqlDateFormat)

	jql := p.project.GetJql() + fmt.Sprintf(" AND updated >= %s AND updated < %s", startDate, endDate)
	log.Debugf("JQL for updated tickets: %s", jql)
	return jql
}

func (p *Project) getOpenStatusMapForJql() string {
	status := strings.Split(p.project.MapOpenStatus, ",")
	for k, v := range status {
		v = strings.TrimSpace(v)
		if strings.Contains(v, "\"") == false {
			status[k] = "\"" + v + "\""
		}
	}

	return strings.Join(status, ",")
}

func (p *Project) getJqlForOpenTickets() string {
	jql := p.project.GetJql() + fmt.Sprintf(" AND status in (%s)", p.getOpenStatusMapForJql())
	log.Debugf("JQL for open tickets: %s", jql)
	return jql
}
