package main

import (
	"database/sql"
	"errors"
	"flag"
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/models"
	"github.com/rebel-l/jirastats/packages/process"
	"github.com/rebel-l/jirastats/packages/utils"
	jp "github.com/rebel-l/jirastats/packages/jira"
	log "github.com/sirupsen/logrus"
	"time"
	"sync"
)

const defaultConcurrentProjects = 5
const defaultInterval = 1

func main() {
	verbose := utils.GetVerboseFlag()
	interval := flag.Int("i", defaultInterval, "Interval of days included, e.g. 3 means update stats starts 3 days before now")
	concurrentProjects := flag.Int("p", defaultConcurrentProjects, "Number of projects to execute concurrent")
	flag.Parse()

	// init log level
	utils.TurnOnVerbose(verbose)

	log.Info("Run collector ...")

	db, err := database.GetDbConnection()
	defer db.Close()
	utils.HandleUnrecoverableError(err)

	c := NewCollector(db, *interval, *concurrentProjects)
	projects, err := c.getProjects()
	if err != nil {
		utils.HandleUnrecoverableError(err)
	} else if len(projects) == 0 {
		utils.HandleUnrecoverableError(errors.New("No projects found"))
	}

	c.distributeProjects(projects)
	c.executeConcurrent()

	// calculate elapsed time
	t := time.Now()
	elapsed := t.Sub(c.start)
	log.Infof("Processed %d projects in %s", len(projects), elapsed.String())
	log.Info("Stopping collector ... Goodbye!")
}

type Collector struct {
	db *sql.DB
	jc *jira.Client
	start time.Time
	interval int
	concurrentProjects int
	distributedProjects [][]*models.Project
	wg sync.WaitGroup
}

func NewCollector(db *sql.DB, interval int, concurrentProjects int) *Collector {
	c := new(Collector)
	c.start = time.Now()
	c.db = db
	c.jc = c.getJiraClient()
	c.interval = interval
	c.concurrentProjects = concurrentProjects
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

func (c *Collector) distributeProjects(projects []*models.Project) {
	c.distributedProjects = make([][]*models.Project, c.concurrentProjects)
	i := 0
	for _, p := range projects {
		c.distributedProjects[i] = append(c.distributedProjects[i], p)
		i++
		if i == c.concurrentProjects {
			i = 0
		}
	}
}

func (c *Collector) executeConcurrent() {
	for i := 0; i < c.concurrentProjects; i++ {
		c.wg.Add(1)
		go c.executeSetOfProjects(c.distributedProjects[i])
	}
	c.wg.Wait()
}

func (c *Collector) executeSetOfProjects(projects []*models.Project) {
	defer c.wg.Done()
	for _, p := range projects {
		pp := process.NewProject(p, c.jc, c.db, c.interval)
		pp.Process()
	}
}
