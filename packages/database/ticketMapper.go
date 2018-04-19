package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/rebel-l/jirastats/packages/models"
	"github.com/rebel-l/jirastats/packages/utils"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type TicketMapper struct {
	table *TicketTable
}

func NewTicketMapper(db *sql.DB) *TicketMapper {
	tm := new(TicketMapper)
	tm.table = NewTicketTable(NewStatement(db))
	return tm
}

func (tm *TicketMapper) Load() (collection []*models.Ticket, err error) {
	rows, err := tm.table.Select("")
	defer rows.Close()
	if err != nil {
		return
	}

	collection = tm.mapRows(rows)
	return
}

func (tm *TicketMapper) LoadNotExpired(projectId int) (collection []*models.Ticket, err error) {
	rows, err := tm.table.Select("`project_id` = ? AND `expired` IS NULL", projectId)
	defer rows.Close()
	if err != nil {
		return
	}

	collection = tm.mapRows(rows)
	return
}

func (tm *TicketMapper) mapRows(rows *sql.Rows) (collection []*models.Ticket) {
	for rows.Next() {
		t := models.NewTicket()
		err := tm.mapRowToModel(rows, t)
		if err != nil {
			log.Warnf("Not able to map ticket: %s", err.Error())
			continue
		}

		collection = append(collection, t)
	}
	return
}

func (tm *TicketMapper) LoadLastByKey(key string, projectId int) (t *models.Ticket, err error) {
	t = models.NewTicket()
	rows, err := tm.table.SelectComplex("`key` = ? AND project_id = ?", "created_at DESC, expired ASC LIMIT 1", "", "", key, projectId)
	defer rows.Close()
	if err != nil {
		return
	}

	if rows.Next() {
		err = tm.mapRowToModel(rows, t)
	}
	return
}

func (tm *TicketMapper) Save(model *models.Ticket) (err error) {
	var expired sql.NullString
	if model.Expired.Year() > 1 {
		expired.Valid = true
		expired.String = model.Expired.Format(dateTimeFormat)
	}

	if model.Id == 0 {
		// insert
		id, err := tm.table.Insert(
			model.Key,
			model.ProjectId,
			model.Summary,
			strings.Join(model.Components, ","),
			strings.Join(model.Labels, ","),
			utils.BtoI(model.IsNew),
			model.StatusByJira,
			model.StatusClustered,
			model.Priority,
			model.Issuetype,
			utils.BtoI(model.Removed),
			model.CreatedAtByJira.Format(dateTimeFormat),
			model.LastUpdatedByJira.Format(dateTimeFormat),
			model.CreatedAt.Format(dateTimeFormat),
			expired)

		if err != nil {
			return errors.New(fmt.Sprintf("Not able to insert ticket in database: %s", err.Error()))
		}
		model.Id = id
	} else {
		// update
		rowsAffect, err := tm.table.Update(
			model.Id,
			model.Key,
			model.ProjectId,
			model.Summary,
			strings.Join(model.Components, ","),
			strings.Join(model.Labels, ","),
			utils.BtoI(model.IsNew),
			model.StatusByJira,
			model.StatusClustered,
			model.Priority,
			model.Issuetype,
			utils.BtoI(model.Removed),
			model.CreatedAtByJira.Format(dateTimeFormat),
			model.LastUpdatedByJira.Format(dateTimeFormat),
			model.CreatedAt.Format(dateTimeFormat),
			expired)
		if err != nil || rowsAffect != 1 {
			return errors.New(fmt.Sprintf("Not able to update ticket in database: %s", err.Error()))
		}
	}

	return
}

func (tm *TicketMapper) CountStatusClusteredAndNotExpired(status string, projectId int) (count int, err error) {
	count, err = tm.table.Count("`status_clustered` = ? AND project_id = ? AND `expired` IS NULL", status, projectId)
	return
}

func (tm *TicketMapper) CountStatusClusteredFromDay(status string, day time.Time, projectId int) (count int, err error) {
	where := "`status_clustered` = ? AND project_id = ?"
	fields := "MAX(`expired`)"
	group := "`key` HAVING COUNT(*) = 1"
	rows, err := tm.table.SelectComplex(where, "", fields, group, status, projectId)
	defer rows.Close()
	if err != nil {
		return
	}

	var maxExpired sql.NullString
	dayString := day.Format(dateFormat)
	for rows.Next() {
		err = rows.Scan(&maxExpired)
		if err != nil {
			return
		}

		if maxExpired.Valid && dayString == maxExpired.String[0:10]{
			count++
		}
	}

	return
}

func (tm *TicketMapper) mapRowToModel(rows *sql.Rows, t *models.Ticket) (err error) {
	var components,
		labels,
		createdAtByJira,
		lastUpdatedAtByJira,
		createdAt string
	var expired sql.NullString
	var isNew,
		removed int

	err = rows.Scan(
		&t.Id,
		&t.Key,
		&t.ProjectId,
		&t.Summary,
		&components,
		&labels,
		&isNew,
		&t.StatusByJira,
		&t.StatusClustered,
		&t.Priority,
		&t.Issuetype,
		&removed,
		&createdAtByJira,
		&lastUpdatedAtByJira,
		&createdAt,
		&expired)
	t.Components = strings.Split(components, ",")
	t.Labels = strings.Split(labels, ",")
	t.CreatedAtByJira, _ = time.Parse(dateTimeFormat, createdAtByJira)
	t.LastUpdatedByJira, _ = time.Parse(dateTimeFormat, lastUpdatedAtByJira)
	t.CreatedAt, _ = time.Parse(dateTimeFormat, createdAt)
	t.IsNew = utils.ItoB(isNew)
	t.Removed = utils.ItoB(removed)
	if expired.Valid {
		t.Expired, _ = time.Parse(dateTimeFormat, expired.String)
	}

	return
}
