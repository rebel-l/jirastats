package endpoints

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/utils"
	"github.com/rebel-l/jirastats/tools/jirastats-server/response"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

const dataStatsOpenTicketsPath = "/data/stats/opentickets/{projectId}"
const defaultKey = "N/A"
const statsTypeNameComponents	= "components"
const statsTypeNameIssueTypes	= "issue_types"
const statsTypeNameLabels		= "labels"
const statsTypeNamePriority		= "priorities"
const statsTypeNameStatus		= "status"
const statsTypeNameTechDebt		= "tech_debt"

type DataStatsOpenTickets struct {
	tm *database.TicketMapper
	pm *database.ProjectMapper
	stats *response.Stats
	series map[string]*response.Serie
}

func NewDataStatsOpenTickets(db *sql.DB, router *mux.Router) {
	ds := new(DataStatsOpenTickets)
	ds.tm = database.NewTicketMapper(db)
	ds.pm = database.NewProjectMapper(db)
	router.HandleFunc(dataStatsOpenTicketsPath, ds.GetStats).Methods(http.MethodGet)
}

func (ds *DataStatsOpenTickets) GetStats(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	projectId, err := strconv.Atoi(vars["projectId"])
	if err != nil {
		msg := "Stats could not be loaded as project id was not provided"
		e := response.NewErrorJson(msg, res)
		e.SendBadRequest()
		return
	}

	log.Debugf("Get all open ticket stats for project: %d", projectId)

	tickets, err := ds.tm.LoadNotExpired(projectId)
	if err != nil {
		msg := fmt.Sprintf("Stats could not be loaded: %s", err.Error())
		e := response.NewErrorJson(msg, res)
		e.SendBadRequest()
		return
	}

	tables := map[string]map[string]*response.TableData{
		statsTypeNameComponents: make(map[string]*response.TableData),
		statsTypeNameIssueTypes: make(map[string]*response.TableData),
		statsTypeNameLabels: make(map[string]*response.TableData),
		statsTypeNamePriority: make(map[string]*response.TableData),
		statsTypeNameStatus: make(map[string]*response.TableData),
		statsTypeNameTechDebt: make(map[string]*response.TableData),
	}
	countTickets := len(tickets)
	for _, t := range tickets {
		// Components
		ds.countTableDataForArrayOfStrings(t.Components, tables[statsTypeNameComponents], "")

		// Issue Types
		ds.countTableDataForSimpleStrings(t.Issuetype, tables[statsTypeNameIssueTypes])

		// Labels
		ds.countTableDataForArrayOfStrings(t.Labels, tables[statsTypeNameLabels], "technicalDebt")

		// Priority
		ds.countTableDataForSimpleStrings(t.Priority, tables[statsTypeNamePriority])

		// Status
		ds.countTableDataForSimpleStrings(t.StatusByJira, tables[statsTypeNameStatus])

		// TechDebt
		if utils.IsValueInMap(t.Labels, "technicalDebt") {
			ds.countTableDataForSimpleStrings("Technical Debt", tables[statsTypeNameTechDebt])
		} else {
			ds.countTableDataForSimpleStrings("Normal", tables[statsTypeNameTechDebt])
		}
		// TODO: parallize processing data
	}

	pc := make(map[string]*response.PieChartTable)
	for k, v := range tables {
		displayName := strings.Title(strings.Replace(k, "_", " ", len(k)))
		finalData := make([]*response.PieChartEntry, len(v))
		pc[k] = ds.createPieChart(displayName, finalData, v, countTickets)
		// TODO: parallize processing data
	}

	success := response.NewSuccessJson(pc, res)
	success.SendOK()
}

func (ds *DataStatsOpenTickets) countTableDataForSimpleStrings(key string, tableData map[string]*response.TableData) {
	if v, ok := tableData[key]; ok {
		v.Value++
	} else {
		tableData[key] = response.NewTableData(key, 1)
	}
}

func (ds *DataStatsOpenTickets) countTableDataForArrayOfStrings(keys []string, tableData map[string]*response.TableData, exclude string) {
	for _, v := range keys {
		if v == "" {
			v = defaultKey
		}

		if v == exclude {
			continue
		}
		ds.countTableDataForSimpleStrings(v, tableData)
	}
}

func (ds *DataStatsOpenTickets) createPieChart(
	name string,
	chart []*response.PieChartEntry,
	data map[string]*response.TableData,
	counter int) *response.PieChartTable {

	pc := response.NewPieChartTable(name, chart)
	i := 0
	maxValue := 0
	maxItem := 0
	for _, d := range data {
		if d.Value > maxValue {
			maxValue = d.Value
			maxItem = i
		}
		chart[i] = response.NewPieChartEntry(d.Name, float64(d.Value) * 100.0 / float64(counter))
		pc.DataTable = append(pc.DataTable, d)
		i++
	}
	chart[maxItem].Sliced = true
	chart[maxItem].Selected = true

	return pc
}
