package response

import (
	"github.com/rebel-l/jirastats/packages/models"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

const versionLeft = "left"
const versionRight = "right"

type StatsComparison struct {
	Project *models.Project `json:"project"`
	Version map[string]string `json:"version"`
	VersionOriginal map[string]time.Time
	Stats map[string]*StatsComparisonEntry `json:"stats"`
	wg sync.WaitGroup
}

func NewStatsComparison(project *models.Project) *StatsComparison {
	sc := new(StatsComparison)
	sc.Project = project
	sc.Version = make(map[string]string, 2)
	sc.VersionOriginal = make(map[string]time.Time, 2)
	sc.Stats = make(map[string]*StatsComparisonEntry)
	return sc
}

func (sc *StatsComparison) SetVersionLeft(value time.Time) {
	sc.setVersion(versionLeft, value)
}

func (sc *StatsComparison) SetVersionRight(value time.Time) {
	sc.setVersion(versionRight, value)
}

func (sc *StatsComparison) setVersion(version string, value time.Time) {
	sc.VersionOriginal[version] = value
	if value.IsZero() {
		sc.Version[version] = "actual"
	} else {
		sc.Version[version] = value.Format(DateFormatDisplay)
	}
}

func (sc *StatsComparison) SetStats(counters map[string]*TicketCounter) {
	versionLeft := counters[versionLeft].Tables
	versionRight := counters[versionRight].Tables

	if len(versionLeft) != len(versionRight) {
		log.Errorf("left version and right version have different dimensions, stats not processed")
		return
	}
	for k, vl := range versionLeft {
		vr := versionRight[k]
		entry := NewStatsComparisonEntry()

		sc.wg.Add(1)
		go sc.processPieChart(vl, entry.Left)

		sc.wg.Add(1)
		go sc.processPieChart(vr, entry.Right)

		sc.wg.Add(1)
		go sc.processDiff(vl, vr, entry.Diff)

		sc.wg.Wait()
	}
}

func (sc *StatsComparison) processPieChart(data map[string]*TableData, entries []*PieChartEntry){
	defer sc.wg.Done()
}

func (sc *StatsComparison) processDiff(left map[string]*TableData, right map[string]*TableData, entries []*TableData){
	defer sc.wg.Done()
}
