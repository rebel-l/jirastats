package response

type PieChartEntry struct {
	Name string `json:"name"`
	Y float64 `json:"y"`
	Sliced bool `json:"sliced"`
	Selected bool `json:"selected"`
}

func NewPieChartEntry(name string, y float64) *PieChartEntry {
	pc := new(PieChartEntry)
	pc.Name = name
	pc.Y = y
	return pc
}
