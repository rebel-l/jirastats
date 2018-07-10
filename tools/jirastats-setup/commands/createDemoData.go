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
)

const periodRegex = "[0-9]+[wWdDmM]{1}"
const periodUnitDay  = "d"
const periodUnitWeek = "w"
const periodUnitMonth = "m"

type CreateDemoData struct {
	periodValue int
	periodUnit string
	reader *bufio.Reader
}

func NewCreateDemoData(period string) *CreateDemoData {
	c := new(CreateDemoData)
	c.setPeriod(period)
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

	c.periodUnit = period[len(period) - 1:]
	c.periodValue, _ = strconv.Atoi(period[:len(period) - 1])
}

func (c *CreateDemoData) getPeriodUnitHr(puShort string) string {
	puHr := ""
	switch strings.ToLower(puShort) {
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
