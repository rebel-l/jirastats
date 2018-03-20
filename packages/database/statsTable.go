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
		"`status` CHAR(50) NOT NULL," +
		"`counter` INTEGER DEFAULT 0 NOT NULL," +
		"`created_at` DATE NOT NULL," +
		"FOREIGN KEY (project_id) REFERENCES project(`id`)" +
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

func (s *StatsTable) Select(where string, args ...interface{}) (rows *sql.Rows, err error){
	statement := s.getSelectAllStatement()
	if where != "" {
		statement += " WHERE " + where
	}

	stmt, err := s.db.Prepare(statement)
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

func (s *StatsTable) getSelectAllStatement() string {
	return createDatabseStatement(SelectAllStatement, statsTableName)
}
