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

func NewProjectTable(db *sql.DB) *ProjectTable {
	p := new(ProjectTable)
	p.db = db
	return p
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
