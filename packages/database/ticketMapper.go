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
	rows, err := tm.table.Load()
	if err != nil {
		return
	}

	for rows.Next() {
		t := new(models.Ticket)
		err = rows.Scan(&t.Id, &t.Key)
		if err != nil {
			log.Warnf("Not able to map ticket")
		}

		collection = append(collection, t)
	}
	return
}
