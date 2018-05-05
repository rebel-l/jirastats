package response

import (
	"github.com/rebel-l/jirastats/packages/models"
	"strings"
)

type StatsPiechartTable struct {
	Project *models.Project `json:"project"`
	Stats map[string]*PieChartTable `json:"stats"`
}

func NewStatsPiechartTable(project *models.Project) *StatsPiechartTable {
	spt := new(StatsPiechartTable)
	spt.Project = project
	spt.Stats = make(map[string]*PieChartTable)
	return spt
}

func (spt *StatsPiechartTable) GeneratePiechartTables(tables map[string]map[string]*TableData, countTickets int) {
	for k, v := range tables {
		displayName := strings.Title(strings.Replace(k, "_", " ", len(k)))
		finalData := make([]*PieChartEntry, len(v))
		spt.Stats[k] = spt.createPieChart(displayName, finalData, v, countTickets)
		// TODO: parallize processing data
	}
}

func (spt *StatsPiechartTable) createPieChart(
	name string,
	chart []*PieChartEntry,
	data map[string]*TableData,
	counter int) *PieChartTable {

	pct := NewPieChartTable(name, chart)
	i := 0
	maxValue := 0
	maxItem := 0
	for _, d := range data {
		if d.Value > maxValue {
			maxValue = d.Value
			maxItem = i
		}
		chart[i] = NewPieChartEntry(d.Name, float64(d.Value) * 100.0 / float64(counter))
		pct.DataTable = append(pct.DataTable, d)
		i++
	}
	chart[maxItem].Sliced = true
	chart[maxItem].Selected = true

	return pct
}
