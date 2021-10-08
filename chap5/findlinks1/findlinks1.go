package findlinks1

import (
	"bytes"
	"errors"
	"golang.org/x/net/html"
)

func visit(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attribute := range n.Attr {
			if attribute.Key == "href" {
				return attribute.Val
			}
		}
	}

	for next := n.FirstChild; next != nil; next = next.NextSibling {
		res := visit(next)
		if res != "" {
			return res
		}
	}

	return ""
}

func visitAndRecord(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attribute := range n.Attr {
			if attribute.Key == "href" {
				links = append(links, attribute.Val)
				return links
			}
		}
	}

	for next := n.FirstChild; next != nil; next = next.NextSibling {
		links = visitAndRecord(links, next)
	}

	return links
}

func FindFirstLink(input []byte) (string, error) {
	node, err := html.Parse(bytes.NewReader(input))
	if err != nil {
		return "", err
	}
	firstLink := visit(node)
	if firstLink == "" {
		return "", errors.New("no link in this html")
	}
	return firstLink, nil
}

func FindAllLinks(input []byte) ([]string, error) {
	node, err := html.Parse(bytes.NewReader(input))

	if err != nil {
		return nil, err
	}
	var links []string
	links = visitAndRecord(links, node)
	return links, nil
}
