package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/jira/search"
	log "github.com/sirupsen/logrus"
)

const method = "POST"
const searchEndpoint = "/rest/api/2/search"

type Search struct {
	client *jira.Client
}

func NewSearch(client *jira.Client) *Search {
	search := new(Search)
	search.client = client
	return search
}

func (s *Search) Do(jql string) (result []jira.Issue, err error) {
	searchRequest := search.NewRequest(jql)

	searchResponse := new(search.Response)

	req, _ := s.client.NewRequest(method, searchEndpoint, searchRequest)
	_, err = s.client.Do(req, searchResponse)
	if err != nil {
		log.Error(err)
		return
	}

	//log.Debugf("Total: %d", searchResponse.Total)
	//for _,v := range searchResponse.Issues {
	//	log.Debugf("Issue: %s", v.Key)
	//}

	result = searchResponse.Issues
	return
}
