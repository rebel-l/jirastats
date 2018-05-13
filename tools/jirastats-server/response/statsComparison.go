package response

import (
	"github.com/rebel-l/jirastats/packages/models"
	log "github.com/sirupsen/logrus"
	"time"
)

const versionLeft = "left"
const versionRight = "right"

type StatsComparison struct {
	Project         *models.Project `json:"project"`
	Version         map[string]string `json:"version"`
	versionOriginal map[string]time.Time
	Stats           map[string]*StatsComparisonEntry `json:"stats"`
}

func NewStatsComparison(project *models.Project) *StatsComparison {
	sc := new(StatsComparison)
	sc.Project = project
	sc.Version = make(map[string]string, 2)
	sc.versionOriginal = make(map[string]time.Time, 2)
	sc.Stats = make(map[string]*StatsComparisonEntry)
	return sc
}

func (sc *StatsComparison) GetVersionOriginal() map[string]time.Time {
	return sc.versionOriginal
}

func (sc *StatsComparison) SetVersionLeft(value time.Time) {
	sc.setVersion(versionLeft, value)
}

func (sc *StatsComparison) SetVersionRight(value time.Time) {
	sc.setVersion(versionRight, value)
}

func (sc *StatsComparison) setVersion(version string, value time.Time) {
	sc.versionOriginal[version] = value
	if value.IsZero() {
		sc.Version[version] = "actual"
	} else {
		sc.Version[version] = value.Format(DateFormatDisplay)
	}
}

func (sc *StatsComparison) SetStats(counters map[string]*TicketCounter) {
	versionLeft := counters[versionLeft]
	versionRight := counters[versionRight]

	if len(versionLeft.Tables) != len(versionRight.Tables) {
		log.Errorf("left version and right version have different dimensions, stats not processed")
		return
	}
	for k, vl := range versionLeft.Tables {
		vr := versionRight.Tables[k]

		chLeft := make(chan []*PieChartEntry)
		go sc.processPieChart(vl, chLeft, versionLeft.NumTickets)

		chRight := make(chan []*PieChartEntry)
		go sc.processPieChart(vr, chRight, versionRight.NumTickets)

		chDiff := make(chan []*TableDiffEntry)
		go sc.processDiff(vl, vr, chDiff)

		entry := NewStatsComparisonEntry(k)
		entry.Left = <- chLeft
		entry.Right = <- chRight
		entry.Diff = <- chDiff
		sc.addStats(k, entry)
	}
}

func (sc *StatsComparison) addStats(key string, stats *StatsComparisonEntry) {
	sc.Stats[key] = stats
}

func (sc *StatsComparison) processPieChart(data map[string]*TableData, chEntries chan []*PieChartEntry, numTickets int){
	entries := make([]*PieChartEntry, 0)
	for _, v := range data {
		pce := NewPieChartEntry(v.Name, float64(v.Value) * 100.0 / float64(numTickets))
		entries = append(entries, pce)
	}
	chEntries  <- entries
}

func (sc *StatsComparison) processDiff(left map[string]*TableData, right map[string]*TableData, chEntries chan []*TableDiffEntry){
	prepareEntries := make(map[string]*TableDiffEntry)
	for k, l := range left {
		newEntry := new(TableDiffEntry)
		newEntry.Name = l.Name
		newEntry.Left = l.Value
		newEntry.Diff = -l.Value
		prepareEntries[k] = newEntry
	}

	for k, r := range right {
		newEntry, ok := prepareEntries[k]
		if ok {
			newEntry.Right = r.Value
			newEntry.Diff += r.Value
		} else {
			newEntry := new(TableDiffEntry)
			newEntry.Name = r.Name
			newEntry.Right = r.Value
			newEntry.Diff = r.Value
			prepareEntries[k] = newEntry
		}
	}

	entries := make([]*TableDiffEntry, 0)
	for _, p := range prepareEntries {
		entries = append(entries, p)
	}
	chEntries <- entries
}
