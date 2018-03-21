package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/rebel-l/jirastats/packages/models"
	log "github.com/sirupsen/logrus"
)

type ProjectMapper struct {
	table *ProjectTable
}

func NewProjectMapper(db *sql.DB) *ProjectMapper {
	pm := new(ProjectMapper)
	pm.table = NewProjectTable(db)
	return pm
}

func (pm *ProjectMapper) Save(model *models.Project) (err error) {
	if model.Id == 0 {
		// insert
		id, err := pm.table.Insert(model.Name, model.Keys, model.Jql, model.KnownSpeed, model.MapOpenStatus, model.MapClosedStatus)
		if err != nil {
			return errors.New(fmt.Sprintf("Not able to insert project in database: %s", err.Error()))
		}
		model.Id = id
	} else {
		// update
		log.Warn("Update for project not implemented yet!")
	}

	return
}

func (pm *ProjectMapper) Load() (projects []*models.Project, err error) {
	rows, err := pm.table.Select("")
	defer rows.Close()
	if err != nil {
		err = errors.New(fmt.Sprintf("Not able to load projects from database: %s", err.Error()))
		return
	}

	var id int
	var name string
	var jql string
	var keys string
	var mapOpenStatus string
	var mapClosedStatus string
	var knownSpeed float32
	for rows.Next() {
		err = rows.Scan(&id, &name, &jql, &keys, &mapOpenStatus, &mapClosedStatus, &knownSpeed)
		if err != nil {
			log.Warn("Mapping of project not possible")
			continue
		}

		project := models.NewProject()
		project.Id = id
		project.Name = name
		project.Jql = jql
		project.Keys = keys
		project.MapOpenStatus = mapOpenStatus
		project.MapClosedStatus = mapClosedStatus
		project.KnownSpeed = knownSpeed

		projects = append(projects, project)
	}

	log.Debugf("Number of projects found: %d", len(projects))
	return
}
