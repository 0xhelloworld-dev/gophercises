package queue

import (
	"fmt"
	"slices"

	"github.com/0xhelloworld-dev/gophercises/sitemapbuilder/xmlutils"
)

type Queue struct {
	UrlQueue    []string
	ScannedURLs []string
}

func (q *Queue) Enqueue(url string) {
	q.UrlQueue = append(q.UrlQueue, url)
}

func (q *Queue) Dequeue() (string, error) {
	//Gets first item from the queue and removes it from the queue
	if len(q.UrlQueue) == 0 {
		return "", fmt.Errorf("[WARNING] Queue is empty")
	}
	nextUrl := q.UrlQueue[0]
	q.UrlQueue = q.UrlQueue[1:len(q.UrlQueue)] //remove first item from queue
	q.ScannedURLs = append(q.ScannedURLs, nextUrl)
	return nextUrl, nil
}

func (q *Queue) InQueue(url string) bool {
	if slices.Contains(q.UrlQueue, url) {
		return true
	} else {
		return false
	}
}

// After we populate the initial queue of URLs
// ProcessQueue is responsible for continuously processing the next item in the queue and passing it to our processFunc()
func (q *Queue) ProcessQueue(processFunc func(string, *Queue, *xmlutils.URLSet)) {
	for {
		if len(q.UrlQueue) == 0 {
			fmt.Printf("Finished queue")
			break
		}
		url, _ := q.Dequeue()
		processFunc(url, q, xmlutils.Sitemap)
	}
}
