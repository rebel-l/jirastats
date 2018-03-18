package database

import (
	"database/sql"
	"github.com/rebel-l/jirastats/packages/models"
	log "github.com/sirupsen/logrus"
)

type ConfigGroupMapper struct {
	table *ConfigGroupTable
}

func NewConfigGroupMapper(db *sql.DB) *ConfigGroupMapper {
	cgm := new(ConfigGroupMapper)
	cgm.table = NewConfigGroupTable(db)
	return cgm
}

func (cgm *ConfigGroupMapper) LoadByName(name string, model *models.ConfigGroup) (err error) {
	rows, err := cgm.table.Select("`name` = ?", name)
	if err != nil {
		return
	}

	if rows.Next() {
		err = rows.Scan(&model.Id, &model.Name)
		if err != nil {
			log.Warn("Mapping of config group not possible")
		}
	}
	return
}

func (cgm *ConfigGroupMapper) Save(model *models.ConfigGroup) (err error) {
	if model.Id == 0 {
		// insert
		id, err := cgm.table.Insert(model.Name)
		if err != nil {
			return err
		}
		model.Id = id
	} else {
		// update
		log.Warn("Update for ConfigGroup not implemented")
	}
	return
}