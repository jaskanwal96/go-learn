package main

import "time"

// ============================================
// URL and Result Definitions
// ============================================
// This file defines the data structures used
// for URL fetching and result tracking.

// URLRequest represents a URL to fetch
type URLRequest struct {
	URL string
	ID  int
}

// FetchResult represents the result of fetching a URL
type FetchResult struct {
	URL        string
	Success    bool
	StatusCode int
	Size       int           // Size of response body in bytes
	Duration   time.Duration // Time taken to fetch
	Error      string        // Error message if failed
}
