package process

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/models"
	jp "github.com/rebel-l/jirastats/packages/jira"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type Project struct {
	search *jp.Search
	project *models.Project
	start time.Time
	counterOpen int
	counterNew int
	counterClosed int
	pm *database.ProjectMapper
}

func NewProject(project *models.Project, jc *jira.Client, pm *database.ProjectMapper) *Project {
	p := new(Project)
	p.start = time.Now()
	p.project = project
	p.search = jp.NewSearch(jc)
	p.pm = pm
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
	tickets, err := p.search.Do(p.getJqlForOpenTieckets())
	if err != nil {
		log.Errorf("Project (Id: %d, Name: %s) was not processed: %s", p.project.Id, p.project.Name, err.Error())
		return
	}

	// TODO: process in channels
	for _, t := range tickets {
		log.Debugf("Ticket (Open): %s", t.Key)
	}
	return
}

func (p *Project) updateStats() (err error){
	tickets, err := p.search.Do(p.getJqlForUpdatedTickets())
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

func (p *Project) getJqlForOpenTieckets() string {
	jql := p.project.GetJql() + fmt.Sprintf(" AND status in (%s)", p.getOpenStatusMapForJql())
	log.Debugf("JQL for open tickets: %s", jql)
	return jql
}
