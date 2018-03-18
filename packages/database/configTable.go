package database

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

const ConfigTableName = "config"
const ConfigTableStructure =
	"CREATE TABLE IF NOT EXISTS `%s` (" +
		"`id` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`name` CHAR(50) NOT NULL," +
		"`config_group_id` INTEGER NOT NULL," +
		"`value` VARCHAR(255)," +
		"FOREIGN KEY(config_group_id) REFERENCES config_group(`id`)" +
");"
const ConfigTableIndex = "CREATE UNIQUE INDEX IF NOT EXISTS config_name_idx ON %s (`name`);"

type ConfigTable struct {
	db *sql.DB
}

func NewConfigTable(db *sql.DB) *ConfigTable {
	c := new(ConfigTable)
	c.db = db
	return c
}

func (c *ConfigTable) Truncate() error {
	truncateNotImplemented(ConfigTableName)
	return nil
}

func (c *ConfigTable) CreateStructure() (err error) {
	log.Debugf("Create structure for %s", ConfigTableName)
	// create table
	err = executeStatement(c.db, c.getCreateTableStatement())
	if err != nil {
		return
	}

	// create index
	err = executeStatement(c.db, c.getCreateIndexStatement())
	return
}

func (c *ConfigTable) getCreateTableStatement() string {
	return createDatabseStatement(ConfigTableStructure, ConfigTableName)
}

func (c *ConfigTable) getCreateIndexStatement() string {
	return createDatabseStatement(ConfigTableIndex, ConfigTableName)
}
