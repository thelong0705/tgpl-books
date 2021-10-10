package main

import (
	"chap5/fetch"
	"chap5/findlinks1"
	"chap5/findlinks2"
	"chap5/outline"
	"chap5/wait"
	"fmt"
	"log"
	"os"
)

func main() {
	url := os.Args[1]
	res, err := fetch.Fetch(url)
	if err != nil {
		log.Fatalf("error when fetching url: %v", err)
	}
	firstLink, err := findlinks1.FindFirstLink(res)
	if err != nil {
		log.Fatalf("error when find link for this %s url: %v", url, err)
	}
	fmt.Println("First link of html: ")
	fmt.Println(firstLink)

	allLinks, err := findlinks1.FindAllLinks(res)
	if err != nil {
		log.Fatalf("error when find link for this %s url: %v", url, err)
	}

	fmt.Println("All links of html: ")
	fmt.Println(allLinks)

	fmt.Println("Outline: ")
	outline.Outline(res)

	allLinks, err = findlinks2.FindLink(url)
	if err != nil {
		log.Fatalf("can not find link for this %s url. error: %v", allLinks, err)
	}

	fmt.Println("All links of html: ")
	fmt.Println(allLinks)

	badUrl := "https://bad.golang.io"
	err = wait.WaitForServer(badUrl)
	if err != nil {
		log.Fatalf("Site is down %v", err)
	}

}

//go run main.go https://golang.com
//https://support.eji.org/give/153413/#!/donation/checkout
