package database

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

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

type ProjectTable struct {
	statement *Statement
}

func NewProjectTable(statement *Statement) *ProjectTable {
	p := new(ProjectTable)
	p.statement = statement
	return p
}

func (p *ProjectTable) Insert(
	name string,
	keys string,
	jql string,
	knownSpeed float32,
	mapOpenStatus string,
	mapClosedStatus string) (id int, err error) {

	res, err := p.statement.execute(
		projectTableInsert,
		projectTableName,
		name,
		jql,
		keys,
		mapOpenStatus,
		mapClosedStatus,
		knownSpeed)
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
	_, err = p.statement.execute(projectTableStructure, projectTableName)
	if err != nil {
		return
	}

	// create index
	_, err = p.statement.execute(projectTableIndex, projectTableName)
	return
}

func (p *ProjectTable) Select(where string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = p.SelectComplex(where, "", "", "", args...)
	return
}

func (p *ProjectTable) SelectComplex(
	where string,
	order string,
	fields string,
	group string,
	args ...interface{}) (rows *sql.Rows, err error){

	rows, err = p.statement.doSelect(projectTableName, where, order, fields, group, args...)
	return
}
