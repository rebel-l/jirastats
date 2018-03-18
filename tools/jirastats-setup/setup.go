package main

import (
	"flag"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/utils"
	log "github.com/sirupsen/logrus"
	"github.com/rebel-l/jirastats/tools/jirastats-setup/commands"
	"database/sql"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Run setup ...")

	resetStats := flag.Bool("resetstats", false, "To reset all stats. !Be careful, there is no recovery!")
	flag.Parse()

	if *resetStats == true {
		doResetStats()
	} else {
		createDatabaseFile()

		db, err := database.GetDbConnection()
		defer db.Close()
		utils.HandleUnrecoverableError(err)

		createDatabaseStructure(db)
		configureApplication(db)
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
		log.Errorf("Stats couldn't be reseted: %s", err.Error())
	}
}

func createDatabaseFile() {
	log.Info("Create database storage")
	dbs := commands.CreateDatabaseStorage{}
	err := dbs.Execute()
	utils.HandleUnrecoverableError(err)
}

func createDatabaseStructure(db *sql.DB) {
	log.Info("Create database structure")
	dbs := commands.NewCreateDatabaseStructure(db)
	err := dbs.Execute()
	utils.HandleUnrecoverableError(err)
}

func configureApplication(db *sql.DB) {
	log.Info("Configure Application")
	ca := commands.NewConfigureApplication(db)
	err := ca.Execute()
	utils.HandleUnrecoverableError(err)
}