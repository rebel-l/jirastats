package response

type PieChart struct {
	Name string `json:"name"`
	Y float64 `json:"y"`
	Sliced bool `json:"sliced"`
	Selected bool `json:"selected"`
}

func NewPieChart(name string, y float64) *PieChart {
	pc := new(PieChart)
	pc.Name = name
	pc.Y = y
	return pc
}
