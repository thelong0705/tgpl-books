package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %d", url, resp.StatusCode)
	}

	contentType := resp.Header.Get("content-type")
	htmlType := "text/html"

	if contentType != "text/html" && !strings.HasPrefix(contentType, htmlType) {
		return fmt.Errorf("%s has %s type not html type", url, contentType)
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return fmt.Errorf("error when parsing %s body. %v", url, err)
	}
	visiter := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visiter, nil)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for next := n.FirstChild; next != nil; next = next.NextSibling {
		forEachNode(next, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func main() {
	url := os.Args[1]
	err := title(url)
	if err != nil {
		log.Fatal(err)
	}
}
