package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	sites := []string{
		"https://golang.org",
		"https://www.google.com",
	}

	codes, err := retrieve(sites)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Status codes:", codes)
}

func getURL(url string) (*http.Response, error) {
	log.Printf("calling %s", url)
	start := time.Now()
	resp, err := http.Get(url)
	log.Printf("completed calling %s in %s", url, time.Since(start))
	return resp, err
}

func retrieve(sites []string) ([]int, error) {
	log.Println("started calling sites")
	start := time.Now()
	wg := &sync.WaitGroup{}

	var responses []int
	errC := &ErrorContainer{}
	for _, v := range sites {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := getURL(url)
			if err != nil {
				errC.Add(err)
				return
			}
			responses = append(responses, resp.StatusCode)
		}(v)
	}

	wg.Wait()
	if errC.IsValid() {
		return responses, nil
	}

	log.Printf("completed calling sites in %s", time.Since(start))
	return responses, errC
}

type ErrorContainer struct {
	Errors []string
}

func (c *ErrorContainer) Add(err error) {
	c.Errors = append(c.Errors, err.Error())
}

func (c *ErrorContainer) IsValid() bool {
	return len(c.Errors) == 0
}

func (c *ErrorContainer) Error() string {
	return fmt.Sprintf("All errors: %s", strings.Join(c.Errors, " "))
}
