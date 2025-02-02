package htmllinkparser

import (
	"fmt"
	"io"

	html "golang.org/x/net/html"
)

// Values in an <a> tag.
type Link struct {
	Href string
	Text string
}

// Parse takes HTML document and returns a slice of our Link type.
func ParseLinks(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	fmt.Println(padding, n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
