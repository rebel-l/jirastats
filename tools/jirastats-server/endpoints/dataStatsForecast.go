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

const dataStatsForecastPath = "/data/stats/forecast/{projectId}"

type DataStatsForecast struct {
	pm *database.ProjectMapper
	sm *database.StatsMapper
}

func NewDataStatsForecast(db *sql.DB, router *mux.Router) {
	dsf := new(DataStatsForecast)
	dsf.pm = database.NewProjectMapper(db)
	dsf.sm = database.NewStatsMapper(db)
	router.HandleFunc(dataStatsForecastPath, dsf.GetStats).Methods(http.MethodGet)
}

func (dsf *DataStatsForecast) GetStats(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	projectId, err := strconv.Atoi(vars["projectId"])
	if err != nil {
		msg := "Stats could not be loaded as project id was not provided"
		e := response.NewErrorJson(msg, res)
		e.SendBadRequest()
		return
	}

	project, err := dsf.pm.LoadProjectById(projectId)
	if err != nil {
		msg := fmt.Sprintf("Project with id %d could not be found", projectId)
		e := response.NewErrorJson(msg, res)
		e.SendBadRequest()
		return
	}

	log.Debugf("Get all forecast stats for project: %d", projectId)

	statsData, err := dsf.sm.LoadByProjectId(projectId)
	if err != nil {
		msg := fmt.Sprintf("Stats for project with id %d could not be found", projectId)
		e := response.NewErrorJson(msg, res)
		e.SendNotFound()
		return
	}

	stats := response.NewStatsForecast(project)
	stats.Calculate(statsData)
	log.Debugf("Data on forecast stats for project: %d processed!", projectId)

	success := response.NewSuccessJson(stats, res)
	success.SendOK()
	log.Debugf("Response on forecast stats for project: %d finished!", projectId)
}
