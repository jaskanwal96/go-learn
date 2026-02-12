package main

import (
	"fmt"
	"time"
	// TODO: Add imports as needed:
	// "encoding/json"
	// "log"
	// "net/http"
	// "sync"
	// "time"
)

// ============================================
// 🎯 CHALLENGE 5.1: Integrated System
// ============================================
// This is a REAL-WORLD challenge combining:
// - Goroutines for concurrent API calls
// - HTTP client to fetch external data
// - Database to store results
// - REST API to serve data
// - Error handling and logging
//
// TASK: Build a data aggregation service
//
// Requirements:
// 1. Fetch data from multiple external APIs concurrently (goroutines)
// 2. Store fetched data in database
// 3. Create REST API to query stored data
// 4. Use worker pool pattern for API calls
// 5. Implement caching mechanism
// 6. Add proper error handling
//
// Architecture:
//   External APIs → Worker Pool → Database → REST API → Client
//
// Concepts you'll practice:
// - Goroutines and channels
// - HTTP clients
// - Database operations
// - REST API design
// - Concurrent programming patterns
// - Error handling strategies

type ExternalData struct {
	ID      int       `json:"id"`
	Source  string    `json:"source"`
	Content string    `json:"content"`
	Fetched time.Time `json:"fetched"`
}

type DataAggregator struct {
	// TODO: Add fields
	// - HTTP client
	// - Database connection
	// - Cache (map with mutex)
	// - Worker pool
}

// TODO: Implement all methods
// - FetchFromAPI(source string) (*ExternalData, error)
// - StoreInDB(data *ExternalData) error
// - GetFromCache(key string) (*ExternalData, bool)
// - FetchMultipleSources(sources []string) error
// - StartAPIServer() error

func main() {
	fmt.Println("🎮 Challenge 5.1: Integrated System")
	fmt.Println("===================================\n")

	// TODO: Initialize aggregator
	// aggregator := NewDataAggregator()

	// TODO: Fetch from multiple sources concurrently
	// sources := []string{"api1", "api2", "api3"}
	// aggregator.FetchMultipleSources(sources)

	// TODO: Start REST API server
	// aggregator.StartAPIServer()

	fmt.Println("✅ Challenge complete!")
	fmt.Println("This is a complex challenge - take your time!")
}
