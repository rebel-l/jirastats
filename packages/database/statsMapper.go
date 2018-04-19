package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/rebel-l/jirastats/packages/models"
	log "github.com/sirupsen/logrus"
	"time"
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

func (sm *StatsMapper) LoadByProjectId(projectId int) (collection []*models.Stats, err error) {
	rows, err := sm.table.SelectComplex("project_id = ?", "created_at ASC", "", "", projectId)
	defer rows.Close()
	if err != nil {
		return
	}

	collection = sm.mapRows(rows)

	return
}

func (sm *StatsMapper) mapRowToModel(rows *sql.Rows, s *models.Stats) (err error) {
	var createdAt string
	err = rows.Scan(&s.Id, &s.ProjectId, &s.Open, &s.Closed, &s.New, &createdAt)
	if err != nil {
		return
	}

	s.CreatedAt, _ = time.Parse(dateTimeFormat, createdAt)
	return
}

func (sm *StatsMapper) mapRows(rows *sql.Rows) (collection []*models.Stats) {
	for rows.Next() {
		s := models.NewStats(0)
		err := sm.mapRowToModel(rows, s)
		if err != nil {
			log.Warnf("Not able to map stats: %s", err.Error())
			continue
		}

		collection = append(collection, s)
	}
	return
}
