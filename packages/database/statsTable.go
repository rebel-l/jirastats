package database

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

type StatsTable struct {
	db *sql.DB
}

const statsTableName = "stats"
const statsTableStructure =
	"CREATE TABLE IF NOT EXISTS `%s` (" +
		"`id` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`project_id` INTEGER NOT NULL," +
		"`status_id` INTEGER NOT NULL," +
		"`counter` INTEGER DEFAULT 0 NOT NULL," +
		"`created_at` DATE NOT NULL," +
		"FOREIGN KEY (project_id) REFERENCES project(`id`)," +
		"FOREIGN KEY (status_id) REFERENCES stats_status(`id`)" +
	");"

func NewSatsTable(db *sql.DB) *StatsTable {
	s := new(StatsTable)
	s.db = db
	return s
}

func (s *StatsTable) Truncate() (err error) {
	stmt, err := s.db.Prepare(s.getTruncateStatement())
	if err != nil {
		return
	}
	_, err = stmt.Exec()
	return
}

func (s *StatsTable) getTruncateStatement() string {
	return createDatabseStatement(TruncateTable, statsTableName)
}

func (s *StatsTable) CreateStructure() (err error) {
	log.Debugf("Create structure for %s", statsTableName)
	// create table
	err = executeStatement(s.db, s.getCreateTableStatement())
	return
}

func (s *StatsTable) getCreateTableStatement() string {
	return createDatabseStatement(statsTableStructure, statsTableName)
}
