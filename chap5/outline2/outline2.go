package outline2

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func Outline(url string) error {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("can not get %s. status: %d", url, res.StatusCode)
	}

	doc, err := html.Parse(res.Body)

	if err != nil {
		return fmt.Errorf("can not parse body of this %s. error: %v", url, err)
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for next := n.FirstChild; next != nil ; next = next.NextSibling {
		forEachNode(next, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, ",", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth *2 , ",", n.Data)
	}
}
