package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/rebel-l/jirastats/packages/database"
	"github.com/rebel-l/jirastats/packages/models"
	"os"
	"strconv"
	"strings"
)

type ConfigureProjects struct {
	reader *bufio.Reader
	mapper *database.ProjectMapper
}

func NewConfigureProjects(db *sql.DB) *ConfigureProjects {
	cp := new(ConfigureProjects)
	cp.reader = bufio.NewReader(os.Stdin)
	cp.mapper = database.NewProjectMapper(db)
	return cp
}

func (cp *ConfigureProjects) Execute() (err error) {
	loop := true

	fmt.Println("")
	fmt.Println("")

	for loop {
		cp.createProject()
		loop = cp.next()
	}


	fmt.Println("")
	fmt.Println("")
	return
}

func (cp *ConfigureProjects) next() bool {
	fmt.Println("")
	fmt.Print("Create another project?  (y/N)")
	confirm, _ := cp.reader.ReadString('\n')
	fmt.Println("")

	confirm = strings.TrimSpace(confirm)
	confirm = strings.ToLower(confirm)
	return confirm == "y"
}

func (cp *ConfigureProjects) createProject() (err error) {
	p := models.NewProject()

	fmt.Print("Project name: ")
	p.Name, _ = cp.reader.ReadString('\n')
	p.Name = strings.TrimSpace(p.Name)

	fmt.Print("Add Jira project keys (can be one or more seperated by comma): ")
	p.Keys, _ = cp.reader.ReadString('\n')
	p.Keys = strings.TrimSpace(p.Keys)

	fmt.Printf("JQL Query, actual is '%s'. Press enter to keep it: ", p.Jql)
	jql, _ := cp.reader.ReadString('\n')
	jql = strings.TrimSpace(jql)
	if jql != "" {
		p.Jql = jql
	}

	fmt.Printf(
		"Known speed (closed tickets per day used only for projects lasting less than 5 days)" +
			", actual value is %f. Press enter to keep: ",
		p.KnownSpeed)
	ks, _ := cp.reader.ReadString('\n')
	ks = strings.TrimSpace(ks)
	if ks != "" {
		ksf, err := strconv.ParseFloat(ks, 32)
		if err == nil {
			p.KnownSpeed = float32(ksf)
		}
	}

	fmt.Printf("Jira status mapped to 'open', actual is '%s'. Press enter to keep it: ", p.MapOpenStatus)
	mos, _ := cp.reader.ReadString('\n')
	mos = strings.TrimSpace(mos)
	if mos != "" {
		p.MapOpenStatus = mos
	}

	fmt.Printf("Jira status mapped to 'closed', actual is '%s'. Press enter to keep it: ", p.MapClosedStatus)
	mcs, _ := cp.reader.ReadString('\n')
	mcs = strings.TrimSpace(mcs)
	if mcs != "" {
		p.MapClosedStatus = mcs
	}

	err = cp.mapper.Save(p)
	fmt.Printf("Id: %d\n", p.Id)

	return
}
