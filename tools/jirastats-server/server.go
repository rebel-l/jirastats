package main

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/utils"
	log "github.com/sirupsen/logrus"
)

const PORT = 3000

func main() {
	log.SetLevel(log.DebugLevel)
	log.Infof("Run server on Port %d ...", PORT)

	db, err := sql.Open(database.DefaultDriver, database.DefaultFile)
	defer db.Close()
	utils.HandleUnrecoverableError(err)

	tm := database.NewTicketMapper(db)
	tickets, err := tm.Load()
	utils.HandleUnrecoverableError(err)
	log.Debugf("Number of results: %d", len(tickets))
	for _, v := range tickets {
		log.Infof("TicketId: %d with Key: %s", v.Id, v.Key)
	}
}


