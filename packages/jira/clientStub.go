package jira

import (
	"errors"
	"net/http"

	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/jira/search"
	"github.com/rebel-l/jirastats/packages/utils/io"
	log "github.com/sirupsen/logrus"
)

// ClientStub is a stub to simulate the API of Jira.
type ClientStub struct {
	AllowEmptyResponse bool
	issues map[string][]jira.Issue
}

// NewClientStub return a new client stub struct
func NewClientStub() *ClientStub {
	cs := new(ClientStub)
	cs.issues = make(map[string][]jira.Issue)
	cs.AllowEmptyResponse = true
	return cs
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

	jRes = new(jira.Response)
	jRes.Response = new(http.Response)
	value, ok := v.(*search.Response)
	if ok == false {
		err = errors.New("not a search response injected")
		jRes.Response.StatusCode = 404
		return
	}

	is, ok := cs.issues[sReq.Jql]
	if (ok == false || len(is) < 1) && cs.AllowEmptyResponse == false {
		err = errors.New("no data injected")
		log.Errorf("jRes: %#v", jRes)
		jRes.Response.StatusCode = 404
		return
	}

	value.Total = len(is)
	end := sReq.StartAt + sReq.MaxResults
	if end > value.Total {
		end = value.Total
	}
	value.Issues = is[sReq.StartAt:end]

	return
}

// AddIssue add Jira issues to a simulated response
func (cs *ClientStub) AddIssue(issue jira.Issue, jql string) {
	_, ok := cs.issues[jql]
	if ok == false {
		cs.issues[jql] = make([]jira.Issue, 0)
	}
	cs.issues[jql] = append(cs.issues[jql], issue)
}
