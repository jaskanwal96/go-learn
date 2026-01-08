package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// 🎯 CHALLENGE 2.2: Worker Pool Pattern
// ============================================
// This challenge teaches:
// - Worker pools (limited concurrent workers)
// - Job queues
// - Mutex for shared state
// - Context for cancellation
//
// TASK: Build a worker pool that processes jobs concurrently
//
// Requirements:
// 1. Create a Job struct with ID and Data
// 2. Create a WorkerPool struct
// 3. Implement ProcessJob(job Job) - simulates work
// 4. Use a channel for job queue
// 5. Limit number of concurrent workers (e.g., 3)
// 6. Track completed jobs with mutex
//
// Concepts you'll practice:
// - Worker pool pattern
// - sync.Mutex for thread-safe operations
// - Channel-based job queue
// - Controlling concurrency

type Job struct {
	ID   int
	Data string
}

type WorkerPool struct {
	workers    int
	jobQueue   chan Job
	completed  int
	mu         sync.Mutex
	wg         sync.WaitGroup
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(workers int, queueSize int) *WorkerPool {
	// TODO: Initialize worker pool
	return nil
}

// Start starts the worker pool
func (wp *WorkerPool) Start() {
	// TODO: Start worker goroutines
	// Each worker should:
	// 1. Loop receiving jobs from jobQueue
	// 2. Process each job
	// 3. Increment completed counter (with mutex)
}

// AddJob adds a job to the queue
func (wp *WorkerPool) AddJob(job Job) {
	// TODO: Send job to jobQueue
}

// ProcessJob simulates processing a job
func (wp *WorkerPool) ProcessJob(job Job) {
	// TODO: Simulate work (time.Sleep(100ms))
	// TODO: Print job processing
	// TODO: Increment completed (with mutex lock)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Processing job %d: %s\n", job.ID, job.Data)
	wp.mu.Lock()
	wp.completed++
	wp.mu.Unlock()
}

// Wait waits for all jobs to complete
func (wp *WorkerPool) Wait() {
	// TODO: Close jobQueue and wait for workers
}

// GetCompleted returns number of completed jobs
func (wp *WorkerPool) GetCompleted() int {
	// TODO: Return completed (with mutex lock)
	return 0
}

func main() {
	fmt.Println("🎮 Challenge 2.2: Worker Pool Pattern")
	fmt.Println("====================================\n")

	// Create worker pool with 3 workers
	pool := NewWorkerPool(3, 10)
	pool.Start()

	// Add jobs
	for i := 1; i <= 10; i++ {
		job := Job{
			ID:   i,
			Data: fmt.Sprintf("Task-%d", i),
		}
		pool.AddJob(job)
	}

	// Wait for completion
	pool.Wait()

	fmt.Printf("✅ Completed %d jobs\n", pool.GetCompleted())
}

