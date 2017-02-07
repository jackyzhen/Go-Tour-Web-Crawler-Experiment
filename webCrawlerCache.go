package main

import (
	"fmt"
)

type NaiveCacheFetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// use a cache to prevent refetching urls
func NaiveCacheCrawl(url string, depth int, fetcher NaiveCacheFetcher, naiveCache map[string]string) {
	if depth <= 0 {
		return
	}
  
	_, urls, err := fetcher.Fetch(url)
  naiveCache[url] = "cached"
	if err != nil {
		// fmt.Println(err)
		return
	}
	// fmt.Printf("found: %s %q\n", url, body)
  
	for _, u := range urls {
    if _, ok := naiveCache[u]; !ok {
		  NaiveCacheCrawl(u, depth-1, fetcher, naiveCache)
    }
	}
	return
}

// func main() {
//   now := time.Now().Nanosecond()
// 	NaiveCacheCrawl("http://golang.org/", 4, naiveCacheFetcher, naiveCache)
//   after := time.Now().Nanosecond()
//   fmt.Println("Time Elapsed:", after - now)
// }

// fakeFetcher is Fetcher that returns canned results.
type naiveCacheFakeFetcher map[string]*naiveCacheFakeResult

type naiveCacheFakeResult struct {
	body string
	urls []string
}

func (f naiveCacheFakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var naiveCache = make(map[string]string)
// fetcher is a populated fakeFetcher.
var naiveCacheFetcher = naiveCacheFakeFetcher{
	"http://golang.org/": &naiveCacheFakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &naiveCacheFakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &naiveCacheFakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &naiveCacheFakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
