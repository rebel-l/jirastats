package response

type Serie struct {
	Name string `json:"name"`
	Data []int `json:"data"`
}

func (s *Serie) AddData(data int) {
	s.Data = append(s.Data, data)
}
