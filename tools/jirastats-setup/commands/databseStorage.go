package commands

import (
	"github.com/rebel-l/jirastats/packages/utils"
	"github.com/rebel-l/jirastats/packages/database"
)

const SQL_SETUP_SCRIPT = "./setup/setup.sql"

type DatabaseStorage struct {}

//Execute creates the path and file for database storage
func (dbs *DatabaseStorage) Execute() (err error) {
	err = utils.CreateDirectoryIfNotExists(database.DefaultStoragePath)
	if err != nil {
		return
	}

	err = utils.CreateFileIfNotExists(database.DefaultStorageFile)
	return
}
