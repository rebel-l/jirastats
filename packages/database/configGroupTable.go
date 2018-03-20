package database

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

const configGroupTableName = "config_group"
const configGroupTableStructure =
	"CREATE TABLE IF NOT EXISTS `%s` (" +
		"`id` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`name` CHAR(50) NOT NULL" +
	");"
const configGroupTableIndex = "CREATE UNIQUE INDEX IF NOT EXISTS config_group_name_idx ON %s (`name`);"
const configGroupInsert = "INSERT INTO %s(`name`) values(?)"

type ConfigGroupTable struct {
	db *sql.DB
}

func NewConfigGroupTable(db *sql.DB) *ConfigGroupTable {
	cg := new(ConfigGroupTable)
	cg.db = db
	return cg
}

func (cg *ConfigGroupTable) CreateStructure() (err error) {
	log.Debugf("Create structure for %s", configGroupTableName)
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
	return createDatabseStatement(configGroupTableStructure, configGroupTableName)
}

func (cg *ConfigGroupTable) getCreateIndexStatement() string {
	return createDatabseStatement(configGroupTableIndex, configGroupTableName)
}

func (cg *ConfigGroupTable) Truncate() error {
	truncateNotImplemented(configGroupTableName)
	return nil
}

func (cg *ConfigGroupTable) Select(where string, args ...interface{}) (rows *sql.Rows, err error){
	statement := cg.getSelectAllStatement()
	if where != "" {
		statement += " WHERE " + where
	}

	stmt, err := cg.db.Prepare(statement)
	if err != nil {
		return
	}

	if args != nil {
		rows, err = stmt.Query(args)
	} else {
		rows, err = stmt.Query()
	}
	return
}

func (cg *ConfigGroupTable) Insert(name string) (id int, err error) {
	stmt, err := cg.db.Prepare(createDatabseStatement(configGroupInsert, configGroupTableName))
	if err != nil {
		return
	}

	res, err := stmt.Exec(name)
	if err != nil {
		return
	}

	id64, err := res.LastInsertId()
	id = int(id64)

	return
}

func (cg *ConfigGroupTable) getSelectAllStatement() string {
	return createDatabseStatement(SelectAllStatement, configGroupTableName)
}
