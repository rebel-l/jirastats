package database

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

type StatsTable struct {
	statement *Statement
}

const statsTableName = "stats"
const statsTableStructure =
	"CREATE TABLE IF NOT EXISTS `%s` (" +
		"`id` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`project_id` INTEGER NOT NULL," +
		"`open` INTEGER DEFAULT 0 NOT NULL," +
		"`closed` INTEGER DEFAULT 0 NOT NULL," +
		"`new` INTEGER DEFAULT 0 NOT NULL," +
		"`created_at` DATE NOT NULL," +
		"FOREIGN KEY (project_id) REFERENCES project(`id`)" +
	");"
const statsTableIndex = "CREATE UNIQUE INDEX IF NOT EXISTS stats_project_created_idx ON %s (`project_id`, `created_at`);"
const statsTableInsert =
	"INSERT INTO %s (" +
		"`project_id`, `open`, `closed`, `new`, `created_at`" +
	") VALUES (?, ?, ?, ?, ?)"
const statsTableUpdate =
	"UPDATE %s SET " +
		"``project_id` = ?, `open` = ?, `closed` = ?, `new` = ?, `created_at` = ? " +
	"WHERE `id` = ?"

func NewStatsTable(statement *Statement) *StatsTable {
	s := new(StatsTable)
	s.statement = statement
	return s
}

func (s *StatsTable) Truncate() (err error) {
	_, err = s.statement.execute(TruncateTable, statsTableName)
	return
}

func (s *StatsTable) CreateStructure() (err error) {
	log.Debugf("Create structure for %s", statsTableName)

	// create table
	_, err = s.statement.execute(statsTableStructure, statsTableName)
	if err != nil {
		return
	}

	// create index
	_, err = s.statement.execute(statsTableIndex, statsTableName)
	return
}

func (s *StatsTable) Select(where string, args ...interface{}) (rows *sql.Rows, err error){
	rows, err = s.SelectComplex(where, "", "", "", args...)
	return
}

func (s *StatsTable) SelectComplex(
	where string,
	order string,
	fields string,
	group string,
	args ...interface{}) (rows *sql.Rows, err error){

	rows, err = s.statement.doSelect(statsTableName, where, order, fields, group, args...)
	return
}

func (s *StatsTable) Insert(
	projectId int,
	open int,
	closed int,
	new int,
	createdAt string) (id int, err error) {

	res, err := s.statement.execute(
		statsTableInsert,
		statsTableName,
		projectId,
		open,
		closed,
		new,
		createdAt)
	if err != nil {
		return
	}

	id64, err := res.LastInsertId()
	id = int(id64)

	return
}

func (s *StatsTable) Update(
	id int,
	projectId int,
	open int,
	closed int,
	new int,
	createdAt string) (rowsAffected int, err error) {

	res, err := s.statement.execute(
		statsTableUpdate,
		statsTableName,
		projectId,
		open,
		closed,
		new,
		createdAt,
		id)
	if err != nil {
		return
	}

	rowsAffected64, err := res.RowsAffected()
	rowsAffected = int(rowsAffected64)
	return
}
