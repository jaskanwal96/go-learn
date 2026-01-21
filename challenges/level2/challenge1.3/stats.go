package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// Thread-Safe Statistics Collector
// ============================================
// This file implements a statistics collector
// that safely tracks metrics from multiple
// goroutines using a mutex.
//
// CHALLENGE: Implement the methods below!
//
// Key concepts:
// - sync.Mutex protects shared data
// - Always Lock() before accessing shared data
// - Always Unlock() after (use defer!)
// - Multiple goroutines will call these methods

// Stats tracks fetching statistics
type Stats struct {
	mu          sync.Mutex
	total       int           // Total URLs processed
	successful  int           // Successfully fetched
	failed      int           // Failed to fetch
	totalBytes  int64         // Total bytes downloaded
	totalTime   time.Duration // Total time spent fetching
	fastestURL  string        // Fastest URL fetched
	fastestTime time.Duration // Time for fastest fetch
	slowestURL  string        // Slowest URL fetched
	slowestTime time.Duration // Time for slowest fetch
}

// NewStats creates a new stats collector
func NewStats() *Stats {
	return &Stats{
		fastestTime: time.Hour, // Initialize to large value
	}
}

// RecordSuccess records a successful fetch
// TODO: Implement this method
// HINT:
// 1. Lock the mutex (s.mu.Lock())
// 2. Use defer to unlock (defer s.mu.Unlock())
// 3. Update: total, successful, totalBytes, totalTime
// 4. Track fastest URL (if duration < s.fastestTime)
// 5. Track slowest URL (if duration > s.slowestTime)
func (s *Stats) RecordSuccess(url string, size int, duration time.Duration) {
	// TODO: Implement me!
}

// RecordFailure records a failed fetch
// TODO: Implement this method
// HINT:
// 1. Lock the mutex
// 2. Update: total, failed, totalTime
// 3. Unlock the mutex
func (s *Stats) RecordFailure(duration time.Duration) {
	// TODO: Implement me!
}

// GetSnapshot returns a copy of current statistics
// TODO: Implement this method
// HINT:
// 1. Lock the mutex
// 2. Create and return a StatsSnapshot with all current values
// 3. Unlock the mutex (use defer!)
// Returns a copy to avoid holding the lock while caller uses the data
func (s *Stats) GetSnapshot() StatsSnapshot {
	// TODO: Implement me!
	return StatsSnapshot{}
}

// StatsSnapshot is a point-in-time copy of statistics
type StatsSnapshot struct {
	Total       int
	Successful  int
	Failed      int
	TotalBytes  int64
	TotalTime   time.Duration
	FastestURL  string
	FastestTime time.Duration
	SlowestURL  string
	SlowestTime time.Duration
}

// Print displays the statistics nicely
func (s StatsSnapshot) Print() {
	fmt.Println("\n📊 Fetching Statistics")
	fmt.Println("=====================")
	fmt.Printf("Total URLs:     %d\n", s.Total)
	fmt.Printf("✅ Successful:  %d\n", s.Successful)
	fmt.Printf("❌ Failed:      %d\n", s.Failed)
	fmt.Printf("📦 Total Bytes: %d (%.2f KB)\n", s.TotalBytes, float64(s.TotalBytes)/1024)
	fmt.Printf("⏱️  Total Time:  %v\n", s.TotalTime)

	if s.Successful > 0 {
		avgTime := s.TotalTime / time.Duration(s.Total)
		fmt.Printf("⏱️  Avg Time:    %v\n", avgTime)
		fmt.Printf("🚀 Fastest:     %s (%v)\n", s.FastestURL, s.FastestTime)
		fmt.Printf("🐌 Slowest:     %s (%v)\n", s.SlowestURL, s.SlowestTime)
	}
}
