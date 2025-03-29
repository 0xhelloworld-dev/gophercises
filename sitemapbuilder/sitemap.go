package sitemapbuilder

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/0xhelloworld-dev/gophercises/sitemapbuilder/queue"
	"github.com/0xhelloworld-dev/gophercises/sitemapbuilder/smUtils"
	"github.com/0xhelloworld-dev/gophercises/sitemapbuilder/xmlutils"

	linkparser "github.com/0xhelloworld-dev/gophercises/htmllinkparser"
)

var globalBaseURL string

func BuildSiteMap(baseURL string) {
	urlsQueue := &queue.Queue{}
	globalBaseURL = baseURL
	xmlutils.Sitemap = &xmlutils.URLSet{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	getLinksFromPage(baseURL, urlsQueue, xmlutils.Sitemap)

	//need to insert next item in urlsQueue.UrlQueue[0] as variable into ProcessQueue(urlsQueue.UrlQueue[0])
	urlsQueue.ProcessQueue(getLinksFromPage)
	saveSitemap(*xmlutils.Sitemap, "sitemap.xml")
}

func getLinksFromPage(targetURL string, queue *queue.Queue, sitemap *xmlutils.URLSet) {
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

	processLinks(links, queue, xmlutils.Sitemap)

	fmt.Printf("[+] CurrentQueue: %v\n----------\n[+] ScannedUrls: %s\n", queue.UrlQueue, queue.ScannedURLs)
	fmt.Printf("----------\n[+] Queue length: %v\n[+] ScannedUrls length: %d\n-------------------------\n", len(queue.UrlQueue), len(queue.ScannedURLs))
	time.Sleep(1 * time.Second)
}

func processLinks(links []linkparser.Link, queue *queue.Queue, sitemap *xmlutils.URLSet) {
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
		*sitemap = xmlutils.URLSet{
			XMLName: sitemap.XMLName,
			Xmlns:   sitemap.Xmlns,
			URLs:    append(sitemap.URLs, xmlutils.URL{Loc: normalizedURL}),
		}
		//fmt.Printf("\t[+] Added Url to Queue:[%s]\n", normalizedURL)
	}
}

// saveSitemap
func saveSitemap(sitemap xmlutils.URLSet, filename string) error {
	output, err := xml.MarshalIndent(sitemap, "", "\t")
	if err != nil {
		return err
	}

	xmlHeader := []byte(xml.Header)
	output = append(xmlHeader, output...)
	return os.WriteFile(filename, output, 0644)
}
