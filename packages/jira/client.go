package jira

import (
	"github.com/andygrunwald/go-jira"
)

func NewClient(config *Config) *jira.Client {
	c, _ := jira.NewClient(nil, config.GetHost())
	c.Authentication.SetBasicAuth(config.GetUsername(), config.GetPassword())
	return c
}
