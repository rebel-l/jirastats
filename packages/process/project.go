package process

import (
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/models"
	jp "github.com/rebel-l/jirastats/packages/jira"
	log "github.com/sirupsen/logrus"
	"time"
)

type Project struct {
	jc *jira.Client
	project *models.Project
	start time.Time
	//elapsed time.Time
	//executionDate time.Time
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

	// TODO 2nd update stats

	search := jp.NewSearch(p.jc)
	log.Debugf("JQL: %s", p.project.GetJql())
	tickets, err := search.Do(p.project.GetJql())
	if err != nil {
		log.Errorf("Project (Id: %d, Name: %s) was not processed: %s", p.project.Id, p.project.Name, err.Error())
		return
	}

	for _, t := range tickets {
		log.Debugf("Ticket: %s", t.Key)
	}

	t := time.Now()
	elapsed := t.Sub(p.start)
	log.Infof("Project (Id: %d, Name: %s) was successful processed in %f seconds", p.project.Id, p.project.Name, elapsed.Seconds())
	return
}
