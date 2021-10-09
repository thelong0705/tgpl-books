package outline

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
)

func outline(stack []string, node *html.Node) {
	if node.Type == html.ElementNode {
		stack = append(stack, node.Data)
		fmt.Println(stack)
	}

	for next := node.FirstChild; next != nil; next = next.NextSibling {
		outline(stack, next)
	}

	return
}

func Outline(input []byte) error {
	node, err := html.Parse(bytes.NewReader(input))
	if err != nil {
		return fmt.Errorf("can not parse input to html node %q", err)
	}

	outline(nil, node)
	return nil
}
