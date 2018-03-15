package main

import (
	"bufio"
	"database/sql"
	log "github.com/sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3"
	"github.com/andygrunwald/go-jira"
	jiraSearch "github.com/rebel-l/jirastats/packages/jira"
	"os"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
	"strings"
)

const StatementInsertTicket = "INSERT INTO ticket(`key`) values(?)"
const StatementSelectTickets = "SELECT * FROM ticket"

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Run collector ...")

	db, err := sql.Open("sqlite3", "./storage/jirastats.db")
	defer db.Close()
	handleUnrecoverableError(err)

	tickets := getTickets()
	storeTickets(tickets, db)
	readTickets(db)

	log.Info("Stopping collector ... Goodbye!")
}

func getTickets() []jira.Issue {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("")
	fmt.Print("Jira Username: ")
	username, _ := r.ReadString('\n')
	fmt.Print("Jira Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)
	fmt.Println("")
	fmt.Println("")

	// init Jira client
	jiraClient := initJiraClient(username, password)
	jiraSearch := jiraSearch.NewSearch(jiraClient)
	result, err := jiraSearch.Do("project = CORE")
	handleUnrecoverableError(err)
	return result
}

func storeTickets(tickets []jira.Issue, db *sql.DB) {
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

// TODO: should be part of central jira package
func initJiraClient(username string, password string) *jira.Client {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	jiraClient, _ := jira.NewClient(nil, "https://jira.home24.de")
	jiraClient.Authentication.SetBasicAuth(username, password)
	return jiraClient
}

func handleUnrecoverableError(err error) {
	if err != nil {
		log.Errorf("Unrecoverable error appeard: %s", err.Error())
		log.Panic("Collector failed ... Goodbye!")
	}
}
