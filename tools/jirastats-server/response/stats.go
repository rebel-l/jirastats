package response

import (
	"fmt"
	"github.com/rebel-l/jirastats/packages/database"
	"net/http"
)

type Stats struct {
	ProjectId int `json:"project_id"`
	ProjectName string `json:"project_name"`
	Categories []string `json:"categories"`
	Series []*Serie `json:"series"`
	res http.ResponseWriter
	pm *database.ProjectMapper
}

func NewStats(projectId int, pm *database.ProjectMapper, res http.ResponseWriter) (s *Stats, ok bool) {
	s = new(Stats)
	s.ProjectId = projectId
	s.pm = pm
	s.res = res
	ok = s.setProjectName()
	return
}

func (s *Stats) AddCategory(name string) {
	s.Categories = append(s.Categories, name)
}

func (s *Stats) setProjectName() bool {
	// TODO: maybe can be set by client, not necessary to send it again
	project, err := s.pm.LoadProjectById(s.ProjectId)
	if err != nil {
		msg := fmt.Sprintf("Not able to load project id %d: %s", s.ProjectId, err.Error())
		e := NewErrorJson(msg, s.res)
		e.SendInternalServerError()
		return false
	}

	if project == nil {
		msg := fmt.Sprintf("No project found for id: %d", s.ProjectId)
		e := NewErrorJson(msg, s.res)
		e.SendNotFound()
		return false
	}

	s.ProjectName = project.Name
	return true
}

func (s *Stats) AddSerie(serie *Serie) {
	s.Series = append(s.Series, serie)
}
