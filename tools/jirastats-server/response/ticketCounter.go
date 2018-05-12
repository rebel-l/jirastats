package response

import (
	"github.com/rebel-l/jirastats/packages/models"
	"github.com/rebel-l/jirastats/packages/utils"
)

const defaultKey = "N/A"
const statsTypeNameComponents	= "components"
const statsTypeNameIssueTypes	= "issue_types"
const statsTypeNameLabels		= "labels"
const statsTypeNamePriority		= "priorities"
const statsTypeNameStatus		= "status"
const statsTypeNameTechDebt		= "tech_debt"

type TicketCounter struct {
	Tables map[string]map[string]*TableData
	NumTickets int
}

func NewTicketCounter() *TicketCounter {
	tc := new(TicketCounter)
	tc.Tables = map[string]map[string]*TableData{
		statsTypeNameComponents: make(map[string]*TableData),
		statsTypeNameIssueTypes: make(map[string]*TableData),
		statsTypeNameLabels: make(map[string]*TableData),
		statsTypeNamePriority: make(map[string]*TableData),
		statsTypeNameStatus: make(map[string]*TableData),
		statsTypeNameTechDebt: make(map[string]*TableData),
	}
	return tc
}

func (tc *TicketCounter) Count(tickets []*models.Ticket) {
	tc.NumTickets = len(tickets)
	for _, ticket := range tickets {
		// Components
		tc.countTableDataForArrayOfStrings(ticket.Components, tc.Tables[statsTypeNameComponents], "")

		// Issue Types
		tc.countTableDataForSimpleStrings(ticket.Issuetype, tc.Tables[statsTypeNameIssueTypes])

		// Labels
		tc.countTableDataForArrayOfStrings(ticket.Labels, tc.Tables[statsTypeNameLabels], "technicalDebt")

		// Priority
		tc.countTableDataForSimpleStrings(ticket.Priority, tc.Tables[statsTypeNamePriority])

		// Status
		tc.countTableDataForSimpleStrings(ticket.StatusByJira, tc.Tables[statsTypeNameStatus])

		// TechDebt
		if utils.IsValueInMap(ticket.Labels, "technicalDebt") {
			tc.countTableDataForSimpleStrings("Technical Debt", tc.Tables[statsTypeNameTechDebt])
		} else {
			tc.countTableDataForSimpleStrings("Normal", tc.Tables[statsTypeNameTechDebt])
		}
		// TODO: parallize processing data
	}
}

func (tc *TicketCounter) countTableDataForSimpleStrings(key string, tableData map[string]*TableData) {
	if v, ok := tableData[key]; ok {
		v.Value++
	} else {
		tableData[key] = NewTableData(key, 1)
	}
}

func (tc *TicketCounter) countTableDataForArrayOfStrings(keys []string, tableData map[string]*TableData, exclude string) {
	for _, v := range keys {
		if v == "" {
			v = defaultKey
		}

		if v == exclude {
			continue
		}
		tc.countTableDataForSimpleStrings(v, tableData)
	}
}
