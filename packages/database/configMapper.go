package database

import (
	"database/sql"
	"github.com/rebel-l/jirastats/packages/models"
	log "github.com/sirupsen/logrus"
)

type ConfigMapper struct {
	table *ConfigTable
}

func NewConfigMapper(db *sql.DB) *ConfigMapper {
	cm := new(ConfigMapper)
	cm.table = NewConfigTable(db)
	return cm
}

func (cm *ConfigMapper) LoadByConfigGroupId(configGroupId int, configs map[string]*models.Config) (err error) {
	rows, err := cm.table.Select("config_group_id = ?", configGroupId)
	defer rows.Close()
	if err != nil {
		return
	}

	var id int
	var name string
	var value string
	var groupId int
	for rows.Next() {
		err = rows.Scan(&id, &name, &groupId, &value)
		if err != nil {
			log.Warn("Mapping of config not possible")
		}

		config, ok := configs[name]
		if ok == false {
			config = new(models.Config)
		}

		config.Id = id
		config.Name = name
		config.Value = value
		config.GroupdId = groupId
		configs[name] = config
	}
	return
}

func (cm *ConfigMapper) Save(model *models.Config) (err error) {
	if model.Id == 0 {
		// insert
		id, err := cm.table.Insert(model.Name, model.Value, model.GroupdId)
		if err != nil {
			return err
		}
		model.Id = id
	} else {
		// update
		err = cm.table.Update(model.Id, model.Name, model.Value)
	}
	return
}
