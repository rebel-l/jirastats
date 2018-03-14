package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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

func handleUnrecoverableError(err error) {
	if err != nil {
		log.Errorf("Unrecoverable error appeard: %s", err.Error())
		log.Panic("Setup failed ... Goodbye!")
	}
}

func createDatabaseFile() {
	_, err := os.Stat(DEFAULT_DB_PATH)
	if err != nil {
		log.Warnf("Database file '%s' does not exist and will be created", DEFAULT_DB_PATH)
		_, err = os.Create(DEFAULT_DB_PATH)
		handleUnrecoverableError(err)
	}
}

func createDatabaseStructure() {
	statements, err := ioutil.ReadFile(SQL_SETUP_SCRIPT)
	handleUnrecoverableError(err)

	db, err := sql.Open("sqlite3", "./storage/jirastats.db")
	defer db.Close()
	handleUnrecoverableError(err)

	stmt, err := db.Prepare(string(statements))
	handleUnrecoverableError(err)

	_, err = stmt.Exec()
	handleUnrecoverableError(err)
}
