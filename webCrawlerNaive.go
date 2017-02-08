// base implementation with no parallel and no cache.

package main

// NaiveCrawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func NaiveCrawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
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
		NaiveCrawl(u, depth-1, fetcher)
	}
	return
}

// func main() {
//   now := time.Now().Nanosecond()
// 	NaiveCrawl("http://golang.org/", 4, fetcher)
//   after := time.Now().Nanosecond()
//   fmt.Println("Time Elapsed:", after - now)
// }

