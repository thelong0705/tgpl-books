package main

import (
	"chap5/fetch"
	"chap5/findlinks1"
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
	fisrtLink, err := findlinks1.FindFirstLink(res)
	if err != nil {
		log.Fatalf("error when find link for this %s url: %v", fisrtLink, err)
	}
	fmt.Println(fisrtLink)
}

//go run main.go https://golang.com
//https://support.eji.org/give/153413/#!/donation/checkout