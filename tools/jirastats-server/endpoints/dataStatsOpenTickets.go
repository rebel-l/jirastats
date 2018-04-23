package endpoints

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/tools/jirastats-server/response"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"fmt"
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

	log.Debugf("Get all open ticket stats for project: %d", projectId)

	tickets, err := ds.tm.LoadNotExpired(projectId)
	if err != nil {
		msg := fmt.Sprintf("Stats could not be loaded: %s", err.Error())
		e := response.NewErrorJson(msg, res)
		e.SendBadRequest()
		return
	}

	/*
		TODO: data to collect
			components
			issue types
			status
			priority
			techdebt
			labels (except tech debt)

			absolute numbers for a data table, relative numbers for chart
	*/
	dataStatus := make(map[string]*response.TableData)
	countTickets := len(tickets)
	for _, t := range tickets {
		if v, ok := dataStatus[t.StatusByJira]; ok {
			v.Value++
		} else {
			dataStatus[t.StatusByJira] = response.NewTableData(t.StatusByJira, 1)
		}
	}

	finalDataStatus := make([]*response.PieChart, len(dataStatus))
	pc := response.NewPieChartTable("Status", finalDataStatus)
	i := 0
	maxValue := 0
	maxItem := 0
	for _, d := range dataStatus {
		if d.Value > maxValue {
			maxValue = d.Value
			maxItem = i
		}
		finalDataStatus[i] = response.NewPieChart(d.Name, float64(d.Value) * 100.0 / float64(countTickets))
		pc.DataTable = append(pc.DataTable, d)
		i++
	}
	finalDataStatus[maxItem].Sliced = true
	finalDataStatus[maxItem].Selected = true

	// TODO: parallize processing data

	success := response.NewSuccessJson(pc, res)
	success.SendOK()
}
