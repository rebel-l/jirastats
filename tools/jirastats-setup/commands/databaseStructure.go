package commands

import (
	"database/sql"
	"github.com/rebel-l/jirastats/packages/database"
)

type DatabaseStructure struct {
	tables []database.Table
}

// NewDatabaseStructure returns a new DatabaseStructure struct
func NewDatabaseStructure(db *sql.DB) *DatabaseStructure {
	dbs := new(DatabaseStructure)
	dbs.tables = append(dbs.tables, database.NewConfigGroupTable(db))
	dbs.tables = append(dbs.tables, database.NewConfigTable(db)) // TODO: is not created :-(
	dbs.tables = append(dbs.tables, database.NewProjectTable(db))
	dbs.tables = append(dbs.tables, database.NewTicketTable(db))
	return dbs
}

// Execute creates the database structure
func (dbs *DatabaseStructure) Execute () (err error) {
	for _, v := range dbs.tables {
		err = v.CreateStructure()
		if err != nil {
			break
		}
	}
	return
}
