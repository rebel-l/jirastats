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
	counterOpen int
	counterNew int
	counterClosed int
	pm *database.ProjectMapper
}

func NewProject(project *models.Project, jc *jira.Client, db *sql.DB) *Project {
	p := new(Project)
	p.start = time.Now()
	p.project = project
	p.jc = jc
	p.pm = database.NewProjectMapper(db)
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
	search := jp.NewSearch(p.jc, p.getJqlForOpenTickets())
	err = p.processTickets(search)
	// TODO: stats save 2 days ago ... and EVERY ticket is NEW
	return
}

func (p *Project) updateStats() (err error){
	log.Debugf("Update project stats: %d (%s)", p.project.Id, p.project.Name)
	search := jp.NewSearch(p.jc, p.getJqlForUpdatedTickets())
	err = p.processTickets(search)

	// TODO: how to figure out the tickets removed from result of jql?

	// TODO: stats save 1 days ago
	return
}

func (p *Project) processTickets(search *jp.Search) (err error) {
	for {
		tickets, err := search.Do()
		if err != nil {
			log.Errorf("Project (Id: %d, Name: %s) was not processed: %s", p.project.Id, p.project.Name, err.Error())
			return err
		}

		// TODO: process in channels
		for _, t := range tickets {
			tm := database.NewTicketMapper(p.db)
			mapOpenStatus := utils.TrimMap(strings.Split(p.project.MapOpenStatus, ","))
			mapClosedStatus := utils.TrimMap(strings.Split(p.project.MapClosedStatus, ","))
			tp := NewTicket(p.project.Id, t, tm, mapOpenStatus, mapClosedStatus)
			tp.Process()
		}

		if search.Next() == false {
			break
		}
	}
	return
}

func (p *Project) getJqlForUpdatedTickets() string {
	startDate := p.start.AddDate(0, 0, -1).Format(jp.JiraJqlDateFormat)
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
