package commands

import (
	"database/sql"
	"github.com/rebel-l/jirastats/packages/database"
)

type CreateDatabaseStructure struct {
	tables []database.Table
}

// NewCreateDatabaseStructure returns a new CreateDatabaseStructure struct
func NewCreateDatabaseStructure(db *sql.DB) *CreateDatabaseStructure {
	dbs := new(CreateDatabaseStructure)
	dbs.tables = append(dbs.tables, database.NewConfigGroupTable(db))
	dbs.tables = append(dbs.tables, database.NewConfigTable(db))
	dbs.tables = append(dbs.tables, database.NewProjectTable(db))
	dbs.tables = append(dbs.tables, database.NewTicketTable(db))
	dbs.tables = append(dbs.tables, database.NewSatsStatusTable(db))
	dbs.tables = append(dbs.tables, database.NewSatsTable(db))
	return dbs
}

// Execute creates the database structure
func (dbs *CreateDatabaseStructure) Execute () (err error) {
	for _, v := range dbs.tables {
		err = v.CreateStructure()
		if err != nil {
			break
		}
	}
	return
}
