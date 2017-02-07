package main

import (
	"sync"
)

// ParallelCrawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// use a sync.WaitGroup to wait for all parallel crawls to finish
func ParallelCrawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}
	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		// fmt.Println(err)
		return
	}
	// fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go ParallelCrawl(u, depth-1, fetcher, wg)
	}
}

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	ParallelCrawl("http://golang.org/", 4, fetcher, &wg)
// 	wg.Wait()
// }
