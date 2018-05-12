package response

import (
	"github.com/rebel-l/jirastats/packages/models"
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

func (sc *StatsComparison) SetVersionLeft(value time.Time) {
	sc.setVersion("left", value)
}

func (sc *StatsComparison) SetVersionRight(value time.Time) {
	sc.setVersion("right", value)
}

func (sc *StatsComparison) setVersion(version string, value time.Time) {
	if value.IsZero() {
		sc.Version[version] = "actual"
	} else {
		sc.Version[version] = value.Format(DateFormatDisplay)
	}
}
