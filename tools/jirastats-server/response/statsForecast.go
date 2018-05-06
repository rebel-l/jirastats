package response

import (
	"github.com/rebel-l/jirastats/packages/models"
	"errors"
	"fmt"
)

type StatsForecast struct {
	Project *models.Project `json:"project"`
	Speed []*TableForecastRow `json:"speed"`
	Summary *StatsForecastEnd `json:"summary"`
	Chart *Stats `json:"chart"`
}

func NewStatsForecast(project *models.Project) *StatsForecast {
	sf := new(StatsForecast)
	sf.Project = project
	sf.Chart = new(Stats)
	sf.Chart.ProjectId = project.Id
	sf.Chart.ProjectName = project.Name
	return sf
}

func (sf *StatsForecast) Calculate(data []*models.Stats) {
	count := len(data)
	if count < 6 {
		// default only
		sf.calcDefaultSpeed(data)
	} else if count < 32 {
		// one month only
	} else if count < 94 {
		// one month & 3 months
	} else {
		// one month, 3 months & all
	}
}

func (sf *StatsForecast) addSpeed(speed *TableForecastRow) {
	sf.Speed = append(sf.Speed, speed)
}

func (sf *StatsForecast) calcDefaultSpeed(data []*models.Stats) {
	newTickets := 0
	for _, v := range data {
		newTickets += v.New
	}

	newTicketsPerDay := float32(newTickets) / float32(len(data))
	speedWithNew := NewTableForcastRow("default speed including new")
	speedWithNew.SetAverageSpeedPerDay(sf.Project.KnownSpeed - newTicketsPerDay)
	sf.addSpeed(speedWithNew)

	speedWithoutNew := NewTableForcastRow("default speed without new")
	speedWithoutNew.SetAverageSpeedPerDay(sf.Project.KnownSpeed)
	sf.addSpeed(speedWithoutNew)
}

func (sf *StatsForecast) calcSpeed(name string, data []*models.Stats, start int, offset int) (err error) {
	max := len(data)
	if start >= max {
		return errors.New(fmt.Sprintf("Start (%d) is higher then max. entries (%d)", start, max))
	}

	if start < 0 {
		start = 0
	}

	for i := start; i < max; i++ {
		if i > offset {
			break
		}


	}

	return
}

/*
	Chart:
		speed last month (no new) ==> defaultSpeed if less than one week
		speed last month (with new) ==> defaultSpeed if less than one week
		speed last 3 months (no new) ==> skip if less than one week
		speed last 3 months (with new) ==> skip if less than one week
		speed over all (no new) ==> skip if less than one week
		speed over all (with new) ==> skip if less than one week
	Table:
		defaultSpeed with new if less than one week: daily / weekly
		defaultSpeed without new if less than one week: daily / weekly

		speed one month with new if more than one week: daily / weekly
		speed one month without new if more than one week: daily / weekly

		speed 3 months with new if more than one week: daily / weekly
		speed 3 months without new if more than one week: daily / weekly

		speed over all with new if more than one week: daily / weekly
		speed over all without new if more than one week: daily / weekly

		guessed project end: days (date) / weeks (week/year)

 */