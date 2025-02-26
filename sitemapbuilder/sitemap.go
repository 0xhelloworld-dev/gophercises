package sitemapbuilder

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/0xhelloworld-dev/gophercises/sitemapbuilder/queue"

	linkparser "github.com/0xhelloworld-dev/gophercises/htmllinkparser"
)

func BuildSiteMap(targetURL string) {
	urlsQueue := &queue.Queue{}
	getLinksFromPage(targetURL, urlsQueue)

	//need to insert next item in urlsQueue.UrlQueue[0] as variable into ProcessQueue(urlsQueue.UrlQueue[0])
	urlsQueue.ProcessQueue(func(targetURL string, q *queue.Queue) { getLinksFromPage(targetURL, urlsQueue) })
}

func getLinksFromPage(targetURL string, queue *queue.Queue) {
	//instead of returning the links, it should add these links to our queue
	fmt.Printf("~~~~~~~~~Executing getLinksFromPage for %v~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\n", targetURL)
	resp, err := http.Get(targetURL)
	if err != nil {
		fmt.Println("Error:", err)
	}
	links, err := linkparser.ParseLinks(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("[+] CurrentQueue: %v\n[+] ScannedUrls: %s\n", queue.UrlQueue, queue.ScannedURLs)
	fmt.Printf("~~~~~~~~~~~~~~~~\n[+] Queue length: %v\n[+] ScannedUrls length: %d\n~~~~~~~~~~~~~~~~\n", len(queue.UrlQueue), len(queue.ScannedURLs))
	for _, link := range links {
		isInScope := inScope(link.Href, targetURL) //important to put link.Href here, not normalized link
		if !isInScope {
			//fmt.Printf("\t[+] Not in scope: %s \n", link.Href)
			continue
		}
		//fmt.Printf("\t[+] Link Href: %s\n", link.Href)
		normalizedURL := normalizeHref(link.Href, targetURL)
		isScanned := isLinkScanned(normalizedURL, queue.ScannedURLs)
		isInQueue := queue.InQueue(normalizedURL)
		if isScanned || isInQueue {
			//we don't want to queue up anything that is already in queue, scanned, or is not in scope
			//fmt.Printf("\t[+] Skipping url: [%s]\n", normalizedURL)
			continue
		}

		queue.Enqueue(normalizedURL)
		time.Sleep(100 * time.Millisecond)
		//fmt.Printf("\t[+] Added Url to Queue:[%s]\n", normalizedURL)
	}
	time.Sleep(5 * time.Second)
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
func normalizeHref(href string, targetURL string) string {
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

// Accepts a raw Href from link.Href
// Ideally you should check if it is in scope before processing
func inScope(href string, targetURL string) bool {
	if strings.HasPrefix(href, targetURL) { //does href have prefix of https://target.com?
		return true
	} else if strings.HasPrefix(href, "/") { //is it a relative href "/about"
		return true
	} else {
		return false
	}

	//need to account for cases where relative link is "#" or ","
}

//there is a bug in how i'm Queuing things up
//it adds the full url
