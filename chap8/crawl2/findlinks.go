package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

// tao 1 cai channel tokens
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	links, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}

	return links
}

func main() {
	worklist := make(chan []string, 20)
	var n int
	n++

	go func() { worklist <- os.Args[1:] }()
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(l string) {
					worklist <- crawl(l)
				}(link)
			}
		}
	}
}
