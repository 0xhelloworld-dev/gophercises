package sitemapbuilder

import (
	"fmt"
	"net/http"
	"time"

	"github.com/0xhelloworld-dev/gophercises/sitemapbuilder/queue"
	"github.com/0xhelloworld-dev/gophercises/sitemapbuilder/smUtils"

	linkparser "github.com/0xhelloworld-dev/gophercises/htmllinkparser"
)

var globalBaseURL string

func BuildSiteMap(baseURL string) {
	urlsQueue := &queue.Queue{}
	globalBaseURL = baseURL
	getLinksFromPage(baseURL, urlsQueue)

	//need to insert next item in urlsQueue.UrlQueue[0] as variable into ProcessQueue(urlsQueue.UrlQueue[0])
	urlsQueue.ProcessQueue(getLinksFromPage)
}

func getLinksFromPage(targetURL string, queue *queue.Queue) {
	//instead of returning the links, it should add these links to our queue
	fmt.Printf("~~~~~~~~~Executing getLinksFromPage for %v~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\n", targetURL)
	resp, err := http.Get(targetURL)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()
	links, err := linkparser.ParseLinks(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}

	processLinks(links, queue)

	fmt.Printf("[+] CurrentQueue: %v\n----------\n[+] ScannedUrls: %s\n", queue.UrlQueue, queue.ScannedURLs)
	fmt.Printf("----------\n[+] Queue length: %v\n[+] ScannedUrls length: %d\n-------------------------\n", len(queue.UrlQueue), len(queue.ScannedURLs))
	time.Sleep(1 * time.Second)
}

func processLinks(links []linkparser.Link, queue *queue.Queue) {
	for _, link := range links {
		isInScope := smUtils.InScope(link.Href, globalBaseURL) //important to put link.Href here, not normalized link
		if !isInScope {
			//fmt.Printf("\t[+] Not in scope: %s \n", link.Href)
			continue
		}
		//fmt.Printf("\t[+] Link Href: %s\n", link.Href)
		normalizedURL := smUtils.NormalizeHref(link.Href, globalBaseURL)
		isScanned := smUtils.IsLinkScanned(normalizedURL, queue.ScannedURLs)
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
}

//there is a bug in how i'm Queuing things up
//it adds the full url
//now that links are being queued and removed properly, i need to add something to generate the sitemap once finished
//i also need to examine my code and make sure it makes sense
