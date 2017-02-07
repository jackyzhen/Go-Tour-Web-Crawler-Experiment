package main/parallelCache

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

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

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	Crawl("http://golang.org/", 4, &fetcher, &wg)
	wg.Wait()  
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher struct {
	v   map[string]*fakeResult
	mux sync.Mutex
}

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Value(key string) (*fakeResult, bool) {
	f.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer f.mux.Unlock()
	elem, ok := f.v[key]
  return elem, ok
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f.Value(url); ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{v: map[string]*fakeResult{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}}
