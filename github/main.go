package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

// Must method parses the template to ensure no errors;
// New method creates and returns a new template
var report = template.Must(template.New("issuelist").Parse(IssueList))

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
