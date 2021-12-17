package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklists := make(chan []string)
	go func() { worklists <- os.Args[1:]}()

	seen := make(map[string]bool)

	for worklist := range worklists {
		for _, link := range worklist {
			if !seen[link] {
				seen[link] = true
				go func(l string) { worklists <- crawl(link) }(link)
			}
		}
	}

}