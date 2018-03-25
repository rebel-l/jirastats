package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/jira/search"
	log "github.com/sirupsen/logrus"
)

const method = "POST"
const searchEndpoint = "/rest/api/2/search"

type Search struct {
	client  *jira.Client
	Request *search.Request
	total   int
}

func NewSearch(client *jira.Client, jql string) *Search {
	s := new(Search)
	s.client = client
	s.Request = search.NewRequest(jql)
	return s
}

func (s *Search) Do() (result []jira.Issue, err error) {

	searchResponse := new(search.Response)

	req, _ := s.client.NewRequest(method, searchEndpoint, s.Request)
	_, err = s.client.Do(req, searchResponse)
	if err != nil {
		log.Error(err)
		return
	}

	s.total = searchResponse.Total
	result = searchResponse.Issues
	return
}

func (s *Search) Next() bool {
	/* EXAMPLES:
		1: t ==> 0, start ==> 0, max ==> 10
		2: t ==> 5, start ==> 10, max ==> 10 ==> exit


		1: t ==> 0, start ==> 0, max ==> 10
		2: t ==> 15, start ==> 10, max ==> 10 ==> continue
		3: t ==> 15, start ==> 20, max ==> 10 ==> exit

		1: t ==> 0, start ==> 0, max ==> 10
		2: t ==> 20, start ==> 10, max ==> 10 ==> continue
		3: t ==> 20, start ==> 20, max ==> 10 ==> exit

		1: t ==> 0, start ==> 0, max ==> 10
		2: t ==> 21, start ==> 10, max ==> 10 ==> continue
		3: t ==> 21, start ==> 20, max ==> 10 ==> continue
		4: t ==> 21, start ==> 30, max ==> 10 ==> exit
	*/
	s.Request.Next()
	return s.total > s.Request.StartAt
}
