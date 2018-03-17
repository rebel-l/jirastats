package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// GetDbConncection returns a db connection
func GetDbConnection() (db *sql.DB, err error) {
	db, err = sql.Open(DefaultDriver, DefaultStorageFile)
	return
}
