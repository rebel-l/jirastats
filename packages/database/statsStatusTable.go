package database

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

type StatsStatusTable struct {
	db *sql.DB
}

const statsStatusTableName = "stats_status"
const statsStatusTableStructure =
	"CREATE TABLE IF NOT EXISTS `%s` (" +
		"`id` INTEGER PRIMARY KEY," +
		"`status` CHAR(50)" +
	");"
const statsStatusTableIndex = "CREATE UNIQUE INDEX IF NOT EXISTS statsstatus_status_idx ON %s (`status`);"

func NewSatsStatusTable(db *sql.DB) *StatsStatusTable {
	ss := new(StatsStatusTable)
	ss.db = db
	return ss
}

func (ss *StatsStatusTable) Truncate() error {
	truncateNotImplemented(statsStatusTableName)
	return nil
}

func (ss *StatsStatusTable) CreateStructure() (err error) {
	log.Debugf("Create structure for %s", statsStatusTableName)
	// create table
	err = executeStatement(ss.db, ss.getCreateTableStatement())
	if err != nil {
		return
	}

	// create index
	err = executeStatement(ss.db, ss.getCreateIndexStatement())
	return
}

func (ss *StatsStatusTable) getCreateTableStatement() string {
	return createDatabseStatement(statsStatusTableStructure, statsStatusTableName)
}

func (ss *StatsStatusTable) getCreateIndexStatement() string {
	return createDatabseStatement(statsStatusTableIndex, statsStatusTableName)
}
