package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/rebel-l/jirastats/packages/database"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

// ResetStats command deletes all stats data
type ResetStats struct {
	db *sql.DB
}

// NewResetsStats returns ResetStats struct
func NewResetStats(db *sql.DB) *ResetStats {
	rs := new(ResetStats)
	rs.db = db
	return rs
}

// Execute ResetStats command
func (rs *ResetStats) Execute() (err error) {
	if rs.getConfirmation() != "y" {
		return
	}

	err = rs.resetTicketTable()
	if err != nil {
		return
	}

	err = rs.resetStatsTable()

	log.Info("Stats data deleted ...")
	return
}

// resetTicketTable deletes all data in table ticket
func (rs *ResetStats) resetTicketTable() error {
	tt := database.NewTicketTable(rs.db)
	return tt.Truncate()
}

// resetStatsTable deletes all data in table stats
func (rs *ResetStats) resetStatsTable() error {
	s := database.NewSatsTable(rs.db)
	return s.Truncate()
}

// getConfirmation asks user for confirmation of the command and returns the answer
func (rs *ResetStats) getConfirmation() string {
	log.Info("Stats data will be deleted ...")
	r := bufio.NewReader(os.Stdin)
	fmt.Println("")
	fmt.Print("Every stats data will be deleted ... are you sure?  (y/N)")
	confirm, _ := r.ReadString('\n')
	fmt.Println("")

	confirm = strings.TrimSpace(confirm)
	confirm = strings.ToLower(confirm)
	return confirm
}
