package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// Priority Queue (Min-Heap)
// ============================================

// PriorityQueue is a thread-safe min-heap.
// Lower Priority value = higher priority.
type PriorityQueue struct {
	tasks []Task
	mu    sync.Mutex
}

// NewPriorityQueue creates an empty priority queue.
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		tasks: make([]Task, 0),
	}
}

// Push inserts a task into the priority queue.
func (pq *PriorityQueue) Push(task Task) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	pq.tasks = append(pq.tasks, task)

	// Bubble up
	for i := len(pq.tasks) - 1; i > 0; {
		parent := (i - 1) / 2
		if pq.tasks[i].Priority >= pq.tasks[parent].Priority {
			break
		}
		pq.tasks[i], pq.tasks[parent] = pq.tasks[parent], pq.tasks[i]
		i = parent
	}
}

// Pop removes and returns the highest-priority task.
// Returns (Task{}, false) if the queue is empty.
func (pq *PriorityQueue) Pop() (Task, bool) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	n := len(pq.tasks)
	if n == 0 {
		return Task{}, false
	}

	root := pq.tasks[0]
	last := pq.tasks[n-1]
	pq.tasks = pq.tasks[:n-1]

	if n == 1 {
		return root, true
	}

	pq.tasks[0] = last
	parent := 0

	// Bubble down
	for {
		left := 2*parent + 1
		if left >= len(pq.tasks) {
			break
		}

		right := left + 1
		smaller := left

		if right < len(pq.tasks) &&
			pq.tasks[right].Priority < pq.tasks[left].Priority {
			smaller = right
		}

		if pq.tasks[parent].Priority <= pq.tasks[smaller].Priority {
			break
		}

		pq.tasks[parent], pq.tasks[smaller] =
			pq.tasks[smaller], pq.tasks[parent]

		parent = smaller
	}

	return root, true
}

// IsEmpty returns true if the queue is empty.
func (pq *PriorityQueue) IsEmpty() bool {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return len(pq.tasks) == 0
}

// Size returns the number of tasks in the queue.
func (pq *PriorityQueue) Size() int {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	return len(pq.tasks)
}

// ============================================
// 🧪 Priority Queue Test Suite
// ============================================

func TestPQ() {
	fmt.Println("🧪 Testing Priority Queue Implementation")

	passed := 0
	failed := 0

	// Test 1
	fmt.Println("Test 1: Creating new priority queue...")
	pq := NewPriorityQueue()
	if pq != nil {
		fmt.Println("✅ PASSED")
		passed++
	} else {
		fmt.Println("❌ FAILED")
		failed++
	}
	fmt.Println()

	// Test 2
	fmt.Println("Test 2: IsEmpty on new queue...")
	if pq.IsEmpty() {
		fmt.Println("✅ PASSED")
		passed++
	} else {
		fmt.Println("❌ FAILED")
		failed++
	}
	fmt.Println()

	// Test 3
	fmt.Println("Test 3: Size on new queue...")
	if pq.Size() == 0 {
		fmt.Println("✅ PASSED")
		passed++
	} else {
		fmt.Println("❌ FAILED")
		failed++
	}
	fmt.Println()

	// Test 4
	fmt.Println("Test 4: Pop from empty queue...")
	if _, ok := pq.Pop(); !ok {
		fmt.Println("✅ PASSED")
		passed++
	} else {
		fmt.Println("❌ FAILED")
		failed++
	}
	fmt.Println()

	// Test 5
	fmt.Println("Test 5: Push single task...")
	task1 := Task{ID: 1, Priority: 3, Name: "Task 1", Duration: time.Second}
	pq.Push(task1)
	if pq.Size() == 1 && !pq.IsEmpty() {
		fmt.Println("✅ PASSED")
		passed++
	} else {
		fmt.Println("❌ FAILED")
		failed++
	}
	fmt.Println()

	// Test 6
	fmt.Println("Test 6: Pop single task...")
	if t, ok := pq.Pop(); ok && t.ID == task1.ID {
		fmt.Println("✅ PASSED")
		passed++
	} else {
		fmt.Println("❌ FAILED")
		failed++
	}
	fmt.Println()

	// Test 7
	fmt.Println("Test 7: Priority ordering...")
	tasks := []Task{
		{ID: 1, Priority: 5},
		{ID: 2, Priority: 1},
		{ID: 3, Priority: 3},
		{ID: 4, Priority: 2},
		{ID: 5, Priority: 4},
		{ID: 6, Priority: 1},
	}

	for _, t := range tasks {
		pq.Push(t)
	}

	expected := []int{1, 1, 2, 3, 4, 5}
	okAll := true

	for _, exp := range expected {
		t, _ := pq.Pop()
		if t.Priority != exp {
			okAll = false
			break
		}
	}

	if okAll {
		fmt.Println("✅ PASSED")
		passed++
	} else {
		fmt.Println("❌ FAILED")
		failed++
	}
	fmt.Println()

	fmt.Println("==========================================")
	fmt.Printf("Test Summary: %d passed, %d failed\n", passed, failed)
}
