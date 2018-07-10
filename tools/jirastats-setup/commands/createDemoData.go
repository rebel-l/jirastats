package commands

import (
	"bufio"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const periodRegex = "[0-9]+[wWdDmM]{1}"
const periodUnitDay  = "d"
const periodUnitWeek = "w"
const periodUnitMonth = "m"

type CreateDemoData struct {
	periodValue int
	periodUnit string
	stopDate time.Time
	reader *bufio.Reader
}

func NewCreateDemoData(period string) *CreateDemoData {
	c := new(CreateDemoData)
	c.setPeriod(period)
	c.stopDate = time.Now()
	c.reader = bufio.NewReader(os.Stdin)
	return c
}

func (c * CreateDemoData) Execute() error {
	if c.validate() == false {
		fmt.Println("")
		fmt.Println("")

		fmt.Print("Period couldn't be parsed. Must be in following format: 3d (3 days), 1w (1 week) or 2m (2 months). Please enter the period: ")
		period, _ := c.reader.ReadString('\n')
		fmt.Println("")

		c.setPeriod(period)
		if c.validate() == false {
			return errors.New("create demo data: not able to parse period")
		}
	}

	log.Debugf("Create Demo Data for %d %s", c.periodValue, c.getPeriodUnitHr(c.periodUnit))

	/**
	TODO: create the demo data
	1. Backlog I (Ideal)
	2. Tech Debt I (Ideal)
	3. Short Term Project I (Ideal) ==> 3 to 4 weeks
	4. Long Term Project I (Ideal) ==> 3 month

	4. Backlog II (Running Out Of Tickets)
	5. Backlog III (Speed / Bug / Priority Chaos)
	6. Backlog IV (Cleanup)
	7. Tech Debt II (No Progress / Growing)
	8. Long Term Project II (Endless Story)
	9. Long Term Project III/I (Top of Iceberg)
	10. Long Term Project III/II (Bottom of Iceberg)
	 */
	 actualDate := c.getStartDate()
	 for c.stopDate.After(actualDate) {
	 	weekday := actualDate.Weekday().String()
	 	if weekday != "Saturday" && weekday != "Sunday" {
			log.Debugf("Actual date: %s", actualDate.Format("02.01.2006"))
		}
	 	actualDate = actualDate.AddDate(0, 0, 1)
	 }

	return nil
}

func (c *CreateDemoData) validate() bool {
	if c.periodValue < 1 {
		return false
	}

	if c.periodUnit == "" {
		return false
	}

	return true
}

func (c *CreateDemoData) setPeriod(period string) {
	pRegex, err := regexp.Compile(periodRegex)
	if err != nil {
		return
	}

	period = pRegex.FindString(period)
	if period == "" {
		return
	}

	c.periodUnit = strings.ToLower(period[len(period) - 1:])
	c.periodValue, _ = strconv.Atoi(period[:len(period) - 1])
}

func (c *CreateDemoData) getPeriodUnitHr(puShort string) string {
	puHr := ""
	switch puShort {
	case periodUnitDay:
		puHr = "day(s)"
		break
	case periodUnitWeek:
		puHr = "week(s)"
		break
	case periodUnitMonth:
		puHr = "month(s)"
		break
	}
	return puHr
}

func (c *CreateDemoData) getStartDate() time.Time {
	startDate := c.stopDate
	switch c.periodUnit {
	case periodUnitDay:
		startDate = startDate.AddDate(0, 0, -c.periodValue)
		break
	case periodUnitWeek:
		startDate = startDate.AddDate(0, 0, -c.periodValue * 7)
		break
	case periodUnitMonth:
		startDate = startDate.AddDate(0, -c.periodValue, 0)
		break
	}
	return startDate
}
