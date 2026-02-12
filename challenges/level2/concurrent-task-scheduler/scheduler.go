package main

import (
	"fmt"
	"sync"
	"time"
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

	queue := NewPriorityQueue()
	channel := make(chan Task, workers*2)
	return &Scheduler{
		queue:     queue,
		taskChan:  channel,
		workers:   workers,
		completed: 0,
	}
}

// AddTask adds a task to the scheduler queue
func (s *Scheduler) AddTask(task Task) {
	// TODO: Add task to priority queue
	//
	// Steps:
	// 1. Call Push method on the priority queue
	//    Example: s.queue.Push(task)
	s.queue.Push(task)
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
	for i := 0; i < s.workers; i++ {
		workerID := i + 1 // Worker IDs start from 1
		s.wg.Add(1)
		go func(id int) {
			defer s.wg.Done()
			for task := range s.taskChan {
				s.processTask(task, id)
			}
		}(workerID)
	}
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
	go func() {
		for !s.queue.IsEmpty() {
			task, ok := s.queue.Pop()
			if ok {
				s.taskChan <- task
			}
		}
		close(s.taskChan)
	}()
}

// processTask simulates processing a task
func (s *Scheduler) processTask(task Task, workerID int) {
	fmt.Printf("👷 Worker %d: 🚀 Starting task %d: %s (Priority: %d)\n", workerID, task.ID, task.Name, task.Priority)
	time.Sleep(task.Duration)
	fmt.Printf("👷 Worker %d: ✅ Completed task %d: %s\n", workerID, task.ID, task.Name)
	s.completedMu.Lock()
	s.completed++
	s.completedMu.Unlock()
}

// Wait waits for all tasks to complete
func (s *Scheduler) Wait() {
	s.wg.Wait()
}

// GetCompleted returns the number of completed tasks
func (s *Scheduler) GetCompleted() int {
	s.completedMu.Lock()
	count := s.completed
	s.completedMu.Unlock()
	return count
}
