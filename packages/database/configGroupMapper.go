package database

import (
	"database/sql"
	"github.com/rebel-l/jirastats/packages/models"
	log "github.com/sirupsen/logrus"
)

type ConfigGroupMapper struct {
	table *ConfigGroupTable
	configMapper *ConfigMapper
}

func NewConfigGroupMapper(db *sql.DB) *ConfigGroupMapper {
	cgm := new(ConfigGroupMapper)
	cgm.table = NewConfigGroupTable(db)
	cgm.configMapper = NewConfigMapper(db)
	return cgm
}

func (cgm *ConfigGroupMapper) LoadByName(name string, model *models.ConfigGroup) (err error) {
	rows, err := cgm.table.Select("`name` = ?", name)
	defer rows.Close()
	if err != nil {
		return
	}

	if rows.Next() {
		err = rows.Scan(&model.Id, &model.Name)
		if err != nil {
			log.Warn("Mapping of config group not possible")
		}
	}

	err = cgm.configMapper.LoadByConfigGroupId(model.Id, model.Configs)
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
	}

	for _, v := range model.Configs {
		v.GroupdId = model.Id
		err = cgm.configMapper.Save(v)
		if err != nil {
			return err
		}
	}

	return
}
