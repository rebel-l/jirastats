package jira

import (
	"errors"
	"net/http"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/jira/search"
)

type ClientStub struct {
	issues []jira.Issue // TODO: split by Request (jql)
}

func (cs *ClientStub) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	return new(http.Request), nil
}

func (cs *ClientStub) Do(req *http.Request, v interface{}) (*jira.Response, error) {
	jRes := new(jira.Response)
	value, ok := v.(*search.Response)
	if ok == false {
		return nil, errors.New("not a search response injected")
	}

	if len(cs.issues) < 1 {
		return nil, errors.New("no data injected")
	}

	value.Total = len(cs.issues)
	value.Issues = cs.issues

	return jRes, nil
}

func (cs *ClientStub) AddIssue(key string, summary string, status string, priority string, iType string, components []string, labels []string, created time.Time, updated time.Time) {
	jPriority := new(jira.Priority)
	jPriority.Name = priority

	jStatus := new(jira.Status)
	jStatus.Name = status

	issueFields := new(jira.IssueFields)
	issueFields.Summary = summary
	issueFields.Type = jira.IssueType{Name: iType}
	issueFields.Status = jStatus
	issueFields.Priority = jPriority
	issueFields.Components = cs.createComponents(components)
	issueFields.Labels = labels
	issueFields.Created = created.Format(JiraDateTimeFormat)
	issueFields.Updated = updated.Format(JiraDateTimeFormat)

	issue := jira.Issue{ID: key, Key: key, Fields: issueFields}
	cs.issues = append(cs.issues, issue)
}

func (cs *ClientStub) createComponents(components []string) []*jira.Component {
	c := make([]*jira.Component, 0)
	for _, v := range components {
		nc := new(jira.Component)
		nc.Name = v
		c = append(c, nc)
	}
	return c
}
