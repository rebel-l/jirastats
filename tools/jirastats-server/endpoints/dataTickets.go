package endpoints

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/tools/jirastats-server/response"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

const getTicketsPath = "/data/tickets"
const getTicketDatesPath = "/data/tickets/dates/{projectId}"

type DataTickets struct {
	tm *database.TicketMapper
}

func NewDataTickets(db *sql.DB, router *mux.Router) {
	dt := new(DataTickets)
	dt.tm = database.NewTicketMapper(db)
	router.HandleFunc(getTicketsPath, dt.GetTickets).Methods(http.MethodGet)
	router.HandleFunc(getTicketDatesPath, dt.GetTicketDates).Methods(http.MethodGet)
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

func (dt *DataTickets) GetTicketDates(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	projectId, err := strconv.Atoi(vars["projectId"])
	if err != nil {
		msg := "Stats could not be loaded as project id was not provided"
		e := response.NewErrorJson(msg, res)
		e.SendBadRequest()
		return
	}

	log.Debugf("Get all dates from tickets for project: %d", projectId)

	data := make([]*response.DateSelect, 0)
	data = append(data, response.NewDateSelect(time.Time{}))

	result, err := dt.tm.GetDistinctCreatedAtByProjectId(projectId)
	if err != nil {
		msg := fmt.Sprintf("Dates of tickets could not be found: %s", err.Error())
		e := response.NewErrorJson(msg, res)
		e.SendNotFound()
		return
	}

	for _, v := range result {
		data = append(data, response.NewDateSelect(v))
	}

	success := response.NewSuccessJson(data, res)
	success.SendOK()
}
