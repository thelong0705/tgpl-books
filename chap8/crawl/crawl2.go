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
		log.Println(err)
	}
	return list
}
var sendlink, receivelink int
func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				receivelink++
				fmt.Println("receive link: ", receivelink)
				wl := crawl(link)
				//go func() { worklist<- wl }()
				worklist<- wl
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				sendlink++
				fmt.Println("sendlink: ",sendlink)

				unseenLinks <- link
			}
		}
	}
}
