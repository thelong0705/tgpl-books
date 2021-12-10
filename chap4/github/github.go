package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesUrl = "https://api.github.com/search/issues"
const SearchField = "?q="

type SearchIssuesResult struct {
	TotalCount int `json:"total_count"`
	Items      []item
}

type item struct {
	Number    int
	Title     string
	CreatedAt time.Time `json:"created_at"`
	User      user
}

type user struct {
	Username string `json:"login"`
}

func SearchIssues(terms []string) (*SearchIssuesResult, error) {
	queryParams := url.QueryEscape(strings.Join(terms, " "))
	searchQuery := IssuesUrl + SearchField + queryParams
	resp, err := http.Get(searchQuery)
	defer func() { resp.Body.Close() }()
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search failed: %s", resp.Status)
	}

	var result SearchIssuesResult
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(respByte, &result)
	return &result, nil

}

