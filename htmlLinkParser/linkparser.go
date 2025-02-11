package htmllinkparser

import (
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
	docNode, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	searchLinks(docNode, "")
	return linkList, nil
}

func searchLinks(n *html.Node, padding string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		var linkEntry Link
		var url string
		for _, attr := range n.Attr { //iterate through all the attributes in the <a> tag
			if attr.Key == "href" {
				url = attr.Val
				linkEntry.Href = url
			}
		}
		linkEntry.Href = url
		//fmt.Printf("Attributes: %v\n", url)
		if n.FirstChild != nil {
			text := n.FirstChild.Data
			linkEntry.Text = text
		}
		linkList = append(linkList, linkEntry)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		searchLinks(c, padding+"  ")
	}
}
