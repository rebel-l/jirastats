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

const dataStatsProgressPath = "/data/stats/progress/{projectId}"

type DataStatsProgress struct {
	pm *database.ProjectMapper
	sm *database.StatsMapper
	stats *response.Stats
}

func NewDataStatsProgress(db *sql.DB, router *mux.Router) {
	ds := new(DataStatsProgress)
	ds.sm = database.NewStatsMapper(db)
	ds.pm = database.NewProjectMapper(db)
	router.HandleFunc(dataStatsProgressPath, ds.GetStats).Methods(http.MethodGet)
}

func (ds *DataStatsProgress) GetStats(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	projectId, err := strconv.Atoi(vars["projectId"])
	if err != nil {
		msg := "Stats could not be loaded as project id was not provided"
		e := response.NewErrorJson(msg, res)
		e.SendBadRequest()
		return
	}

	log.Debugf("Get all progress stats for project: %d", projectId)

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



func (ds *DataStatsProgress) setStats(res http.ResponseWriter) bool {
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

	openSeries := new(response.Serie)
	openSeries.Name = "Open"
	closedSeries := new(response.Serie)
	closedSeries.Name = "Closed"
	newSeries := new(response.Serie)
	newSeries.Name = "New"

	for _, v := range stats {
		ds.stats.AddCategory(v.CreatedAt.Format(response.DateFormatDisplay))
		openSeries.AddDataInt(v.Open)
		closedSeries.AddDataInt(v.Closed)
		newSeries.AddDataInt(v.New)
	}

	ds.stats.Series = append(ds.stats.Series, openSeries, closedSeries, newSeries)

	return true
}
