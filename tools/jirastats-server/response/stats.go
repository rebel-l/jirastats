package response

type Stats struct {
	ProjectId int `json:"project_id"`
	ProjectName string `json:"project_name"`
	Categories []string `json:"categories"`
	Series []*Serie `json:"series"`
}

func (s *Stats) AddCategory(name string) {
	s.Categories = append(s.Categories, name)
}
