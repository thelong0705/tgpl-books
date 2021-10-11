package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"database":              {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t %s \n", i+1, course)
	}
	fmt.Println(topoSort(prereqs))
}

func topoSort(m map[string][]string) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	seen := make(map[string]bool)
	var visitAll func(course []string)
	var orders []string

	visitAll = func(courses []string) {
		for _, course := range courses {
			if !seen[course] {
				seen[course] = true
				visitAll(m[course])
				orders = append(orders, course)
			}
		}
	}
	visitAll(keys)
	return orders
}
