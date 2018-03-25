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
const ticketTableIndex = "CREATE UNIQUE INDEX IF NOT EXISTS ticket_key_idx ON %s (`key`, `expired`);"
const ticketTableInsert =
	"INSERT INTO %s (" +
		"`key`, `project_id`, `summary`, `components`, `labels`, " +
		"`status_by_jira`, `status_clustered`, `priority`, `issuetype`, `created_at_by_jira`, " +
		"`last_updated_by_jira`, `created_at`, `expired`" +
	") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
const ticketTableUpdate =
	"UPDATE %s SET " +
		"`key` = ?, `project_id` = ?, `summary` = ?, `components` = ?, `labels` = ?, " +
		"`status_by_jira` = ?, `status_clustered` = ?, `priority` = ?, `issuetype` = ?, `created_at_by_jira` = ?, " +
		"`last_updated_by_jira` = ?, `created_at` = ?, `expired` = ? " +
	"WHERE `id` = ?"


// TicketTable represents the database table for tickets
type TicketTable struct {
	statement *Statement
}

func NewTicketTable(statement *Statement) *TicketTable {
	t := new(TicketTable)
	t.statement = statement
	return t
}

func (t *TicketTable) Select(where string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = t.SelectComplex(where, "", "", "", args...)
	return
}

func (t *TicketTable) SelectComplex(
	where string,
	order string,
	fields string,
	group string,
	args ...interface{}) (rows *sql.Rows, err error){

	rows, err = t.statement.doSelect(ticketTableName, where, order, fields, group, args...)
	return
}

func (t *TicketTable) Insert(
	key string,
	projectId int,
	summary string,
	components string,
	labels string,
	statusByJira string,
	statusClustered string,
	priority string,
	issueType string,
	createdAtByJira string,
	lastUpdatedAtByJira string,
	createdAt string,
	expired sql.NullString) (id int, err error) {

	res, err := t.statement.execute(
		ticketTableInsert,
		ticketTableName,
		key,
		projectId,
		summary,
		components,
		labels,
		statusByJira,
		statusClustered,
		priority,
		issueType,
		createdAtByJira,
		lastUpdatedAtByJira,
		createdAt,
		expired)
	if err != nil {
		return
	}

	id64, err := res.LastInsertId()
	id = int(id64)

	return
}

func (t *TicketTable) Update(
	id int,
	key string,
	projectId int,
	summary string,
	components string,
	labels string,
	statusByJira string,
	statusClustered string,
	priority string,
	issueType string,
	createdAtByJira string,
	lastUpdatedAtByJira string,
	createdAt string,
	expired sql.NullString) (rowsAffected int, err error) {

	res, err := t.statement.execute(
		ticketTableUpdate,
		ticketTableName,
		key,
		projectId,
		summary,
		components,
		labels,
		statusByJira,
		statusClustered,
		priority,
		issueType,
		createdAtByJira,
		lastUpdatedAtByJira,
		createdAt,
		expired,
		id)
	if err != nil {
		return
	}

	rowsAffected64, err := res.RowsAffected()
	rowsAffected = int(rowsAffected64)
	return
}

func (t *TicketTable) Count(where string, args ...interface{}) (counter int, err error) {
	rows, err := t.statement.doSelect(ticketTableName, where, "", exprCount, "", args...)
	defer rows.Close()
	if err != nil {
		return
	}

	if rows.Next() {
		err = rows.Scan(&counter)
	}

	return
}

func (t *TicketTable) Truncate() error {
	_, err := t.statement.execute(TruncateTable, ticketTableName)
	return err
}

func (t *TicketTable) CreateStructure() (err error) {
	log.Debugf("Create structure for %s", ticketTableName)
	// create table
	_, err = t.statement.execute(ticketTableStructure, ticketTableName)
	if err != nil {
		return
	}

	// create index
	_, err = t.statement.execute(ticketTableIndex, ticketTableName)
	return
}
