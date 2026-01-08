package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// 🎯 CHALLENGE 2.1: Goroutines & Channels
// ============================================
// This challenge teaches:
// - Goroutines (concurrent functions)
// - Channels (communication between goroutines)
// - WaitGroups (synchronization)
// - Select statements
//
// TASK: Build a concurrent web scraper simulator
//
// Requirements:
// 1. Create a function fetchURL(url string, ch chan<- string)
//    - Simulates fetching a URL (use time.Sleep)
//    - Sends result to channel
// 2. Create a function processResults(ch <-chan string, wg *sync.WaitGroup)
//    - Receives results from channel
//    - Processes them concurrently
// 3. Use WaitGroup to wait for all goroutines
// 4. Use buffered channels
//
// Concepts you'll practice:
// - go keyword for goroutines
// - Channel communication (send/receive)
// - sync.WaitGroup
// - Channel direction (chan<- and <-chan)

type Result struct {
	URL     string
	Content string
	Time    time.Duration
}

// fetchURL simulates fetching a URL
// Sends result to channel when done
func fetchURL(url string, ch chan<- Result) {
	// TODO: Implement this
	// 1. Record start time
	// 2. Simulate network delay (time.Sleep(100-500ms))
	// 3. Create Result struct
	// 4. Send to channel
	start := time.Now()
	delay := time.Duration(100+len(url)*10) * time.Millisecond
	time.Sleep(delay)

	result := Result{
		URL:     url,
		Content: fmt.Sprintf("Content from %s", url),
		Time:    time.Since(start),
	}
	ch <- result
}

// processResults receives results from channel and processes them
func processResults(ch <-chan Result, wg *sync.WaitGroup) {
	// TODO: Implement this
	// 1. Loop receiving from channel
	// 2. Process each result (print it)
	// 3. Call wg.Done() when channel closes
	defer wg.Done()
	for result := range ch {
		fmt.Printf("✅ Fetched %s in %v\n", result.URL, result.Time)
	}
}

func main() {
	fmt.Println("🎮 Challenge 2.1: Goroutines & Channels")
	fmt.Println("======================================\n")

	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://reddit.com",
	}

	// TODO: Create a buffered channel for results
	// TODO: Create a WaitGroup
	// TODO: Start goroutines to fetch URLs
	// TODO: Start goroutine to process results
	// TODO: Wait for all to complete

	ch := make(chan Result, len(urls))
	var wg sync.WaitGroup

	// Start fetchers
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			fetchURL(u, ch)
		}(url)
	}

	// Start processor
	wg.Add(1)
	go processResults(ch, &wg)

	// Wait for all fetchers to complete, then close channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Wait a bit for processing
	time.Sleep(2 * time.Second)

	fmt.Println("\n✅ Challenge complete!")
}
