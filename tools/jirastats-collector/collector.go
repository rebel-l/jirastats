package main

import (
	"database/sql"
	"errors"
	"flag"
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/models"
	"github.com/rebel-l/jirastats/packages/utils"
	jp "github.com/rebel-l/jirastats/packages/jira"
	log "github.com/sirupsen/logrus"
	"github.com/rebel-l/jirastats/packages/process"
	"time"
)


func main() {
	verbose := utils.GetVerboseFlag()
	flag.Parse()

	// init log level
	utils.TurnOnVerbose(verbose)

	log.Info("Run collector ...")

	db, err := database.GetDbConnection()
	defer db.Close()
	utils.HandleUnrecoverableError(err)

	c := NewCollector(db)
	jc := c.getJiraClient()
	projects, err := c.getProjects()
	if err != nil {
		utils.HandleUnrecoverableError(err)
	} else if len(projects) == 0 {
		utils.HandleUnrecoverableError(errors.New("No projects found"))
	}

	// TODO: use channels to parallize
	for _, p := range projects {
		pp := process.NewProject(p, jc, db)
		pp.Process()
	}

	// calculate elapsed time
	t := time.Now()
	elapsed := t.Sub(c.start)
	log.Infof("Processed %d projects in %f seconds", len(projects), elapsed.Seconds())
	log.Info("Stopping collector ... Goodbye!")
}

type Collector struct {
	db *sql.DB
	start time.Time
}

func NewCollector(db *sql.DB) *Collector {
	c := new(Collector)
	c.db = db
	c.start = time.Now()
	return c
}

func (c *Collector) getProjects() (projects []*models.Project, err error) {
	pm := database.NewProjectMapper(c.db)
	projects, err = pm.Load()
	return
}

func (c *Collector) getJiraClient() *jira.Client {
	config := jp.NewConfig(c.db)
	return jp.NewClient(config)
}
