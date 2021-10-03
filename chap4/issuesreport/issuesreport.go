package main

import (
	"chap4/github"
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{ range .Items}}-------------------------------------
Number: {{ .Number}}
User:   {{ .User.Username}}
Title: {{ .Title | printf "%.55s" }}
Created At: {{ .CreatedAt | daysAgo }} days ago
{{ end }}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("Issue Report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
