package search

import (
	"github.com/andygrunwald/go-jira"
)

type Response struct {
	Total int `json:"total"`
	Issues []jira.Issue `json:"issues"`
}
