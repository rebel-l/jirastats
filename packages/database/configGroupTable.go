package database

import (
	"database/sql"
)

const ConfigGroupTableName = "config_group"
const ConfigGroupTableStructure =
	"CREATE TABLE IF NOT EXISTS `%s` (" +
		"`id` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`name` CHAR(50) NOT NULL" +
	");"
const ConfigGroupTableIndex = "CREATE UNIQUE INDEX IF NOT EXISTS config_group_name_idx ON %s (`name`);"

type ConfigGroupTable struct {
	db *sql.DB
}

func NewConfigGroupTable(db *sql.DB) *ConfigGroupTable {
	cg := new(ConfigGroupTable)
	cg.db = db
	return cg
}

func (cg *ConfigGroupTable) CreateStructure() (err error) {
	// create table
	err = executeStatement(cg.db, cg.getCreateTableStatement())
	if err != nil {
		return
	}

	// create index
	err = executeStatement(cg.db, cg.getCreateIndexStatement())
	return
}

func (cg *ConfigGroupTable) getCreateTableStatement() string {
	return createDatabseStatement(ConfigGroupTableStructure, ConfigGroupTableName)
}

func (cg *ConfigGroupTable) getCreateIndexStatement() string {
	return createDatabseStatement(ConfigGroupTableIndex, ConfigGroupTableName)
}

func (cg *ConfigGroupTable) Truncate() error {
	truncateNotImplemented(ConfigGroupTableName)
	return nil
}
