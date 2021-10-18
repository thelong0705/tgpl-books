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
		return fmt.Errorf("error getting %s. error: %v", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %d", url, resp.StatusCode)
	}

	ct := resp.Header.Get("content-type")
	htmlType := "text/html"

	if ct != htmlType && !strings.HasPrefix(ct, htmlType) {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not html type", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visiter := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			if n.FirstChild != nil {
				fmt.Println(n.FirstChild.Data)
			}
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
