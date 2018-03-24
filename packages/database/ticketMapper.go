package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/rebel-l/jirastats/packages/models"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type TicketMapper struct {
	table *TicketTable
}

func NewTicketMapper(db *sql.DB) *TicketMapper {
	tm := new(TicketMapper)
	tm.table = NewTicketTable(db)
	return tm
}

func (tm *TicketMapper) Load() (collection []*models.Ticket, err error) {
	rows, err := tm.table.Select("")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		t := models.NewTicket()
		err = tm.mapRowToModel(rows, t)
		if err != nil {
			log.Warnf("Not able to map ticket")
		}

		collection = append(collection, t)
	}
	return
}

func (tm *TicketMapper) LoadByKey(key string) (t *models.Ticket, err error) {
	t = models.NewTicket()
	rows, err := tm.table.Select("`key` = ? AND `expired` IS NULL", key)
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
	if model.Id == 0 {
		// insert
		id, err := tm.table.Insert(
			model.Key,
			model.ProjectId,
			model.Summary,
			strings.Join(model.Components, ","),
			strings.Join(model.Labels, ","),
			model.StatusByJira,
			model.StatusClustered,
			model.Priority,
			model.Issuetype,
			model.CreatedAtByJira.Format(dateTimeFormat),
			model.LastUpdatedByJira.Format(dateTimeFormat),
			model.CreatedAt.Format(dateTimeFormat))

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
			model.StatusByJira,
			model.StatusClustered,
			model.Priority,
			model.Issuetype,
			model.CreatedAtByJira.Format(dateTimeFormat),
			model.LastUpdatedByJira.Format(dateTimeFormat),
			model.CreatedAt.Format(dateTimeFormat),
			model.Expired.Format(dateTimeFormat))
		if err != nil || rowsAffect != 1 {
			return errors.New(fmt.Sprintf("Not able to update ticket in database: %s", err.Error()))
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

	err = rows.Scan(
		&t.Id,
		&t.Key,
		&t.ProjectId,
		&t.Summary,
		&components,
		&labels,
		&t.StatusByJira,
		&t.StatusClustered,
		&t.Priority,
		&t.Issuetype,
		&createdAtByJira,
		&lastUpdatedAtByJira,
		&createdAt,
		&expired)
	t.Components = strings.Split(components, ",")
	t.Labels = strings.Split(labels, ",")
	t.CreatedAtByJira, _ = time.Parse(dateTimeFormat, createdAtByJira)
	t.LastUpdatedByJira, _ = time.Parse(dateTimeFormat, lastUpdatedAtByJira)
	t.CreatedAt, _ = time.Parse(dateTimeFormat, createdAt)
	if expired.Valid {
		t.Expired, _ = time.Parse(dateTimeFormat, expired.String)
	}

	return
}
