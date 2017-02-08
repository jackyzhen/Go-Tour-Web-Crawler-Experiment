# Go-Tour-Web-Crawler-Experiment

Implementations of [https://tour.golang.org/concurrency/10](https://tour.golang.org/concurrency/10) and some speed comparisons.

### Usage


```
bash-3.2$ go run experiment.go webCrawlerCache.go webCrawlerParallel.go webCrawlerNaive.go webCrawlerParallelCache.go common.go onlineSoluti
on.go
NAIVE CRAWL: 100000 repetitions took 237.59997ms
NAIVE CACHE: 100000 repetitions took 17.067391ms
PARALLEL CRAWL: 100000 repetitions took 2.569889913s
PARALLEL CACHE CRAWL: 100000 repetitions took 161.429999ms
ONLINE CRAWL: 100000 repetitions took 754.953116ms
bash-3.2$
```