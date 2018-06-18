package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/rebel-l/jirastats/packages/jira/search"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
	"github.com/rebel-l/jirastats/packages/utils"
)

const method = "POST"
const searchEndpoint = "/rest/api/2/search"
const retryMax = 3
const retryWait = 200

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

	for i := 0; i < retryMax; i++ {
		req, _ := s.client.NewRequest(method, searchEndpoint, s.Request)
		res, err := s.client.Do(req, searchResponse)
		if err != nil && res != nil && res.StatusCode == 500 {
			if i < retryMax {
				log.Warn(fmt.Sprintf("Jira responded with an error (try: %d): %s", i, err.Error()))
				wt := time.Duration(rand.Intn(retryWait))
				time.Sleep(wt * time.Millisecond)
				continue
			} else {
				log.Error(fmt.Sprintf("Unrecoverable error connecting Jira (retry limit reached): %s", err))
				return result, err
			}
		} else if err != nil {
			resStr, _ := utils.IoRCTS(res.Body)
			log.Error(fmt.Sprintf("Unrecoverable error connecting Jira: %s", err))
			log.Debug(fmt.Sprintf("Unrecoverable Error - Debug Respose: %s", resStr))
			return result, err
		}

		s.total = searchResponse.Total
		result = searchResponse.Issues
		return result, err
	}
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
