package main

import (
	"flag"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/utils"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/rebel-l/jirastats/tools/jirastats-setup/commands"
)

const DEFAULT_DB_PATH = "./storage/jirastats.db"
const SQL_SETUP_SCRIPT = "./setup/setup.sql"

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Run setup ...")

	resetStats := flag.Bool("resetstats", false, "To reset all stats. !Be careful, there is no recovery!")
	flag.Parse()

	if *resetStats == true {
		doResetStats()
	} else {
		createDatabaseFile()
		createDatabaseStructure()
	}

	log.Info("Setup finished successful ... Goodbye!")
}

func doResetStats() {
	db, err := database.GetDbConnection()
	defer db.Close()
	utils.HandleUnrecoverableError(err)

	rsc := commands.NewResetStats(db)
	err = rsc.Execute()
	if err != nil {
		log.Errorf("Stats couldn't be resetted: %s", err.Error())
	}
}

func createDatabaseFile() {
	_, err := os.Stat(DEFAULT_DB_PATH)
	if err != nil {
		log.Warnf("Database file '%s' does not exist and will be created", DEFAULT_DB_PATH)
		_, err = os.Create(DEFAULT_DB_PATH)
		utils.HandleUnrecoverableError(err)
	}
}

func createDatabaseStructure() {
	statements, err := ioutil.ReadFile(SQL_SETUP_SCRIPT)
	utils.HandleUnrecoverableError(err)

	db, err := database.GetDbConnection()
	defer db.Close()
	utils.HandleUnrecoverableError(err)

	stmt, err := db.Prepare(string(statements))
	utils.HandleUnrecoverableError(err)

	_, err = stmt.Exec()
	utils.HandleUnrecoverableError(err)
}
