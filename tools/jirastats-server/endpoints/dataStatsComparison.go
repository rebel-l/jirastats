package endpoints

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/models"
	"github.com/rebel-l/jirastats/tools/jirastats-server/response"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

const dataStatsComparisonPath = "/data/stats/comparison/{projectId}/{versionLeft}/{versionRight}"

type DataStatsComparison struct {
	tm *database.TicketMapper
	pm *database.ProjectMapper
	stats *response.StatsComparison
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
	dsc.stats = response.NewStatsComparison(project)
	dsc.stats.SetVersionLeft(versionLeft)
	dsc.stats.SetVersionRight(versionRight)

	data := make(map[string][]*models.Ticket, 2)
	data["left"], err = dsc.getData(versionLeft)
	if err != nil {
		msg := fmt.Sprintf(
			"Could not find tickets for project with id %d and date %s",
			projectId,
			versionLeft.Format(response.DateFormatInternal))
		e := response.NewErrorJson(msg, res)
		e.SendNotFound()
		return
	}

	data["right"], err = dsc.getData(versionRight)
	if err != nil {
		msg := fmt.Sprintf(
			"Could not find tickets for project with id %d and date %s",
			projectId,
			versionLeft.Format(response.DateFormatInternal))
		e := response.NewErrorJson(msg, res)
		e.SendNotFound()
		return
	}

	for k, v := range data {
		for _, t := range v {
			log.Debugf("Version: %s, Ticket: %s", k, t.Key)
		}
	}

	success := response.NewSuccessJson(dsc.stats, res)
	success.SendOK()
}

func (dsc *DataStatsComparison) getParams(req *http.Request) (projectId int, versionLeft time.Time, versionRight time.Time, err error) {
	vars := mux.Vars(req)
	projectId, err = strconv.Atoi(vars["projectId"])
	if err != nil {
		err = errors.New("stats could not be loaded as project id was not provided")
		return
	}

	errMsg := "stats could not be loaded as left or right version was not provided: %s"
	versionLeft, err = time.Parse(response.DateFormatInternal, vars["versionLeft"])
	if err != nil {
		err = errors.New(fmt.Sprintf(errMsg, err.Error()))
		return
	}

	versionRight, err = time.Parse(response.DateFormatInternal, vars["versionRight"])
	if err != nil {
		err = errors.New(fmt.Sprintf(errMsg, err.Error()))
		return
	}

	return
}

func (dsc *DataStatsComparison) getData(start time.Time) (data []*models.Ticket, err error) {
	if start.IsZero() {
		data, err = dsc.tm.LoadNotExpired(dsc.stats.Project.Id)
		return
	}

	end := start.AddDate(0, 0, 1)
	data, err = dsc.tm.LoadHistorical(dsc.stats.Project.Id, start, end)
	return
}
