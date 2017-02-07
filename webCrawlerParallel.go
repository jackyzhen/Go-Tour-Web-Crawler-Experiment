package main

import (
	"fmt"
	"sync"
)

type ParallelFetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// use a sync.WaitGroup to wait for all parallel crawls to finish
func ParallelCrawl(url string, depth int, fetcher ParallelFetcher, wg *sync.WaitGroup) {
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
// 	ParallelCrawl("http://golang.org/", 4, parallelFetcher, &wg)
// 	wg.Wait()
// }

// fakeFetcher is Fetcher that returns canned results.
type parallelFakeFetcher map[string]*parallelFakeResult

type parallelFakeResult struct {
	body string
	urls []string
}

func (f parallelFakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var parallelFetcher = parallelFakeFetcher{
	"http://golang.org/": &parallelFakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &parallelFakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &parallelFakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &parallelFakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
