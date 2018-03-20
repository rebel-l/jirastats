package database

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

type ProjectTable struct {
	db *sql.DB
}

const projectTableName = "project"
const projectTableStructure =
	"CREATE TABLE IF NOT EXISTS `%s` (" +
		"`id` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`name` CHAR(50) NOT NULL," +
		"`jql` VARCHAR(2000) NOT NULL," +
		"`keys` VARCHAR(255) NULL," +
		"`map_open_status` VARCHAR(2000) NOT NULL," +
		"`map_closed_status` VARCHAR(2000) NOT NULL," +
		"`known_speed` REAL NOT NULL" +
");"
const projectTableIndex = "CREATE UNIQUE INDEX IF NOT EXISTS project_name_idx ON %s (`name`);"
const projectTableInsert =
	"INSERT INTO %s (" +
		"`name`, `jql`, `keys`, `map_open_status`, `map_closed_status`, `known_speed`" +
	")" +
	" VALUES (?, ?, ?, ?, ?, ?)"

func NewProjectTable(db *sql.DB) *ProjectTable {
	p := new(ProjectTable)
	p.db = db
	return p
}

func (p *ProjectTable) Insert(name string, keys string, jql string, knownSpeed float32, mapOpenStatus string, mapClosedStatus string) (id int, err error) {
	stmt, err := p.db.Prepare(createDatabseStatement(projectTableInsert, projectTableName))
	if err != nil {
		return
	}

	res, err := stmt.Exec(name, jql, keys, mapOpenStatus, mapClosedStatus, knownSpeed)
	if err != nil {
		return
	}

	id64, err := res.LastInsertId()
	id = int(id64)

	return
}

func (p *ProjectTable) Truncate() error {
	truncateNotImplemented(projectTableName)
	return nil
}

func (p *ProjectTable) CreateStructure() (err error) {
	log.Debugf("Create structure for %s", projectTableName)
	// create table
	err = executeStatement(p.db, p.getCreateTableStatement())
	if err != nil {
		return
	}

	// create index
	err = executeStatement(p.db, p.getCreateIndexStatement())
	return
}

func (p *ProjectTable) getCreateTableStatement() string {
	return createDatabseStatement(projectTableStructure, projectTableName)
}

func (p *ProjectTable) getCreateIndexStatement() string {
	return createDatabseStatement(projectTableIndex, projectTableName)
}

func (p *ProjectTable) Select(where string, args ...interface{}) (rows *sql.Rows, err error){
	statement := p.getSelectAllStatement()
	if where != "" {
		statement += " WHERE " + where
	}
	log.Debugf("Project select statement: %s", statement)

	stmt, err := p.db.Prepare(statement)
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

func (p *ProjectTable) getSelectAllStatement() string {
	return createDatabseStatement(SelectAllStatement, projectTableName)
}
