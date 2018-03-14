package main

import (
	//"bufio"
	"fmt"
	"github.com/andygrunwald/go-jira"
	//jiraSearch "github.com/rebel-l/jirastats/server/jira"
	log "github.com/sirupsen/logrus"
	"io"
	//"os"
	//"golang.org/x/crypto/ssh/terminal"
	//"syscall"
	"strings"
)

func main() {
	log.SetLevel(log.DebugLevel)
	//log.Debug("Run Jira Stats ...")
	//
	//// get credentials
	//r := bufio.NewReader(os.Stdin)
	//fmt.Println("")
	//fmt.Print("Jira Username: ")
	//username, _ := r.ReadString('\n')
	//fmt.Print("Jira Password: ")
	//bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	//password := string(bytePassword)
	//fmt.Println("")
	//fmt.Println("")

	// init Jira client
	//jiraClient := initJiraClient(username, password)
	//jiraExampleTicket(jiraClient)
	//jiraExampleSearch(jiraClient)
	//jiraSearch := jiraSearch.NewSearch(jiraClient)
	//jiraSearch.Do("project = CORE")
	log.Debug("Stopping Jira Stats ... Goodbye!")
}

func printBody(body io.ReadCloser) {
	var p []byte
	body.Read(p)
	log.Debugf("Body: %s", string(p))
}

func jiraExampleProjects() {
	jiraClient, _ := jira.NewClient(nil, "https://jira.atlassian.com/")
	req, _ := jiraClient.NewRequest("GET", "/rest/api/2/project", nil)

	projects := new([]jira.Project)
	_, err := jiraClient.Do(req, projects)
	if err != nil {
		panic(err)
	}


	for _, project := range *projects {
		fmt.Printf("%s: %s\n", project.Key, project.Name)
	}
}

func jiraExampleTicket(client *jira.Client) {
	issue, response, err := client.Issue.Get("CORE-1674", nil)

	if err != nil {
		log.Errorf("Error: %s", err.Error())
		printBody(response.Body)
		return
	}

	printBody(response.Body)

	log.Debugf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	log.Debugf("Type: %s\n", issue.Fields.Type.Name)
	log.Debugf("Priority: %s\n", issue.Fields.Priority.Name)
}

func initJiraClient(username string, password string) *jira.Client {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	jiraClient, _ := jira.NewClient(nil, "https://jira.home24.de")
	jiraClient.Authentication.SetBasicAuth(username, password)
	return jiraClient
}
