package links

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func GetLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("error when getting %s. Error: %v", err, url)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s. Status code: %d", url, resp.StatusCode)
	}

	htmlBody, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error when parsing %s. error: %v", url, err)
	}

	var links []string
	pre := func(n *html.Node) {
		// <a href="xyz" style="hoge">example </a>
		// => data == "a"
		// key = href val = xyz
		// key = style val = hoge
		// n.FirstChild.Data = example
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue
					}
					links = append(links, link.String())
				}
			}
		}
	}
	forEachNode(pre, nil, htmlBody)
	return links, nil
}

func forEachNode(pre, post func(n *html.Node), n *html.Node) {
	if pre != nil {
		pre(n)
	}

	for next := n.FirstChild; next != nil; next = next.NextSibling {
		forEachNode(pre, post, next)
	}

	if post != nil {
		post(n)
	}
}
