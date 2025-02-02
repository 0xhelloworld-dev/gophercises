package main

import (
	"flag"
	"fmt"
	"os"

	linkparser "github.com/0xhelloworld-dev/gophercises/htmllinkparser"
)

func main() {
	htmlFilename := flag.String("filename", "/Users/jonathanchua/Desktop/goProjects/gophercises/htmlLinkParser/sampleHtml/ex1.html", "name of the html file you want to parse (ex: hello.html)")
	flag.Parse()

	f, err := os.Open(*htmlFilename)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	//verify we are opening a file
	//log.Printf("Printing value of 'f': %v", f)

	links, err := linkparser.ParseLinks(f)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(links)

	/*
		rootNode, err := html.Parse(f)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		for node := range rootNode.Descendants() {
			if node.Type == html.ElementNode {
				fmt.Printf("Node data: %v\n", node.Data)
			}
		}
	*/
}
