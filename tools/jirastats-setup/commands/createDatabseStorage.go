package commands

import (
	"github.com/rebel-l/jirastats/packages/utils"
	"github.com/rebel-l/jirastats/packages/database"
)

type CreateDatabaseStorage struct {}

//Execute creates the path and file for database storage
func (dbs *CreateDatabaseStorage) Execute() (err error) {
	err = utils.CreateDirectoryIfNotExists(database.DefaultStoragePath)
	if err != nil {
		return
	}

	err = utils.CreateFileIfNotExists(database.DefaultStorageFile)
	return
}
