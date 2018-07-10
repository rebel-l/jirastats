package main

import (
	"database/sql"
	"flag"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/utils"
	"github.com/rebel-l/jirastats/tools/jirastats-setup/commands"
	log "github.com/sirupsen/logrus"
)

func main() {
	// init flags
	verbose := utils.GetVerboseFlag()
	resetStats := flag.Bool("r", false, "To reset all stats. !Be careful, there is no recovery!")
	confProject := flag.Bool("p", false, "Creates a new project only")
	createDemo := flag.String("demo", "", "Creates demo data. You need to specify a value like 3d (3 days), 1w (1 week) or 2m (2 months)")
	flag.Parse()

	// init log level
	utils.TurnOnVerbose(verbose)


	// start setup
	log.Info("Run setup ...")

	if *resetStats == true {
		doResetStats()
	} else if *confProject == true {
		db, err := database.GetDbConnection()
		defer db.Close()
		utils.HandleUnrecoverableError(err)
		configureProjects(db)
	} else if *createDemo != "" {
		createDemoData(*createDemo)
	} else {
		createDatabaseFile()

		db, err := database.GetDbConnection()
		defer db.Close()
		utils.HandleUnrecoverableError(err)

		createDatabaseStructure(db)
		configureApplication(db)
		configureProjects(db)
		createDemoData("")
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

func configureProjects(db *sql.DB) {
	log.Info("Configure Projects")
	cp := commands.NewConfigureProjects(db)
	err := cp.Execute()
	utils.HandleUnrecoverableError(err)
}

func createDemoData(period string) {
	log.Info("Create Demo")
	cd := commands.NewCreateDemoData(period)
	err := cd.Execute()
	utils.HandleUnrecoverableError(err)
}
