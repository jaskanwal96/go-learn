package main

import (
	"sync"
)

// ============================================
// Scheduler Implementation
// ============================================
// This file implements a concurrent task scheduler that:
// - Uses goroutines to process tasks concurrently
// - Uses channels to communicate between dispatcher and workers
// - Uses WaitGroups to synchronize completion
//
// Architecture:
// - Dispatcher goroutine: Pops tasks from priority queue and sends to channel
// - Worker goroutines: Receive tasks from channel and process them
// - WaitGroup: Ensures all workers finish before program exits
//
// HINT: Review challenge1 to understand:
//   - How goroutines work with channels
//   - How WaitGroups synchronize goroutines
//   - How to send/receive from channels

// Scheduler manages concurrent task processing
type Scheduler struct {
	queue       *PriorityQueue
	taskChan    chan Task      // Channel for sending tasks to workers
	workers     int            // Number of worker goroutines
	wg          sync.WaitGroup // Waits for all workers to finish
	completed   int            // Counter for completed tasks
	completedMu sync.Mutex     // Protects completed counter
}

// NewScheduler creates a new scheduler with specified number of workers
func NewScheduler(workers int) *Scheduler {
	// TODO: Initialize scheduler
	//
	// Steps:
	// 1. Create new PriorityQueue using NewPriorityQueue()
	// 2. Create buffered channel for tasks (size = workers * 2)
	//    Example: make(chan Task, workers*2)
	// 3. Set workers count
	// 4. Initialize WaitGroup (it's a struct, no initialization needed)
	// 5. Initialize completed to 0
	// 6. Initialize completedMu (it's a struct, no initialization needed)
	// 7. Return pointer to Scheduler
	return nil
}

// AddTask adds a task to the scheduler queue
func (s *Scheduler) AddTask(task Task) {
	// TODO: Add task to priority queue
	//
	// Steps:
	// 1. Call Push method on the priority queue
	//    Example: s.queue.Push(task)
}

// Start begins processing tasks with worker goroutines
func (s *Scheduler) Start() {
	// TODO: Start worker goroutines and dispatcher
	//
	// Part 1: Start worker goroutines
	// For each worker (loop from 0 to s.workers):
	//   1. Add to WaitGroup: s.wg.Add(1)
	//   2. Start goroutine that:
	//      a. Defers wg.Done() to signal completion when done
	//      b. Loops receiving tasks from taskChan:
	//         for task := range s.taskChan {
	//             s.processTask(task)
	//         }
	//      c. When channel closes, loop exits and defer calls wg.Done()
	//
	// Part 2: Start dispatcher goroutine
	// Start a goroutine that:
	//   1. Loops while queue is not empty:
	//      for !s.queue.IsEmpty() {
	//          task, ok := s.queue.Pop()
	//          if ok {
	//              s.taskChan <- task
	//          }
	//      }
	//   2. Close taskChan when queue is empty: close(s.taskChan)
	//
	// HINT: Review challenge1's processResults function for the worker pattern
	// HINT: The dispatcher should run in a separate goroutine
}

// processTask simulates processing a task
func (s *Scheduler) processTask(task Task) {
	// TODO: Process a task
	//
	// NOTE: You'll need to add these imports at the top:
	//   import (
	//       "fmt"
	//       "sync"
	//       "time"
	//   )
	//
	// Steps:
	// 1. Print that task is starting (include priority and name)
	//    Example: fmt.Printf("🚀 Starting task %d: %s (Priority: %d)\n", ...)
	// 2. Sleep for task.Duration to simulate work: time.Sleep(task.Duration)
	// 3. Print that task is completed
	//    Example: fmt.Printf("✅ Completed task %d: %s\n", ...)
	// 4. Increment completed counter (with mutex lock):
	//    s.completedMu.Lock()
	//    s.completed++
	//    s.completedMu.Unlock()
}

// Wait waits for all tasks to complete
func (s *Scheduler) Wait() {
	// TODO: Wait for all worker goroutines to finish
	//
	// Steps:
	// 1. Call Wait() on the WaitGroup
	//    Example: s.wg.Wait()
	//
	// This will block until all workers have called wg.Done()
}

// GetCompleted returns the number of completed tasks
func (s *Scheduler) GetCompleted() int {
	// TODO: Return completed count (with mutex lock)
	//
	// Steps:
	// 1. Lock the mutex: s.completedMu.Lock()
	// 2. Get the value: count := s.completed
	// 3. Unlock the mutex: s.completedMu.Unlock()
	// 4. Return the count
	return 0
}
