package commands

import (
	"database/sql"
	"io/ioutil"
)

const SQL_SETUP_SCRIPT = "./setup/setup.sql"

type DatabaseStructure struct {
	db *sql.DB
}

// NewDatabaseStructure returns a new DatabaseStructure struct
func NewDatabaseStructure(db *sql.DB) *DatabaseStructure {
	dbs := new(DatabaseStructure)
	dbs.db = db
	return dbs
}

// Execute creates the database structure
func (dbs *DatabaseStructure) Execute () (err error) {
	statements, err := ioutil.ReadFile(SQL_SETUP_SCRIPT)
	if err != nil {
		return
	}

	stmt, err := dbs.db.Prepare(string(statements))
	if err != nil {
		return
	}

	_, err = stmt.Exec()
	return
}
