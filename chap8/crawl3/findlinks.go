package main

import (
	"fmt"
	links2 "gopl.io/ch5/links"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	links, err := links2.Extract(url)
	if err != nil {
		fmt.Println(err)
	}
	return links
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- os.Args[1:] }()
	for i := 0; i < 20; i++ {
		go func() {
			for l := range unseenLinks {
				foundLinks := crawl(l)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := map[string]bool{}

	for list := range worklist {
		for _, l := range list {
			if !seen[l] {
				seen[l] = true
				unseenLinks <- l
			}
		}
	}
}
