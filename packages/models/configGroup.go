package models

type ConfigGroup struct {
	Id int
	Name string
	Configs map[string]*Config
}

func NewConfigGroup() *ConfigGroup {
	cg := new(ConfigGroup)
	cg.Configs = make(map[string]*Config)
	return cg
}
