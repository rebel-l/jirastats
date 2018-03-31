package endpoints

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rebel-l/jirastats/packages/database"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type DataTickets struct {
	db *sql.DB
}

func NewDataTickets(db *sql.DB, router *mux.Router) {
	dt := new(DataTickets)
	dt.db = db
	router.HandleFunc("/data/tickets", dt.Handler).Methods(http.MethodGet)
}

func (dt *DataTickets) Handler(res http.ResponseWriter, req *http.Request) {
	log.Debug("Get all data from tickets")
	status := http.StatusOK

	tm := database.NewTicketMapper(dt.db)
	tickets, err := tm.Load()
	if err != nil {
		log.Errorf("Not able to load tickets: %s", err.Error())
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	err = json.NewEncoder(res).Encode(tickets)
	if err != nil {
		log.Errorf("Not able to convert tickets to json: %s", err.Error())
		return
	}
}
