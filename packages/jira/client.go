package jira

import (
	"github.com/andygrunwald/go-jira"
	"net/http"
)

type Client interface {
	NewRequest(method, urlStr string, body interface{}) (*http.Request, error)
	Do(req *http.Request, v interface{}) (*jira.Response, error)
}

func NewClient(config *Config) *jira.Client {
	c, _ := jira.NewClient(nil, config.GetHost())
	c.Authentication.SetBasicAuth(config.GetUsername(), config.GetPassword())
	return c
}
