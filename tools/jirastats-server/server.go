package main

import (
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
	log.SetLevel(log.DebugLevel)
	log.Infof("Run server on Port %d ...", PORT)

	// Init database
	db, err := database.GetDbConnection()
	defer db.Close()
	utils.HandleUnrecoverableError(err)

	// Init router
	router := mux.NewRouter()

	// Init Endpoints
	_ = endpoints.NewDataTickets(db, router)

	// start server
	err = http.ListenAndServe(":" + strconv.Itoa(PORT), router)
	utils.HandleUnrecoverableError(err)
}
