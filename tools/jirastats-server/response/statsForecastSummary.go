package response

import "math"

type StatsForecastSummary struct {
	Days int `json:"days"`
	LastDay string `json:"last_day"`
	Weeks int `json:"weeks"`
	LastWeek string `json:"last_week"`
}

func NewStatsForecastSummary() *StatsForecastSummary {
	sfs := new(StatsForecastSummary)
	sfs.Days = -1
	sfs.Weeks = -1
	return sfs
}

func (sfs *StatsForecastSummary) SetDaysAndWeeks(days int) {
	sfs.Days = days
	weeks := float64(days) / 5.0
	sfs.Weeks = int(math.Ceil(weeks))
}
