package response

import (
	"github.com/rebel-l/jirastats/packages/models"
	log "github.com/sirupsen/logrus"
	"math"
	"time"
)

type StatsForecast struct {
	Project *models.Project `json:"project"`
	Speed []*TableForecastRow `json:"speed"`
	Summary *StatsForecastEnd `json:"summary"`
	Chart *Stats `json:"chart"`
	maxDaysForecast int
	numberOfOpenTickets int
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
	sf.numberOfOpenTickets = data[count - 1].Open

	if count < 6 {
		// default only
		sf.calcDefaultSpeed(data)
	} else if count < 32 {
		// one month only
		sf.calcSpeed("overall speed", data, 0)
	} else if count < 94 {
		// one month & 3 months
		sf.calcSpeed("overall speed", data, 0)
		sf.calcSpeed("last month speed", data, count - 31)
	} else {
		// one month, 3 months & all
		sf.calcSpeed("overall speed", data, 0)
		sf.calcSpeed("last 3 months speed", data, count - 93)
		sf.calcSpeed("last month speed", data, count - 31)
	}

	sf.calcMaxDaysForecast()
	sf.calcCategories()
	for _, speed := range sf.Speed {
		sf.calcSerie(speed)
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

func (sf *StatsForecast) calcSpeed(name string, data []*models.Stats, start int) {
	max := len(data)
	if start >= max {
		log.Errorf("Start (%d) is higher then max. entries (%d)", start, max)
		return
	}

	if start < 0 {
		start = 0
	}

	numberOfDays := max - start
	numberOfNewTickets := 0
	numberOfClosedTickets := 0
	for i := start; i < max; i++ {
		numberOfNewTickets += data[i].New
		numberOfClosedTickets += data[i].Closed
	}

	newTicketsPerDay := float32(numberOfNewTickets) / float32(numberOfDays)
	closedTicketsPerDay := float32(numberOfClosedTickets) / float32(numberOfDays)

	speedWithNew := NewTableForcastRow(name + " including new")
	speedWithNew.SetAverageSpeedPerDay(closedTicketsPerDay - newTicketsPerDay)
	sf.addSpeed(speedWithNew)

	speedWithoutNew := NewTableForcastRow(name + " without new")
	speedWithoutNew.SetAverageSpeedPerDay(closedTicketsPerDay)
	sf.addSpeed(speedWithoutNew)

	return
}

func (sf *StatsForecast) calcMaxDaysForecast() {
	for _, speed := range sf.Speed {
		if speed.AverageSpeedPerDay < 0 {
			continue
		}

		maxDays := float32(sf.numberOfOpenTickets) / speed.AverageSpeedPerDay
		maxDaysInt := int(math.Ceil(float64(maxDays))) + 1
		if sf.maxDaysForecast < maxDaysInt {
			sf.maxDaysForecast = maxDaysInt
		}
	}
}

func (sf *StatsForecast) calcCategories() {
	day := time.Now()
	for i := sf.maxDaysForecast; i > 0; {
		if day.Weekday().String() == "Saturday" || day.Weekday().String() == "Sunday" {
			day = day.AddDate(0, 0, 1)
			continue
		}

		sf.Chart.AddCategory(day.Format(dateFormat))
		day = day.AddDate(0, 0, 1)
		i--
	}
}

func (sf *StatsForecast) calcSerie(speed *TableForecastRow) {
	remainingTickets := float32(sf.numberOfOpenTickets)
	serie := new(Serie)
	serie.Name = speed.Name
	for i := 0; i < sf.maxDaysForecast; i++ {
		serie.AddData(remainingTickets)

		remainingTickets -= speed.AverageSpeedPerDay
		if remainingTickets < 0 {
			remainingTickets = 0
		}
	}

	sf.Chart.AddSerie(serie)
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