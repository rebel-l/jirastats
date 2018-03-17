package database

import (
	"database/sql"
	"fmt"
)

const TicketTableName = "ticket"

// TicektTable represents the database table for tickets
type TicketTable struct {
	db *sql.DB
}

func NewTicketTable(db *sql.DB) *TicketTable {
	t := new(TicketTable)
	t.db = db
	return t
}

func (t *TicketTable) Load() (rows *sql.Rows, err error){
	stmt, err := t.db.Prepare(t.getSelectAllStatement())
	if err != nil {
		return
	}

	rows, err = stmt.Query()
	return
}

func (t *TicketTable) getStatement(statement string) string {
	return fmt.Sprintf(statement, TicketTableName)
}

func (t *TicketTable) getSelectAllStatement() string {
	return t.getStatement(SelectAllStatement)
}

func (t *TicketTable) Truncate() (err error) {
	stmt, err := t.db.Prepare(t.getStatement(TruncateTable))
	if err != nil {
		return
	}
	_, err = stmt.Exec()
	return
}
