package database

import (
	"database/sql"
)

type Table interface {
	CreateStructure() error
	Select(where string, args ...interface{}) (rows *sql.Rows, err error)
	Truncate() error
	// TODO: SelectComplex(where string, order string, fields string, group string, args ...interface{})
}
