package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func getHTMLDOC(url string) (doc *html.Node, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s got status code %d", url, resp.StatusCode)
	}

	doc, err = html.Parse(resp.Body)
	if err != nil {
		fmt.Errorf("can not parse body of this %s. err: %s", url)
	}

	return doc, nil
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}

	}()

	getTitle := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}

	forEachNode(getTitle, nil, doc)

	if title == "" {
		return title, fmt.Errorf("doesnt have title")
	}
	return title, nil
}

func forEachNode(pre, post func(n *html.Node), node *html.Node) {
	if pre != nil {
		pre(node)
	}

	for next := node.FirstChild; next != nil; next = next.NextSibling {
		forEachNode(pre, post, next)
	}

	if post != nil {
		post(node)
	}
}

func main() {
	url := os.Args[1]
	doc, err := getHTMLDOC(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(soleTitle(doc))
}
