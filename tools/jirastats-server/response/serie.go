package response

type Serie struct {
	Name string `json:"name"`
	Data []float32 `json:"data"`
}

func (s *Serie) AddDataInt(data int) {
	s.Data = append(s.Data, float32(data))
}

func (s *Serie) AddData(data float32) {
	s.Data = append(s.Data, data)
}
