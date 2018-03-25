package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/rebel-l/jirastats/packages/models"
)

type StatsMapper struct {
	table *StatsTable
}

func NewStatsMapper(db *sql.DB) *StatsMapper {
	sm := new(StatsMapper)
	sm.table = NewStatsTable(NewStatement(db))
	return sm
}

func (sm *StatsMapper) Save(model *models.Stats) (err error) {
	if model.Id == 0 {
		// insert
		id, err := sm.table.Insert(
			model.ProjectId,
			model.Open,
			model.Closed,
			model.New,
			model.CreatedAt.Format(dateFormat))

		if err != nil {
			return errors.New(fmt.Sprintf("Not able to insert stats in database: %s", err.Error()))
		}
		model.Id = id
	} else {
		// update
		rowsAffect, err := sm.table.Update(
			model.Id,
			model.ProjectId,
			model.Open,
			model.Closed,
			model.New,
			model.CreatedAt.Format(dateFormat))
		if err != nil || rowsAffect != 1 {
			return errors.New(fmt.Sprintf("Not able to update stats in database: %s", err.Error()))
		}
	}

	return
}
