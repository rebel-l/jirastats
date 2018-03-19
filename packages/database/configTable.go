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
const ConfigInsert = "INSERT INTO %s(`name`, `value`, `config_group_id`) values(?, ?, ?)"
const ConfigUpdate = "UPDATE %s SET `name`=?, `value`=? where id=?"

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

func (c *ConfigTable) Insert(name string, value string, configGroupId int) (id int, err error) {
	stmt, err := c.db.Prepare(createDatabseStatement(ConfigInsert, ConfigTableName))
	if err != nil {
		return
	}

	res, err := stmt.Exec(name, value, configGroupId)
	if err != nil {
		return
	}

	id64, err := res.LastInsertId()
	id = int(id64)

	return
}

func (c *ConfigTable) Update(id int, name string, value string) (err error) {
	stmt, err := c.db.Prepare(createDatabseStatement(ConfigUpdate, ConfigTableName))
	if err != nil {
		return
	}

	_, err = stmt.Exec(name, value, id)
	return
}

func (c *ConfigTable) Select(where string, args interface{}) (rows *sql.Rows, err error){
	statement := c.getSelectAllStatement()
	if where != "" {
		statement += " WHERE " + where
	}

	stmt, err := c.db.Prepare(statement)
	if err != nil {
		return
	}

	rows, err = stmt.Query(args)
	return
}

func (c *ConfigTable) getSelectAllStatement() string {
	return createDatabseStatement(SelectAllStatement, ConfigTableName)
}
