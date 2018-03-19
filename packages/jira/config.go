package jira

import (
	"database/sql"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/models"
	log "github.com/sirupsen/logrus"
)

const ConfigGroupName = "jira"
const configHostName = "host"
const configUsernameName = "username"
const configPasswordName = "password"

type Config struct {
	cgm *database.ConfigGroupMapper
	configGroup *models.ConfigGroup
}

func NewConfig(db *sql.DB) *Config {
	c := new(Config)
	c.cgm = database.NewConfigGroupMapper(db)
	c.configGroup = models.NewConfigGroup()
	c.configGroup.Name = ConfigGroupName
	err := c.load()
	if err != nil {
		log.Errorf("Wasn't able to load Jira config: %s", err.Error())
	}
	return c
}

func (c *Config) GetHost() string {
	return c.getConfig(configHostName)
}

func (c *Config) GetUsername() string {
	return c.getConfig(configUsernameName)
}

func (c *Config) GetPassword() string {
	return c.getConfig(configPasswordName)
}

func (c *Config) SetHost(host string) {
	c.setConfig(configHostName, host)
}

func (c *Config) SetUsername(username string) {
	c.setConfig(configUsernameName, username)
}

func (c *Config) SetPassword(password string) {
	c.setConfig(configPasswordName, password)
}

func (c *Config) getConfig(name string) string {
	v, ok := c.configGroup.Configs[name]
	if ok == false {
		return ""
	}

	return v.Value
}

func (c *Config) setConfig(name string, value string) {
	v, ok := c.configGroup.Configs[name]
	if ok == false {
		config := new(models.Config)
		config.Name = name
		config.Value = value
		config.GroupdId = c.configGroup.Id
		c.configGroup.Configs[name] = config
		return
	}

	v.Value = value
}

func (c *Config) load() (err error) {
	err = c.cgm.LoadByName(ConfigGroupName, c.configGroup)
	return
}

func (c *Config) Save() (err error) {
	err = c.cgm.Save(c.configGroup)
	return
}