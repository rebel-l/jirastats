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

const dataStatsSpeedPath = "/data/stats/speed/{projectId}"

type DataStatsSpeed struct {
	sm *database.StatsMapper
	pm *database.ProjectMapper
	stats *response.Stats
	series map[string]*response.Serie
}

func NewDataStatsSpeed(db *sql.DB, router *mux.Router) {
	ds := new(DataStatsSpeed)
	ds.sm = database.NewStatsMapper(db)
	ds.pm = database.NewProjectMapper(db)
	router.HandleFunc(dataStatsSpeedPath, ds.GetStats).Methods(http.MethodGet)
}

func (ds *DataStatsSpeed) GetStats(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	projectId, err := strconv.Atoi(vars["projectId"])
	if err != nil {
		msg := "Stats could not be loaded as project id was not provided"
		e := response.NewErrorJson(msg, res)
		e.SendBadRequest()
		return
	}

	log.Debugf("Get all speed stats for project: %d", projectId)

	var ok bool
	ds.stats, ok = response.NewStats(projectId, ds.pm, res)
	if ok == false {
		return
	}

	ok = ds.setStats(res)
	if ok == false {
		return
	}

	success := response.NewSuccessJson(ds.stats, res)
	success.SendOK()
}

func (ds *DataStatsSpeed) setStats(res http.ResponseWriter) bool {
	stats, err := ds.sm.LoadByProjectId(ds.stats.ProjectId)
	if err != nil {
		msg := fmt.Sprintf("Not able to load stats for project id %d: %s", ds.stats.ProjectId, err.Error())
		e := response.NewErrorJson(msg, res)
		e.SendInternalServerError()
		return false
	}

	if len(stats) == 0 {
		msg := fmt.Sprintf("No stats found for project id: %d", ds.stats.ProjectId)
		e := response.NewErrorJson(msg, res)
		e.SendNotFound()
		return false
	}

	ds.series = make(map[string]*response.Serie, 3)
	ds.series["closed"] = new(response.Serie)
	ds.series["closed"].Name = "Closed"
	ds.series["new"] = new(response.Serie)
	ds.series["new"].Name = "New"
	ds.series["speed"] = new(response.Serie)
	ds.series["speed"].Name = "Speed Index"

	var previousYear, previousWeek int
	var closedTickets, newTickets int
	for i := 0; i < len(stats); i++ {
		/*
		Iteration	|	previous year/week	|	actual year/week	|	actions
			0		|		0/0 			| 		2018/2 			| 	sum up
			1		|		2018/2			| 		2018/2 			| 	sum up
			2		|		2018/2			| 		2018/3 			| 	add actual serie, reset counter, sum up
			3		|		2018/3			| 		2018/3 			| 	sum up + add serie as it is last
		*/

		v := stats[i]
		actualYear, actualWeek := v.CreatedAt.ISOWeek()

		if i == 0 {
			previousYear = actualYear
			previousWeek = actualWeek
		}

		if actualYear != previousYear || actualWeek != previousWeek {
			// add numbers to serie
			ds.addData(previousYear, previousWeek, closedTickets, newTickets)

			// reset counters
			closedTickets = 0
			newTickets = 0
		}

		// sum up
		closedTickets += v.Closed
		newTickets += v.New

		if i == len(stats) - 1 {
			// add numbers of last run to serie
			ds.addData(actualYear, actualWeek, closedTickets, newTickets)
		}

		previousYear = actualYear
		previousWeek = actualWeek
	}

	ds.stats.Series = append(ds.stats.Series, ds.series["closed"], ds.series["new"], ds.series["speed"])

	return true
}

func (ds *DataStatsSpeed) addData(year int, week int, closedTickets int, newTickets int) {
	ds.stats.Categories = append(ds.stats.Categories, fmt.Sprintf("%d/%d", year, week))
	ds.series["closed"].AddDataInt(closedTickets)
	ds.series["new"].AddDataInt(newTickets)
	ds.series["speed"].AddDataInt(closedTickets - newTickets)
}
