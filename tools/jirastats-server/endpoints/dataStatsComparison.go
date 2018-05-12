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
	"errors"
)

const dataStatsComparisonPath = "/data/stats/comparison/{projectId}/{versionLeft}/{versionRight}"

type DataStatsComparison struct {
	tm *database.TicketMapper
	pm *database.ProjectMapper
}

func NewDataStatsComparison(db *sql.DB, router *mux.Router) {
	dsc := new(DataStatsComparison)
	dsc.tm = database.NewTicketMapper(db)
	dsc.pm = database.NewProjectMapper(db)
	router.HandleFunc(dataStatsComparisonPath, dsc.GetStats).Methods(http.MethodGet)
}

func (dsc *DataStatsComparison) GetStats(res http.ResponseWriter, req *http.Request) {
	projectId, versionLeft, versionRight, err := dsc.getParams(req)
	if err != nil {
		e := response.NewErrorJson(err.Error(), res)
		e.SendBadRequest()
		return
	}

	project, err := dsc.pm.LoadProjectById(projectId)
	if err != nil {
		msg := fmt.Sprintf("Project with id %d could not be found", projectId)
		e := response.NewErrorJson(msg, res)
		e.SendNotFound()
		return
	}

	log.Debugf("Get all comparison stats for project: %d, parameters: %s (left), %s (right)", projectId, versionLeft, versionRight)
	stats := response.NewStatsComparison(project)
	stats.SetVersionLeft(versionLeft)
	stats.SetVersionRight(versionRight)

	success := response.NewSuccessJson(stats, res)
	success.SendOK()
}

func (dsc *DataStatsComparison) getParams(req *http.Request) (projectId int, versionLeft string, versionRight string, err error) {
	vars := mux.Vars(req)
	projectId, err = strconv.Atoi(vars["projectId"])
	if err != nil {
		err = errors.New("stats could not be loaded as project id was not provided")
		return
	}

	versionLeft = vars["versionLeft"]
	versionRight = vars["versionRight"]
	if versionLeft == "" ||  versionRight == "" {
		err = errors.New("stats could not be loaded as left or right version was not provided")
	}
	return
}
