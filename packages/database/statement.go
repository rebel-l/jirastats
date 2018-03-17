package database

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func executeStatement(db *sql.DB, statement string) error {
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	return err
}

func createDatabseStatement(statement string, tablename string) string {
	return fmt.Sprintf(statement, tablename)
}

func truncateNotImplemented (table string) {
	log.Infof("Truncate not implemented for table %s", table)
}
