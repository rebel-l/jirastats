package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/utils"
	"github.com/rebel-l/jirastats/tools/jirastats-server/endpoints"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const PORT = 3000

func main() {
	verbose := utils.GetVerboseFlag()
	flag.Parse()

	// init log level
	utils.TurnOnVerbose(verbose)

	log.Infof("Run server on Port %d ...", PORT)

	// Init database
	db, err := database.GetDbConnection()
	defer db.Close()
	utils.HandleUnrecoverableError(err)

	// Init router
	router := mux.NewRouter()

	// Init Endpoints
	endpoints.NewDataTickets(db, router)
	endpoints.NewProjects(db, router)

	// Init File Server ... must be the last one
	endpoints.NewPublic(router)

	// start server
	err = http.ListenAndServe(":" + strconv.Itoa(PORT), router)
	utils.HandleUnrecoverableError(err)
}
