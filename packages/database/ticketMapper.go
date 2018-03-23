package database

import (
	"database/sql"
	"github.com/rebel-l/jirastats/packages/models"
	log "github.com/sirupsen/logrus"
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

func (tm *TicketMapper) mapRowToModel(rows *sql.Rows, t *models.Ticket) (err error) {
	err = rows.Scan(&t.Id, &t.Key)
	return
}
