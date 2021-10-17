package main

import (
	"chap5/links"
	"fmt"
	"log"
	"os"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil

		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.GetLinks(url)
	if err != nil {
		log.Print(err)
	}

	return list
}


func main() {
	urls := os.Args[1:]
	breadthFirst(crawl, urls)
}
