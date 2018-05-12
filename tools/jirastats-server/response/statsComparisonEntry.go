package response

type StatsComparisonEntry struct {
	Left []*PieChartEntry `json:"left"`
	Right []*PieChartEntry `json:"right"`
	Diff []*TableData `json:"diff"`
}

func NewStatsComparisonEntry() *StatsComparisonEntry {
	sce := new(StatsComparisonEntry)
	return sce
}
