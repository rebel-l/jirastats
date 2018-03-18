package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/rebel-l/jirastats/packages/jira"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)

type ConfigureApplication struct {
	jiraConfig *jira.Config
}

func NewConfigureApplication(db *sql.DB) *ConfigureApplication {
	ca := new(ConfigureApplication)
	ca.jiraConfig = jira.NewConfig(db)
	return ca
}

func (ca *ConfigureApplication) Execute() (err error) {
	r := bufio.NewReader(os.Stdin)

	fmt.Println("")
	fmt.Println("")

	// host
	if ca.jiraConfig.GetHost() != "" {
		fmt.Printf("Enter host to jira, actual is: %s. Press enter to keep it.", ca.jiraConfig.GetHost())
	} else {
		fmt.Println("Enter host to jira, e.g. https.//example.jira.com")
	}
	host, _ := r.ReadString('\n')
	host = strings.TrimSpace(host)
	if host != "" {
		ca.jiraConfig.SetHost(host)
	}

	// username
	if ca.jiraConfig.GetUsername() != "" {
		fmt.Printf("Enter username for jira, actual is: %s. Press enter to keep it.", ca.jiraConfig.GetUsername())
	} else {
		fmt.Println("Enter username for jira")
	}
	username, _ := r.ReadString('\n')
	username = strings.TrimSpace(username)
	if username != "" {
		ca.jiraConfig.SetUsername(username)
	}

	// password
	if ca.jiraConfig.GetPassword() != "" {
		fmt.Println("Enter password for jira. Press enter to keep the current one.")
	} else {
		fmt.Println("Enter password for jira")
	}
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := strings.TrimSpace(string(bytePassword))
	if password != "" {
		ca.jiraConfig.SetPassword(password)
	}

	fmt.Println("")
	fmt.Println("")

	err = ca.jiraConfig.Save()
	return
}
