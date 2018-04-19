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
const dateFormat = "02.01.2006"

type DataStatsProgress struct {
	sm *database.StatsMapper
	pm *database.ProjectMapper
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

	s := new(response.Stats)
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

func (ds *DataStatsProgress) setProjectName(s *response.Stats, res http.ResponseWriter) bool {
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

func (ds *DataStatsProgress) setStats(s *response.Stats, res http.ResponseWriter) bool {
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

	openSeries := new(response.Serie)
	openSeries.Name = "Open"
	closedSeries := new(response.Serie)
	closedSeries.Name = "Closed"
	newSeries := new(response.Serie)
	newSeries.Name = "New"

	for _, v := range stats {
		s.Categories = append(s.Categories, v.CreatedAt.Format(dateFormat))
		openSeries.Data = append(openSeries.Data, v.Open)
		closedSeries.Data = append(closedSeries.Data, v.Closed)
		newSeries.Data = append(newSeries.Data, v.New)
	}

	s.Series = append(s.Series, openSeries, closedSeries, newSeries)

	return true
}
