package endpoints

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/tools/jirastats-server/response"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const getTicketsPath = "/data/tickets"

type DataTickets struct {
	tm *database.TicketMapper
}

func NewDataTickets(db *sql.DB, router *mux.Router) {
	dt := new(DataTickets)
	dt.tm = database.NewTicketMapper(db)
	router.HandleFunc(getTicketsPath, dt.GetTickets).Methods(http.MethodGet)
}

func (dt *DataTickets) GetTickets(res http.ResponseWriter, req *http.Request) {
	log.Debug("Get all data from tickets")

	tickets, err := dt.tm.Load()
	if err != nil {
		msg := fmt.Sprintf("Not able to load tickets: %s", err.Error())
		e := response.NewErrorJson(msg, res)
		e.SendBadRequest()
		return
	}

	success := response.NewSuccessJson(tickets, res)
	success.SendOK()
}
