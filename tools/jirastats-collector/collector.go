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
)

const StatementInsertTicket = "INSERT INTO ticket(`key`) values(?)" // TODO: deprecated
const StatementSelectTickets = "SELECT * FROM ticket" // TODO: deprecated


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
		log.Infof("Process project ... Id: %d, Name: %s", p.Id, p.Name)
		search := jp.NewSearch(jc)
		log.Debugf("JQL: %s", p.GetJql())
		tickets, err := search.Do(p.GetJql())
		if err != nil {
			log.Errorf("Project (Id: %d, Name: %s) was not processed: %s", p.Id, p.Name, err.Error())
			continue
		}

		for _, t := range tickets {
			log.Debugf("Ticket: %s", t.Key)
		}
	}


	//
	//tickets := getTickets()
	//storeTickets(tickets, db)
	//readTickets(db)

	log.Info("Stopping collector ... Goodbye!")
}

type Collector struct {
	db *sql.DB
}

func NewCollector(db *sql.DB) *Collector {
	c := new(Collector)
	c.db = db
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



func storeTickets(tickets []jira.Issue, db *sql.DB) {
	// TODO: deprecated
	stmt, err := db.Prepare(StatementInsertTicket)
	handleUnrecoverableError(err)

	for _, v := range tickets {
		_, err = stmt.Exec(v.Key)
		if err != nil {
			log.Warnf("Couldn't save ticket %s: %s", v.Key, err.Error())
		}
	}
}

func readTickets(db *sql.DB) {
	// TODO: deprecated
	stmt, err := db.Prepare(StatementSelectTickets)
	handleUnrecoverableError(err)
	rows, err := stmt.Query()
	if err != nil {
		log.Warn("Couldn't read tickets from database")
		return
	}

	var id int
	var key string

	for rows.Next() {
		err = rows.Scan(&id, &key)
		log.Infof("TicketId: %d with Key: %s", id, key)
	}
}

func handleUnrecoverableError(err error) {
	// TODO: deprecated
	if err != nil {
		log.Errorf("Unrecoverable error appeard: %s", err.Error())
		log.Panic("Collector failed ... Goodbye!")
	}
}
