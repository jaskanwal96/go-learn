package main

import (
	"fmt"
	"time"
)

// ============================================
// 🎯 CHALLENGE 2.3: Concurrent URL Fetcher
// ============================================
// This challenge teaches:
// - Goroutines (concurrent HTTP requests)
// - Channels (URL queue and results)
// - WaitGroups (synchronization)
// - Mutex (thread-safe statistics)
// - Real-world patterns (worker pools, rate limiting)
//
// TASK: Build a concurrent URL fetcher that processes multiple URLs
//
// File Structure:
// - url.go: Defines URLRequest and FetchResult structs
// - stats.go: Thread-safe statistics collector using mutex
// - fetcher.go: Concurrent fetcher with goroutines and channels
// - main.go: This file - sets up and runs the challenge
//
// Requirements:
// 1. Implement Stats collector with mutex protection
// 2. Create Fetcher that uses goroutines to fetch URLs concurrently
// 3. Use channels to send URLs to workers and collect results
// 4. Use WaitGroup to wait for all work to complete
// 5. Handle errors gracefully
//
// Concepts you'll practice:
// - Worker pool pattern
// - Buffered vs unbuffered channels
// - Channel closing and range loops
// - Mutex for protecting shared data
// - Real HTTP requests

func main() {
	fmt.Println("🎮 Challenge 2.3: Concurrent URL Fetcher")
	fmt.Println("========================================")
	fmt.Println()

	startTime := time.Now()

	// Create fetcher with 5 concurrent workers
	fetcher := NewFetcher(5)

	// List of URLs to fetch
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
		"https://www.wikipedia.org",
		"https://www.youtube.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.amazon.com",
		"https://www.netflix.com",
		"https://www.invalid-url-that-will-fail-12345.com", // This will fail
		"https://www.apple.com",
		"https://www.microsoft.com",
		"https://www.golang.org",
		"https://www.rust-lang.org",
	}

	fmt.Printf("📋 Fetching %d URLs with %d workers...\n\n", len(urls), 5)

	// Start the fetcher (launches workers and collector)
	fetcher.Start()

	// Add URLs to the queue
	for i, url := range urls {
		fetcher.AddURL(url, i+1)
	}

	// Signal that no more URLs will be added
	fetcher.FinishAdding()

	// Wait for all work to complete
	fmt.Println("\n⏳ Waiting for all fetches to complete...\n")
	fetcher.Wait()

	// Get and display statistics
	stats := fetcher.GetStats()
	stats.Print()

	fmt.Printf("\n⏱️  Total execution time: %v\n", time.Since(startTime))

	fmt.Println("\n🎓 What you learned:")
	fmt.Println("- Goroutines fetch URLs concurrently (5 at a time)")
	fmt.Println("- Channels coordinate work between goroutines")
	fmt.Println("- WaitGroups ensure all work completes")
	fmt.Println("- Mutex protects shared statistics from race conditions")
	fmt.Println("- Worker pool pattern limits concurrent requests")
}
