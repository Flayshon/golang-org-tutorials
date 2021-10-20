package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type CrawlerCache struct {
	mu sync.Mutex
	cache  map[string]int
}

func (cc *CrawlerCache) Crawl(url string, depth int, fetcher Fetcher) {

	defer wg.Done()

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		cc.mu.Lock()
		cc.cache[url] = 1
		cc.mu.Unlock()

		fmt.Println(err)
		return
	} else {
		cc.mu.Lock()
		cc.cache[url] = 1
		cc.mu.Unlock()
	}


	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if cc.cache[u] != 1 {
			wg.Add(1)
			go cc.Crawl(u, depth-1, fetcher)
		}
	}
}

var wg sync.WaitGroup

func main() {
	cc := CrawlerCache{cache: make(map[string]int)}
	wg.Add(1)
	go cc.Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
}


type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}


var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
