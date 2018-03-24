package database

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Statement struct {
	db *sql.DB
}

func NewStatement(db *sql.DB) *Statement {
	s := new(Statement)
	s.db = db
	return s
}

func (s *Statement) doSelect(
	table string,
	where string,
	order string,
	fields string,
	group string,
	args ...interface{}) (rows *sql.Rows, err error) {

	if fields == "" {
		fields = "*"
	}

	if where != "" {
		where = fmt.Sprintf("WHERE %s", where)
	}

	if group != "" {
		group = fmt.Sprintf("GROUP BY %s", group)
	}

	if order != "" {
		order = fmt.Sprintf("ORDER BY %s", order)
	}

	statement := fmt.Sprintf(SelectStatement, fields, table, where, group, order)
	stmt, err := s.db.Prepare(statement)
	if err != nil {
		return
	}

	if args != nil {
		rows, err = stmt.Query(args...)
	} else {
		rows, err = stmt.Query()
	}
	return
}

func (s *Statement) execute(statement string, table string, args ...interface{}) (res sql.Result, err error) {
	stmt, err := s.db.Prepare(createDatabseStatement(statement, table))
	if err != nil {
		return
	}

	if args != nil {
		res, err = stmt.Exec(args...)
	} else {
		res, err = stmt.Exec()
	}
	return
}

func executeStatement(db *sql.DB, statement string) error {
	// TODO: deprecated, use statement struct instead
	s := NewStatement(db)
	_, err := s.execute(statement, "")
	return err
}

func createDatabseStatement(statement string, tablename string) string {
	return fmt.Sprintf(statement, tablename)
}

func truncateNotImplemented (table string) {
	log.Infof("Truncate not implemented for table %s", table)
}
