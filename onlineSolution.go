// based on https://gist.github.com/avalanche123/4127253
// implements parallel and cache via channels, not wait groups.

package main

type result struct {
	url, body string
	urls []string
	err error
	depth int
}

// OnlineCrawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func OnlineCrawl(url string, depth int, fetcher Fetcher) {
	results := make(chan *result)
	fetched := make(map[string]bool)
	fetch   := func(url string, depth int) {
		body, urls, err := fetcher.Fetch(url)
		results <- &result{url, body, urls, err, depth}
	}

	go fetch(url, depth)
	fetched[url] = true

	// 1 url is currently being fetched in background, loop while fetching
	for fetching := 1; fetching > 0; fetching-- {
		res := <- results

		// skip failed fetches
		if res.err != nil {
			// fmt.Println(res.err)
			continue
		}

		// fmt.Printf("found: %s %q\n", res.url, res.body)

		// follow links if depth has not been exhausted
		if res.depth > 0 {
			for _, u := range res.urls {
				// don't attempt to re-fetch known url, decrement depth
				if !fetched[u] {
        fetching++
        go fetch(u, res.depth - 1)
					fetched[u] = true
				}
			}
		}
	}

	close(results)
}
