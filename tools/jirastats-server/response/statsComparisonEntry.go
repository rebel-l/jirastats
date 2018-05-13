package response

import (
	"strings"
)

type StatsComparisonEntry struct {
	Name string`json:"name"`
	Left []*PieChartEntry `json:"left"`
	Right []*PieChartEntry `json:"right"`
	Diff []*TableDiffEntry `json:"diff"`
}

func NewStatsComparisonEntry(name string) *StatsComparisonEntry {
	sce := new(StatsComparisonEntry)
	sce.Name = strings.Title(strings.Replace(name, "_", " ", len(name)))
	return sce
}
