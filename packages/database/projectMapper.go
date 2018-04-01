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
	pm.table = NewProjectTable(NewStatement(db))
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

func (pm *ProjectMapper) LoadProjectById(id int) (p *models.Project, err error) {
	rows, err := pm.table.Select("id = ?", id)
	defer rows.Close()
	if err != nil {
		return
	}

	p = models.NewProject()
	if rows.Next() {
		p = models.NewProject()
		err = pm.mapRowToModel(rows, p)
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

	projects = pm.mapRows(rows)
	return
}

func (pm *ProjectMapper) HasTickets(project *models.Project) bool {
	tt := NewTicketTable(pm.table.statement)
	counter, err := tt.Count("project_id = ?", project.Id)
	if err != nil {
		log.Errorf("Tickets are not countable from database: %s", err.Error())
		return false
	}

	log.Debugf("%d Tickets found for project %d (%s)", counter, project.Id, project.Name)

	return counter != 0
}

func (pm *ProjectMapper) mapRowToModel(rows *sql.Rows, p *models.Project) (err error) {
	err = rows.Scan(&p.Id, &p.Name, &p.Jql, &p.Keys, &p.MapOpenStatus, &p.MapClosedStatus, &p.KnownSpeed)
	return
}

func (pm *ProjectMapper) mapRows(rows *sql.Rows) (collection []*models.Project) {
	for rows.Next() {
		p := models.NewProject()
		err := pm.mapRowToModel(rows, p)
		if err != nil {
			log.Warnf("Not able to map project: %s", err.Error())
			continue
		}

		collection = append(collection, p)
	}
	return
}
