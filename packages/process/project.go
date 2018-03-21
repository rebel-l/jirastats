package process

import (
	"database/sql"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/models"
	jp "github.com/rebel-l/jirastats/packages/jira"
	log "github.com/sirupsen/logrus"
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
}

func NewProject(project *models.Project, jc *jira.Client) *Project {
	p := new(Project)
	p.project = project
	p.jc = jc
	p.start = time.Now()
	return p
}

func (p *Project) Process() {
	log.Infof("Process project ... Id: %d, Name: %s", p.project.Id, p.project.Name)

	// TODO: 1st no stats
	p.initStats()

	// TODO 2nd update stats
	p.updateStats()


	t := time.Now()
	elapsed := t.Sub(p.start)
	log.Infof("Project (Id: %d, Name: %s) was successful processed in %f seconds", p.project.Id, p.project.Name, elapsed.Seconds())
	return
}

func (p *Project) initStats() (err error){
	return
}

func (p *Project) updateStats() (err error){
	search := jp.NewSearch(p.jc)
	tickets, err := search.Do(p.getJqlForUpdatedTickets())
	if err != nil {
		log.Errorf("Project (Id: %d, Name: %s) was not processed: %s", p.project.Id, p.project.Name, err.Error())
		return
	}

	// TODO: process in channels
	for _, t := range tickets {
		log.Debugf("Ticket (Updated): %s", t.Key)
	}
	return
}

func (p *Project) getJqlForUpdatedTickets() string {
	startDate := p.start.AddDate(0, 0, -1).Format(jp.JiraDateFormat)
	endDate := p.start.Format(jp.JiraDateFormat)

	jql := p.project.GetJql() + fmt.Sprintf(" AND updated >= %s AND updated < %s", startDate, endDate)
	log.Debugf("JQL for updated tickets: %s", jql)
	return jql
}


