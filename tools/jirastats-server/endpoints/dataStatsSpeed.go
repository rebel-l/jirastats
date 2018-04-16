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

const dataStatsSpeedPath = "/data/stats/speed/{projectId}"

type DataStatsSpeed struct {
	sm *database.StatsMapper
	pm *database.ProjectMapper
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

	s := new(Stats)
	s.ProjectId = projectId
	ok := ds.setProjectName(s, res)
	if ok == false {
		return
	}

	ok = ds.setStats(s, res)
	if ok == false {
		return
	}

	success := response.NewSuccessJson(s, res)
	success.SendOK()
}

func (ds *DataStatsSpeed) setProjectName(s *Stats, res http.ResponseWriter) bool {
	// TODO: maybe can be set by client, not necessary to send it again
	// TODO: add to Stats struct
	project, err := ds.pm.LoadProjectById(s.ProjectId)
	if err != nil {
		msg := fmt.Sprintf("Not able to load project id %d: %s", s.ProjectId, err.Error())
		e := response.NewErrorJson(msg, res)
		e.SendInternalServerError()
		return false
	}

	if project == nil {
		msg := fmt.Sprintf("No project found for id: %d", s.ProjectId)
		e := response.NewErrorJson(msg, res)
		e.SendNotFound()
		return false
	}

	s.ProjectName = project.Name
	return true
}

func (ds *DataStatsSpeed) setStats(s *Stats, res http.ResponseWriter) bool {
	// TODO: cluster by week
	stats, err := ds.sm.LoadByProjectId(s.ProjectId)
	if err != nil {
		msg := fmt.Sprintf("Not able to load stats for project id %d: %s", s.ProjectId, err.Error())
		e := response.NewErrorJson(msg, res)
		e.SendInternalServerError()
		return false
	}

	if len(stats) == 0 {
		msg := fmt.Sprintf("No stats found for project id: %d", s.ProjectId)
		e := response.NewErrorJson(msg, res)
		e.SendNotFound()
		return false
	}

	closedSeries := new(Serie)
	closedSeries.Name = "Closed"
	newSeries := new(Serie)
	newSeries.Name = "New"
	speedSeries := new(Serie)
	speedSeries.Name = "Open"

	for _, v := range stats {
		s.Categories = append(s.Categories, v.CreatedAt.Format(dateFormat))
		closedSeries.Data = append(closedSeries.Data, v.Closed)
		newSeries.Data = append(newSeries.Data, v.New)
		speedSeries.Data = append(speedSeries.Data, v.Closed - v.New)
	}

	s.Series = append(s.Series, closedSeries, newSeries, speedSeries)

	return true
}

