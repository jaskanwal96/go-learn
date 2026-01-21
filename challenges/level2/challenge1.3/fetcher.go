package main

import (
	"io"
	"net/http"
	"sync"
	"time"
)

// ============================================
// Concurrent URL Fetcher
// ============================================
// This file implements a concurrent URL fetcher
// using goroutines, channels, and WaitGroups.
//
// CHALLENGE: Implement the methods below!
//
// Architecture:
// 1. Main goroutine adds URLs to urlChan
// 2. Worker goroutines fetch URLs from urlChan
// 3. Workers send results to resultsChan
// 4. Collector goroutine processes results
// 5. WaitGroup ensures all workers finish

// Fetcher manages concurrent URL fetching
type Fetcher struct {
	workers     int
	urlChan     chan URLRequest
	resultsChan chan FetchResult
	stats       *Stats
	workerWg    sync.WaitGroup // Waits for workers to finish
	collectorWg sync.WaitGroup // Waits for collector to finish
}

// NewFetcher creates a new fetcher with specified number of workers
func NewFetcher(workers int) *Fetcher {
	return &Fetcher{
		workers:     workers,
		urlChan:     make(chan URLRequest, workers*2), // Buffered channel
		resultsChan: make(chan FetchResult, workers),  // Buffered for results
		stats:       NewStats(),
	}
}

// AddURL adds a URL to the fetch queue
func (f *Fetcher) AddURL(url string, id int) {
	f.urlChan <- URLRequest{URL: url, ID: id}
}

func (f *Fetcher) Start() {
	// TODO: Implement me!
}

// worker is a goroutine that fetches URLs from the channel
// TODO: Implement this method
// HINT:
// 1. Use defer f.workerWg.Done() at the start
// 2. Loop: for req := range f.urlChan
// 3. Print a message showing which URL you're fetching
// 4. Call f.fetchURL(req) to get the result
// 5. Send result to f.resultsChan
// 6. Print when worker finishes
func (f *Fetcher) worker(id int) {
	// TODO: Implement me!
}

// fetchURL performs the actual HTTP GET request
// This one is implemented for you since it's HTTP-specific
func (f *Fetcher) fetchURL(req URLRequest) FetchResult {
	startTime := time.Now()

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(req.URL)
	if err != nil {
		return FetchResult{
			URL:      req.URL,
			Success:  false,
			Duration: time.Since(startTime),
			Error:    err.Error(),
		}
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return FetchResult{
			URL:        req.URL,
			Success:    false,
			StatusCode: resp.StatusCode,
			Duration:   time.Since(startTime),
			Error:      err.Error(),
		}
	}

	return FetchResult{
		URL:        req.URL,
		Success:    resp.StatusCode == 200,
		StatusCode: resp.StatusCode,
		Size:       len(body),
		Duration:   time.Since(startTime),
	}
}

// collector processes results from workers
// TODO: Implement this method
// HINT:
// 1. Use defer f.collectorWg.Done() at the start
// 2. Loop: for result := range f.resultsChan
// 3. If result.Success is true:
//   - Print success message with URL, status, size, time
//   - Call f.stats.RecordSuccess(...)
//
// 4. If result.Success is false:
//   - Print failure message with URL, error, time
//   - Call f.stats.RecordFailure(...)
//
// 5. Print when collector finishes
func (f *Fetcher) collector() {
	// TODO: Implement me!
}

// FinishAdding signals that no more URLs will be added
func (f *Fetcher) FinishAdding() {
	close(f.urlChan)
}

// Wait waits for all workers and collector to finish
func (f *Fetcher) Wait() {
	f.collectorWg.Wait()
}

// GetStats returns the final statistics
func (f *Fetcher) GetStats() StatsSnapshot {
	return f.stats.GetSnapshot()
}
