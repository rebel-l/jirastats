package database

import (
	"database/sql"
	"github.com/rebel-l/jirastats/packages/models"
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
			return err
		}
		model.Id = id
	} else {
		// update
	}

	return
}
