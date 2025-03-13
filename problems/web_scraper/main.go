package main

import (
	"fmt"
	"io"
	"net/http"
)

type FetchResult struct {
	url    string
	length int
	err    string
}

func main() {
	sites := []string{
		"https://go.dev/",
		"https://www.google.com/",
	}

	// var wg sync.WaitGroup
	// for _, site := range sites {
	// 	wg.Add(1)
	// 	go fetchURL(site, &wg)
	// }
	// wg.Wait()

	results := make(chan FetchResult)
	for _, site := range sites {
		go fetchURL(site, results)
	}

	for range sites {
		res := <-results
		fmt.Printf("URL: %s, body length: %d\n", res.url, res.length)
	}

	fmt.Println("All done")
}

func fetchURL(url string, ch chan<- FetchResult) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- FetchResult{
			err: "could not execute get request",
		}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- FetchResult{
			err: "could not read body",
		}
		return
	}

	ch <- FetchResult{
		url:    url,
		length: len(body),
	}
}

/*
func fetchURL(url string, wg *sync.WaitGroup) error {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("could not execute get request")
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read body")
		return err
	}

	fmt.Printf("URL: %s, body length: %d\n", url, len(body))

	return nil
}
*/
