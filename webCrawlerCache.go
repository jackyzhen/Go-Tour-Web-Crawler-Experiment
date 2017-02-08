package main

// NaiveCacheCrawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// use a cache to prevent refetching urls
func NaiveCacheCrawl(url string, depth int, fetcher Fetcher, cache map[string]string) {
	if depth <= 0 {
		return
	}
  
	_, urls, err := fetcher.Fetch(url)
  cache[url] = "cached"
	if err != nil {
		// fmt.Println(err)
		return
	}
	// fmt.Printf("found: %s %q\n", url, body)
  
	for _, u := range urls {
    if _, ok := cache[u]; !ok {
		  NaiveCacheCrawl(u, depth-1, fetcher, cache)
    }
	}
	return
}

// func main() {
//   now := time.Now().Nanosecond()
// 	NaiveCacheCrawl("http://golang.org/", 4, fetcher, naiveCache)
//   after := time.Now().Nanosecond()
//   fmt.Println("Time Elapsed:", after - now)
// }