package main

import (
	"flag"
	"fmt"
	"os"

	html "golang.org/x/net/html"
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

	z := html.NewTokenizer(f)
	for {
		line := z.Next()
		if line == html.ErrorToken {
			fmt.Println("Error parsing html")
			return
		}

		switch line {
		case html.StartTagToken, html.EndTagToken: //we only care about tag types because we're looking for links
			tn, _ := z.TagName()
			fmt.Printf("Tag name: %v\n", string(tn))
		case html.TextToken:
			fmt.Println("Text value", z.Text())
		}
		fmt.Println("Line value: ", line)
	}
}
