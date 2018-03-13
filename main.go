package main

import (
	"bufio"
	"fmt"
	"github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
	"strings"
)

var jiraClient *jira.Client

func main() {
	log.SetLevel(log.DebugLevel)
	log.Debug("Run Jira Stats ...")

	// get credentials
	r := bufio.NewReader(os.Stdin)
	fmt.Println("")
	fmt.Print("Jira Username: ")
	username, _ := r.ReadString('\n')
	fmt.Print("Jira Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)
	fmt.Println("")
	fmt.Println("")

	// init Jira client
	initJiraClient(username, password)
	//jiraExampleTicket()
	jiraExampleSearch()
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

func jiraExampleTicket() {
	issue, response, err := jiraClient.Issue.Get("CORE-1674", nil)

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

func initJiraClient(username string, password string) {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	jiraClient, _ = jira.NewClient(nil, "https://jira.home24.de")
	jiraClient.Authentication.SetBasicAuth(username, password)
}

func jiraExampleSearch() {
	search := JiraSearchRequest{
		Jql: "project = CORE",
		StartAt: 0,
		MaxResults: 1,
	}

	result := new(JiraSearchResponse)

	req, _ := jiraClient.NewRequest("POST", "/rest/api/2/search", search)
	_, err := jiraClient.Do(req, result)
	if err != nil {
		log.Error(err)
		return
	}

	log.Debugf("Total: %d", result.Total)
	for _,v := range result.Issues {
		log.Debugf("Issue: %s", v.Key)
	}
	//printBody(response.Body)
}

type JiraSearchRequest struct {
	Jql string `json:"jql"`
	StartAt int `json:"startAt"`
	MaxResults int `json:"maxResults"`
}

type JiraSearchResponse struct {
	Total int `json:"total"`
	Issues []jira.Issue `json:"issues"`
}