package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeUrlCache struct {
	cache_url map[string]int
	mux       sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, sucache *SafeUrlCache, ch chan int) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	//fmt.Printf("Entering Crawl : url=%s\n", url)
	if depth <= 0 {
		ch <- 3
		return
	}
	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		ch <- 0
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	sucache.mux.Lock()
	// base url has been crawled so invalidate that entry in cache
	sucache.cache_url[url] = 2
	for _, u := range urls {
		//fmt.Println(u)
		_, ok := sucache.cache_url[u]
		if ok == false {
			sucache.cache_url[u] = 1
		} else {
			//fmt.Printf(" Already Crawled url: %s \n", u)
			sucache.cache_url[u]++
			/*
				sucache.mux.Unlock()
				ch <- 1
				return
			*/
		}
	}
	sucache.mux.Unlock()

	for key := range sucache.cache_url {
		//fmt.Println(key)
		if sucache.cache_url[key] == 1 {
			go Crawl(key, depth-1, fetcher, sucache, ch)
			x := <-ch
			fmt.Printf("thread exist rc=%d\n", x)
		}
	}
	ch <- 4
	return
}

func main() {
	surl := SafeUrlCache{cache_url: make(map[string]int)}
	ch := make(chan int)
	Crawl("http://golang.org/", 4, fetcher, &surl, ch)
	//x := <-ch
	//fmt.Printf("main thread exited rc=%d\n", x)
	fmt.Printf("main thread exited rc=\n")
}

// fakeFetcher is Fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
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
}
