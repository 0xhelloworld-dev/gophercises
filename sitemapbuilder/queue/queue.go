package queue

import (
	"fmt"
	"slices"
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

func (q *Queue) ProcessQueue(processFunc func(string, *Queue)) {
	for {
		if len(q.UrlQueue) == 0 {
			break
		}
		url, _ := q.Dequeue()
		processFunc(url, q)
	}
}
