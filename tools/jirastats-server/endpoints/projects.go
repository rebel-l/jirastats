package endpoints

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/models"
	"net/http"
	log "github.com/sirupsen/logrus"
)

const projectsPath = "/projects"

type Projects struct {
	pm *database.ProjectMapper
}

type ProjectList struct {
	Projects []*models.Project
}

func NewProjects(db *sql.DB, router *mux.Router) {
	dp := new(Projects)
	dp.pm = database.NewProjectMapper(db)
	router.HandleFunc(projectsPath, dp.GetAllProjects).Methods(http.MethodGet)
}

func (dp *Projects) GetAllProjects(res http.ResponseWriter, req *http.Request) {
	log.Debug("Get all projects")
	projects, err := dp.pm.Load()
	if err != nil {
		msg := fmt.Sprintf("Not able to load tickets: %s", err.Error())
		e := NewErrorJson(msg, res)
		e.SendInternalServerError()
		return
	}

	if len(projects) == 0 {
		msg := "No Projects found. Please create at least one to proceed."
		e := NewErrorJson(msg, res)
		e.SendNotFound()
		return
	}

	payload := ProjectList{projects}
	s := NewSuccessJson(payload, res)
	s.SendOK()
}
