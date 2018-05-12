package models

import (
	"github.com/rebel-l/jirastats/packages/models"
	"github.com/rebel-l/jirastats/packages/utils"
	"github.com/rebel-l/jirastats/tools/jirastats-server/response"
)

const defaultKey = "N/A"
const statsTypeNameComponents	= "components"
const statsTypeNameIssueTypes	= "issue_types"
const statsTypeNameLabels		= "labels"
const statsTypeNamePriority		= "priorities"
const statsTypeNameStatus		= "status"
const statsTypeNameTechDebt		= "tech_debt"

type TicketCounter struct {
	Tables map[string]map[string]*response.TableData
}

func NewTicketCounter() *TicketCounter {
	tc := new(TicketCounter)
	tc.Tables = map[string]map[string]*response.TableData{
		statsTypeNameComponents: make(map[string]*response.TableData),
		statsTypeNameIssueTypes: make(map[string]*response.TableData),
		statsTypeNameLabels: make(map[string]*response.TableData),
		statsTypeNamePriority: make(map[string]*response.TableData),
		statsTypeNameStatus: make(map[string]*response.TableData),
		statsTypeNameTechDebt: make(map[string]*response.TableData),
	}
	return tc
}

func (tc *TicketCounter) Count(tickets []*models.Ticket) {
	for _, t := range tickets {
		// Components
		tc.countTableDataForArrayOfStrings(t.Components, tc.Tables[statsTypeNameComponents], "")

		// Issue Types
		tc.countTableDataForSimpleStrings(t.Issuetype, tc.Tables[statsTypeNameIssueTypes])

		// Labels
		tc.countTableDataForArrayOfStrings(t.Labels, tc.Tables[statsTypeNameLabels], "technicalDebt")

		// Priority
		tc.countTableDataForSimpleStrings(t.Priority, tc.Tables[statsTypeNamePriority])

		// Status
		tc.countTableDataForSimpleStrings(t.StatusByJira, tc.Tables[statsTypeNameStatus])

		// TechDebt
		if utils.IsValueInMap(t.Labels, "technicalDebt") {
			tc.countTableDataForSimpleStrings("Technical Debt", tc.Tables[statsTypeNameTechDebt])
		} else {
			tc.countTableDataForSimpleStrings("Normal", tc.Tables[statsTypeNameTechDebt])
		}
		// TODO: parallize processing data
	}
}

func (tc *TicketCounter) countTableDataForSimpleStrings(key string, tableData map[string]*response.TableData) {
	if v, ok := tableData[key]; ok {
		v.Value++
	} else {
		tableData[key] = response.NewTableData(key, 1)
	}
}

func (tc *TicketCounter) countTableDataForArrayOfStrings(keys []string, tableData map[string]*response.TableData, exclude string) {
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
