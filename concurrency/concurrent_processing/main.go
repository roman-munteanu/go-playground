package main

import (
	"fmt"
	"sync"
	"time"
)

type Status string

const (
	Success Status = "success"
	Failure        = "failure"
)

type Song struct {
	Title  string
	Artist string
}

type Platform struct {
	Name string
}

type ReleaseRequest struct {
	Song     Song
	Platform Platform
}

type ReleaseResult struct {
	Req    ReleaseRequest
	Status Status
	Error  error
}

type PlatformResults struct {
	Results []ReleaseResult
}

func main() {
	song := Song{
		Title:  "Sunshine",
		Artist: "RMHighlander",
	}
	platforms := []Platform{
		{Name: "Spotify"},
		{Name: "SoundCloud"},
		{Name: "Apple Music"},
		{Name: "YouTube Music"},
	}

	in := make(chan ReleaseResult)
	out := make(chan PlatformResults)
	defer close(out)

	var wg sync.WaitGroup
	for _, platform := range platforms {
		wg.Add(1)
		req := ReleaseRequest{
			Song:     song,
			Platform: platform,
		}
		go processRelease(req, in)
	}

	go handleResults(in, out, &wg)

	wg.Wait()
	close(in)

	results := <-out
	fmt.Printf("Results: %v\n", results)
}

func releaseToPlatform(req ReleaseRequest) ReleaseResult {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Released song: `%v` to platform: `%v`\n", req.Song, req.Platform)

	return ReleaseResult{
		Req:    req,
		Status: Success,
		Error:  nil,
	}
}

func processRelease(req ReleaseRequest, in chan ReleaseResult) {
	in <- releaseToPlatform(req)
}

func handleResults(in chan ReleaseResult, out chan PlatformResults, wg *sync.WaitGroup) {
	var results PlatformResults
	for res := range in {
		results.Results = append(results.Results, res)
		wg.Done()
	}
	out <- results
}
