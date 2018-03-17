package endpoints

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rebel-l/jirastats/packages/database"
	"net/http"
	log "github.com/sirupsen/logrus"
	"encoding/json"
)

type DataTickets struct {
	db *sql.DB
}

func NewDataTickets(db *sql.DB, router *mux.Router) *DataTickets {
	dt := new(DataTickets)
	dt.db = db
	router.HandleFunc("/data/tickets", dt.Handler).Methods(http.MethodGet)

	return dt
}

func (dt *DataTickets) Handler(res http.ResponseWriter, req *http.Request) {
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
