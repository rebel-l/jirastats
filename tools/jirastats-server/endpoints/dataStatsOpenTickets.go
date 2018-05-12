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
)

const dataStatsOpenTicketsPath = "/data/stats/opentickets/{projectId}"

type DataStatsOpenTickets struct {
	tm *database.TicketMapper
	pm *database.ProjectMapper
	stats *response.Stats
	series map[string]*response.Serie
}

func NewDataStatsOpenTickets(db *sql.DB, router *mux.Router) {
	ds := new(DataStatsOpenTickets)
	ds.tm = database.NewTicketMapper(db)
	ds.pm = database.NewProjectMapper(db)
	router.HandleFunc(dataStatsOpenTicketsPath, ds.GetStats).Methods(http.MethodGet)
}

func (ds *DataStatsOpenTickets) GetStats(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	projectId, err := strconv.Atoi(vars["projectId"])
	if err != nil {
		msg := "Stats could not be loaded as project id was not provided"
		e := response.NewErrorJson(msg, res)
		e.SendBadRequest()
		return
	}

	project, err := ds.pm.LoadProjectById(projectId)
	if err != nil {
		msg := fmt.Sprintf("Project with id %d could not be found", projectId)
		e := response.NewErrorJson(msg, res)
		e.SendNotFound()
		return
	}

	log.Debugf("Get all open ticket stats for project: %d", projectId)

	tickets, err := ds.tm.LoadNotExpired(projectId)
	if err != nil {
		msg := fmt.Sprintf("Stats could not be loaded: %s", err.Error())
		e := response.NewErrorJson(msg, res)
		e.SendNotFound()
		return
	}

	ticketCounter := response.NewTicketCounter()
	ticketCounter.Count(tickets)

	stats := response.NewStatsPiechartTable(project)
	stats.GeneratePiechartTables(ticketCounter.Tables, len(tickets))

	success := response.NewSuccessJson(stats, res)
	success.SendOK()
}
