package response

import "time"

type StatsForecastEnd struct {
	Days int `json:"days"`
	Day time.Time `json:"day"`
	Weeks int `json:"weeks"`
	Week string `json:"week"`
}
