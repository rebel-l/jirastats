package main

import (
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/utils"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"os"
)

const DEFAULT_DB_PATH = "./storage/jirastats.db"
const SQL_SETUP_SCRIPT = "./setup/setup.sql"

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Run setup ...")

	createDatabaseFile()
	createDatabaseStructure()

	log.Info("Setup finished successful ... Goodbye!")
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
