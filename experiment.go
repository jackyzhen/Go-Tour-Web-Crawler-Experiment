package main


import (
  "fmt"
  "time"
	"sync"
)
const repetitions = 100000


func naiveCrawl() {
  start := time.Now()
  for i := 0; i < repetitions; i ++ {
    NaiveCrawl("http://golang.org/", 4, fetcher)
  }
  elapsed := time.Since(start)
  fmt.Printf("NAIVE CRAWL: %d repetitions took %s\n", repetitions, elapsed)
}
func naiveCacheCrawl() {
  start := time.Now()
  for i := 0; i < repetitions; i ++ {
  	NaiveCacheCrawl("http://golang.org/", 4, naiveCacheFetcher, naiveCache)
  }
  elapsed := time.Since(start)
  fmt.Printf("NAIVE CACHE: %d repetitions took %s\n", repetitions, elapsed)

}
func parallelCrawl() {
  start := time.Now()
  var wg sync.WaitGroup
  for i := 0; i < repetitions; i ++ {
    wg.Add(1)
    go ParallelCrawl("http://golang.org/", 4, &parallelFetcher, &wg)
  }
  wg.Wait()
  elapsed := time.Since(start)
  fmt.Printf("PARALLEL CRAWL: %d repetitions took %s\n", repetitions, elapsed)
}


func main() {
  naiveCrawl()
  naiveCacheCrawl()
  parallelCrawl()
}