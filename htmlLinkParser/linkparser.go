package htmllinkparser

import (
	"fmt"
	"io"

	html "golang.org/x/net/html"
)

var linkList []Link

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
	return linkList, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
		if n.Data == "a" {
			var linkEntry Link
			url := n.Attr[0].Val
			linkEntry.Href = url
			//fmt.Printf("Attributes: %v\n", url)
			if n.FirstChild != nil {
				text := n.FirstChild.Data
				linkEntry.Text = text
			}
			linkList = append(linkList, linkEntry)
		}
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
