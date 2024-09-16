package main

import (
	"mutex/cache"
	"mutex/crawl"
	"mutex/fetcher"
	"sync"
)

var fakeFetcher = fetcher.FakeFetcher{
	"https://golang.org/": &fetcher.FakeResult{
		Body: "The Go Programming Language",
		Urls: []string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fetcher.FakeResult{
		Body: "Packages",
		Urls: []string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fetcher.FakeResult{
		Body: "Package fmt",
		Urls: []string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fetcher.FakeResult{
		Body: "Package os",
		Urls: []string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func main() {
	var wg sync.WaitGroup
	safeCache := cache.CreateSafeCache[crawl.CachedPage]()
	wg.Add(1) // Add the first goroutine (main call)
	go crawl.Crawl("https://golang.org/", 4, fakeFetcher, &safeCache, &wg)
	wg.Wait()
}
