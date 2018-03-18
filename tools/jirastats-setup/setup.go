package main

import (
	"flag"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/utils"
	log "github.com/sirupsen/logrus"
	"github.com/rebel-l/jirastats/tools/jirastats-setup/commands"
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
		log.Errorf("Stats couldn't be reseted: %s", err.Error())
	}
}

func createDatabaseFile() {
	log.Info("Create database storage")
	dbs := commands.CreateDatabaseStorage{}
	err := dbs.Execute()
	utils.HandleUnrecoverableError(err)
}

func createDatabaseStructure() {
	log.Info("Create database structure")
	db, err := database.GetDbConnection()
	defer db.Close()
	utils.HandleUnrecoverableError(err)

	dbs := commands.NewCreateDatabaseStructure(db)
	err = dbs.Execute()
	utils.HandleUnrecoverableError(err)
}
