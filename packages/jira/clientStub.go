package jira

import (
	"errors"
	"net/http"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/jira/search"
	"github.com/rebel-l/jirastats/packages/utils/io"
	log "github.com/sirupsen/logrus"
)

// ClientStub is a stub to simulate the API of Jira.
type ClientStub struct {
	issues []jira.Issue // TODO: split by Request (jql)
}

// NewRequest returns a new request
func (cs *ClientStub) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	c, err := jira.NewClient(new(http.Client), "")
	if err != nil {
		return nil, err
	}

	r, err := c.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Do simulates a request against the Jira API and returns a response depending on the request
func (cs *ClientStub) Do(req *http.Request, v interface{}) (jRes *jira.Response, err error) {
	body, err := io.ReadCloserToByte(req.Body)
	if err != nil {
		return
	}

	sReq, err := search.NewRequestFromJson(body)
	if err != nil {
		return
	}
	log.Debugf("Body: %s", sReq.Jql)

	jRes = new(jira.Response)
	value, ok := v.(*search.Response)
	if ok == false {
		err = errors.New("not a search response injected")
		return
	}

	if len(cs.issues) < 1 {
		err = errors.New("no data injected")
		return
	}

	value.Total = len(cs.issues)
	value.Issues = cs.issues

	return
}

// AddIssue add Jira issues to a simulated response
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
