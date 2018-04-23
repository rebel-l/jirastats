package response

type PieChartTable struct {
	Name string            `json:"name"`
	DataTable []*TableData `json:"data_table"`
	DataChart []*PieChart  `json:"data_chart"`
}

func NewPieChartTable(name string, chartData []*PieChart) *PieChartTable {
	pct := new(PieChartTable)
	pct.Name = name
	pct.DataChart = chartData
	return pct
}
