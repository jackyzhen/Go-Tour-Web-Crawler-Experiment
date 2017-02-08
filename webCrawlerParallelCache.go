// implements parallel and cache via sync wait group and a concurrent map

package main

import (
	"sync"
)

// ParallelCacheCrawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func ParallelCacheCrawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup, safeCache *SafeCache) {

  defer wg.Done() 
	if depth <= 0 {
		return
	}
	_, urls, err := fetcher.Fetch(url)
	safeCache.Write(url)
	if err != nil {
		// fmt.Println(err)
		return
	}
	// fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if _, ok := safeCache.Value(u); !ok {
			wg.Add(1)
			go ParallelCacheCrawl(u, depth-1, fetcher, wg, safeCache)
    }
	}
}

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go ParallelCacheCrawl("http://golang.org/", 4, fetcher, &wg, cache)
// 	wg.Wait()  
// }

