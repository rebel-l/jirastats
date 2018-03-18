package commands

import (
	"database/sql"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/models"
	"fmt"
)

type ConfigureApplication struct {
	mapper *database.ConfigGroupMapper
}

func NewConfigureApplication(db *sql.DB) *ConfigureApplication {
	ca := new(ConfigureApplication)
	ca.mapper = database.NewConfigGroupMapper(db)
	return ca
}

func (ca *ConfigureApplication) Execute() (err error) {
	cg := new(models.ConfigGroup)
	cg.Name = "jira" // TODO put to constant somewhere
	err = ca.mapper.Save(cg)
	// fmt.Printf("ConfigurationGroup saved: Id: %d, Name %s", cg.Id, cg.Name)
	return
}
