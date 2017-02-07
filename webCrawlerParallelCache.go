package main

import (
	"fmt"
	"sync"
)


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
  defer func(){ 
    wg.Done() 
  }()
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, wg)
	}
}

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	Crawl("http://golang.org/", 4, fetcher, &wg)
// 	wg.Wait()  
// }

