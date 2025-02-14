package main

import (
	"flag"
	"fmt"

	sitemapbuilder "github.com/0xhelloworld-dev/gophercises/sitemapbuilder"
)

func main() {
	targetURL := flag.String("url", "https://calhoun.io/", "target url. include trailing slash (use 'https://google.com/' instead of 'https://google.com'")
	flag.Parse()

	fmt.Println("")

	sitemapbuilder.BuildSiteMap(*targetURL)
}
