package database

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

const ticketTableName = "ticket"
const ticketTableStructure =
	"CREATE TABLE IF NOT EXISTS `%s` (" +
		"`id` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`key` CHAR(50) NOT NULL," +
		"`project_id` INTEGER NOT NULL," +
		"`summary` VARCHAR(2000) NOT NULL," +
		"`components` VARCHAR(2000) NULL," +
		"`labels` VARCHAR(2000) NULL," +
		"`status_by_jira` CHAR(50) NOT NULL," +
		"`status_clustered` CHAR(50) NOT NULL," +
		"`priority` CHAR(50) NOT NULL," +
		"`issuetype` CHAR(50) NULL," +
		"`created_at_by_jira` DATETIME NOT NULL," +
		"`last_updated_by_jira` DATETIME NOT NULL," +
		"`created_at` DATETIME NOT NULL," +
		"`expired` DATETIME NULL," +
		"FOREIGN KEY (project_id) REFERENCES project(`id`)" +
");"
const ticketTableIndex = "CREATE UNIQUE INDEX IF NOT EXISTS ticket_key_idx ON %s (`key`);"


// TicektTable represents the database table for tickets
type TicketTable struct {
	db *sql.DB
}

func NewTicketTable(db *sql.DB) *TicketTable {
	t := new(TicketTable)
	t.db = db
	return t
}

func (t *TicketTable) Select(where string, args ...interface{}) (rows *sql.Rows, err error){
	statement := t.getSelectAllStatement()
	if where != "" {
		statement += " WHERE " + where
	}

	stmt, err := t.db.Prepare(statement)
	if err != nil {
		return
	}

	if args != nil {
		rows, err = stmt.Query(args...)
	} else {
		rows, err = stmt.Query()
	}
	return
}

func (t *TicketTable) Count(where string, args ...interface{}) (counter int, err error) {
	statement := t.getSelectCountStatement()
	if where != "" {
		statement += " WHERE " + where
	}

	stmt, err := t.db.Prepare(statement)
	if err != nil {
		return
	}

	var rows *sql.Rows
	if args != nil {
		rows, err = stmt.Query(args...)
	} else {
		rows, err = stmt.Query()
	}
	defer rows.Close()
	if err != nil {
		return
	}

	if rows.Next() {
		err = rows.Scan(&counter)
	}

	return
}

func (t *TicketTable) getSelectAllStatement() string {
	return createDatabseStatement(SelectAllStatement, ticketTableName)
}


func (t *TicketTable) getSelectCountStatement() string {
	return createDatabseStatement(SelectCountStatement, statsTableName)
}

func (t *TicketTable) getTruncateStatement() string {
	return createDatabseStatement(TruncateTable, ticketTableName)
}

func (t *TicketTable) Truncate() (err error) {
	stmt, err := t.db.Prepare(t.getTruncateStatement())
	if err != nil {
		return
	}
	_, err = stmt.Exec()
	return
}

func (t *TicketTable) CreateStructure() (err error) {
	log.Debugf("Create structure for %s", ticketTableName)
	// create table
	err = executeStatement(t.db, t.getCreateTableStatement())
	if err != nil {
		return
	}

	// create index
	err = executeStatement(t.db, t.getCreateIndexStatement())
	return
}

func (t *TicketTable) getCreateTableStatement() string {
	return createDatabseStatement(ticketTableStructure, ticketTableName)
}

func (t *TicketTable) getCreateIndexStatement() string {
	return createDatabseStatement(ticketTableIndex, ticketTableName)
}
