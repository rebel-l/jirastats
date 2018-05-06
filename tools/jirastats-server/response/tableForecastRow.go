package response

type TableForecastRow struct {
	Name                string  `json:"name"`
	AverageSpeedPerDay  float32 `json:"average_speed_per_day"`
	AverageSpeedPerWeek float32 `json:"average_speed_per_week"`
}

func NewTableForcastRow(name string) *TableForecastRow {
	tfr := new(TableForecastRow)
	tfr.Name = name
	return tfr
}

func (tfr *TableForecastRow) SetAverageSpeedPerDay(speed float32) {
	tfr.AverageSpeedPerDay = speed
	tfr.AverageSpeedPerWeek = speed * 5
}
