package response

import (
	"github.com/rebel-l/jirastats/packages/models"
	log "github.com/sirupsen/logrus"
	"time"
)

type StatsComparison struct {
	Project *models.Project `json:"project"`
	Version map[string]string `json:"version"`
	Stats map[string]*StatsComparisonEntry `json:"stats"`
}

func NewStatsComparison(project *models.Project) *StatsComparison {
	sc := new(StatsComparison)
	sc.Project = project
	sc.Version = make(map[string]string, 2)
	sc.Stats = make(map[string]*StatsComparisonEntry)
	return sc
}

func (sc *StatsComparison) SetVersionLeft(value string) {
	sc.setVersion("left", value)
}

func (sc *StatsComparison) SetVersionRight(value string) {
	sc.setVersion("right", value)
}

func (sc *StatsComparison) setVersion(version string, value string) {
	date, err := time.Parse(dateFormatInternal, value)
	if err != nil {
		log.Warnf("Date couldn't be converted: %s", err.Error())
		return
	}

	if date.IsZero() {
		sc.Version[version] = "actual"
	} else {
		sc.Version[version] = date.Format(dateFormatDisplay)
	}
}
