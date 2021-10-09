package findlinks2

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for next := node.FirstChild; next != nil; next = next.NextSibling {
		links = visit(links, next)
	}

	return links
}

func FindLink(url string) (links []string, err error) {
	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("can not get %s. error: %v", url, err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can not get %s. response status: %d", url, res.StatusCode)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error when parsing html from %s. error: %v", url, err)
	}

	return visit(nil, doc), nil

}
