package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens

	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		for list := range worklist {
			fmt.Println("List: ", list)
			fmt.Println("-------------")
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					n++
					go func(url string) {
						worklist <- crawl(url)
					}(link)
				}
			}
			fmt.Println("N: ", n)
			fmt.Println("++++++++++++")
		}
	}
}
