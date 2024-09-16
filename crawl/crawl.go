package crawl

import (
	"fmt"
	"mutex/cache"
	"mutex/fetcher"
	"sync"
)

type CachedPage struct {
	page *fetcher.FakeResult
	err  error
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetch fetcher.Fetcher, safeCache cache.IMap[CachedPage], wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer wg.Done()
	if depth <= 0 {
		return
	}

	cachedResults, found := safeCache.Get(url)
	if found {
		if cachedResults.err != nil {
			fmt.Println("cache", cachedResults.err)
			return
		} else {
			fmt.Printf("cache found: %s %q\n", url, cachedResults.page.Body)
			for _, u := range cachedResults.page.Urls {
				wg.Add(1)
				Crawl(u, depth-1, fetch, safeCache, wg)
			}
			return
		}
	}

	body, urls, err := fetch.Fetch(url)

	if !found {
		safeCache.Set(url, CachedPage{page: &fetcher.FakeResult{Body: body, Urls: urls}, err: err})
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		Crawl(u, depth-1, fetch, safeCache, wg)
	}
}
