package sitemapbuilder

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	linkparser "github.com/0xhelloworld-dev/gophercises/htmllinkparser"
)

var scannedLinks []string

func BuildSiteMap(targetURL string) {
	links := getAndParsePage(targetURL)
	callNumber := 1
	scanLinks(links, targetURL, targetURL, callNumber)
}

func getAndParsePage(targetURL string) []linkparser.Link {
	resp, err := http.Get(targetURL)
	if err != nil {
		fmt.Println("Error:", err)
	}
	links, err := linkparser.ParseLinks(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return links
}

// DFS search
// Links are gathered from every page
// for each link found from the page:
//
//	/index.html
//		-> /hello.html (scan)
//			-> /about.html (scan)
//				-> /deep.html (scan)
//				-> /profile.html (scan)
//			-> /profile.html (skip, we've seen it in about.html)
//		-> /profile.html (skip, we've seen it in about.html)
//			-> deep.html (skip, we've seen it in about.html)
//	check if we scanned it before
//		if we have skip it
//	if we havent scanned it, scan it
func scanLinks(links []linkparser.Link, domainTarget string, currentURL string, callNumber int) {
	fmt.Printf("~~~~~Running scanLinks() for %s~~~~~~~~~~~~ callNumber: %d\n\n", currentURL, callNumber)
	//var currentQueue []string
	fmt.Printf("Iterating through every link gathered from %s\n\n", currentURL)
	for _, link := range links {
		isInScope := inScope(link.Href, domainTarget)
		if isInScope {
			//check if it has been scanned
			targetLink := formatHref(link.Href, domainTarget)
			isScanned := isLinkScanned(targetLink, scannedLinks)
			//fmt.Printf("Target Link: %s\n\n", targetLink)
			if isScanned {
				//fmt.Printf("%v has been scanned, skipping\n", targetLink)
				continue
			}
			targetPageLinks := getAndParsePage(targetLink)
			fmt.Printf("Got links for %s\n\n", targetLink)

			//currentQueue = addToQueue(targetPageLinks, currentQueue)

			scannedLinks = append(scannedLinks, targetLink)
			fmt.Printf("Scanned Links: %v\n\n", scannedLinks)
			time.Sleep(1 * time.Second)

			scanLinks(targetPageLinks, domainTarget, targetLink, callNumber+1)
		} else {
			//fmt.Printf("Skipping %s - not in scope\n\n", link.Href)
		}
	}
	fmt.Printf("Finished iterating through links for %s\n", currentURL)
}

func addToQueue(targetPageLinks []linkparser.Link, currentQueue []string) []string {
	for _, link := range targetPageLinks {
		for _, inQueue := range currentQueue {
			if link.Href == inQueue {
				fmt.Printf("Link from new page %s matches link inQueue [%s]\n\n", link.Href, inQueue)
				continue
			}
			currentQueue = append(currentQueue, link.Href)
			fmt.Printf("Current Queue: %v\n\n", currentQueue)
		}
	}
	return currentQueue
}

func isLinkScanned(targetLink string, scannedLinks []string) bool {
	for _, link := range scannedLinks {
		if link == targetLink {
			return true
		}
	}
	return false
}

// Transforms all links to a universal format: https://{domain}/{path}
func formatHref(href string, targetURL string) string {
	if strings.HasPrefix(href, targetURL) {
		//Addresses "http://test.com/about" cases
		return href
	} else if strings.HasPrefix(href, "/") {
		//Addresses "/about" cases. Returns "https://test.com/about"
		formattedURL := targetURL[:len(targetURL)-1] + href
		return formattedURL
	} else {
		//Addresses "#" cases. Returns https://test.com/#
		formattedURL := targetURL + href
		return formattedURL
	}
}

func inScope(href string, targetURL string) bool {
	if strings.HasPrefix(href, targetURL) {
		return true
	} else if strings.HasPrefix(href, "/") {
		return true
	} else {
		return false
	}

	//need to account for cases where relative link is "#" or ","
}
