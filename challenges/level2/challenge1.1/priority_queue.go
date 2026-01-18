package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// Priority Queue (Min-Heap) Implementation
// ============================================
// This file implements a priority queue using a min-heap.
// A min-heap ensures the highest priority task (lowest number)
// is always at the root and can be retrieved quickly.
//
// Key concepts:
// - Heap: A complete binary tree stored in an array
// - Min-heap: Parent priority ≤ child priority
// - Bubble Up: When adding, move element up if it has higher priority
// - Bubble Down: When removing, move element down to maintain heap property
//
// Heap properties:
// - Parent index: (child_index - 1) / 2
// - Left child: 2 * index + 1
// - Right child: 2 * index + 2

// PriorityQueue is a min-heap for tasks (lower priority number = higher priority)
type PriorityQueue struct {
	tasks []Task
	mu    sync.Mutex // Protects concurrent access to the queue
}

// NewPriorityQueue creates a new priority queue
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{tasks: []Task{}}
}

// Push adds a task to the priority queue
func (pq *PriorityQueue) Push(task Task) {
	// TODO: Implement heap push operation
	pq.mu.Lock()
	pq.tasks = append(pq.tasks, task)
	// 	Index:  0   1   2    3   4    5    6
	// Value:  1 | 4 | 15 | 20| 50 | 100| 10
	// 	     1
	//     /   \
	//    4     15
	//   / \   /  \
	//  20 50 100 10
	child_index := len(pq.tasks) - 1
	parent_index := (child_index - 1) / 2
	if child_index == 0 {
		pq.mu.Unlock()
		return
	}
	for pq.tasks[child_index].Priority < pq.tasks[parent_index].Priority {
		pq.tasks[child_index], pq.tasks[parent_index] = pq.tasks[parent_index], pq.tasks[child_index]
		child_index = parent_index
		parent_index = (child_index - 1) / 2
	}
	pq.mu.Unlock()
	//
	// Steps:
	// 1. Lock the mutex to ensure thread-safe access
	// 2. Append task to the end of tasks slice
	// 3. Bubble up to maintain heap property (min-heap):
	//    - Start from the newly added element (last index)
	//    - Compare with parent: if task.Priority < parent.Priority, swap
	//    - Move up to parent and repeat until heap property is satisfied
	//    - Stop when you reach root (index 0) or parent has higher priority
	// 4. Unlock the mutex
	//
	// HINT: Parent index = (child_index - 1) / 2
	// HINT: Use integer division: (index - 1) / 2
}

// Pop removes and returns the highest priority task (lowest priority number)
func (pq *PriorityQueue) Pop() (Task, bool) {
	// TODO: Implement heap pop operation
	//
	// Steps:
	// 1. Lock the mutex
	// 2. Check if queue is empty (len(tasks) == 0), return (Task{}, false) if empty
	// 3. Get root task (tasks[0]) - this is the highest priority
	// 4. Move last element to root: tasks[0] = tasks[len(tasks)-1]
	// 5. Remove last element: tasks = tasks[:len(tasks)-1]
	// 6. If queue is not empty, bubble down to maintain heap property:
	//    - Start from root (index 0)
	//    - Compare with children: swap with smaller priority child
	//    - Move down to swapped child and repeat until heap property is satisfied
	// 7. Unlock the mutex
	// 8. Return the saved root task and true
	//
	// HINT: Left child = 2*index + 1, Right child = 2*index + 2
	// HINT: Check bounds before accessing children (left < len(tasks))
	pq.mu.Lock()
	if len(pq.tasks) == 0 {
		pq.mu.Unlock()
		return Task{}, false
	}
	rootTask := pq.tasks[0]
	pq.tasks[0] = pq.tasks[len(pq.tasks)-1]
	pq.tasks = pq.tasks[:len(pq.tasks)-1]
	parent_index := 0
	left_child_index := 2*parent_index + 1
	right_child_index := 2*parent_index + 2

	for left_child_index < len(pq.tasks) && pq.tasks[left_child_index].Priority < pq.tasks[parent_index].Priority ||
		right_child_index < len(pq.tasks) && pq.tasks[right_child_index].Priority < pq.tasks[parent_index].Priority {
		smaller_index := 0
		if pq.tasks[left_child_index].Priority < pq.tasks[right_child_index].Priority {
			smaller_index = left_child_index
		} else {
			smaller_index = right_child_index
		}
		pq.tasks[smaller_index], pq.tasks[parent_index] = pq.tasks[parent_index], pq.tasks[smaller_index]
		parent_index = smaller_index
		left_child_index = 2*parent_index + 1
		right_child_index = 2*parent_index + 2
	}
	pq.mu.Unlock()
	return rootTask, true
}

// IsEmpty checks if the priority queue is empty
func (pq *PriorityQueue) IsEmpty() bool {
	IsEmpty := true
	pq.mu.Lock()
	IsEmpty = len(pq.tasks) == 0
	pq.mu.Unlock()
	return IsEmpty
}

// Size returns the number of tasks in the queue
func (pq *PriorityQueue) Size() int {
	// TODO: Return queue size (with mutex lock)
	//
	// Steps:
	// 1. Lock the mutex
	// 2. Get len(pq.tasks)
	// 3. Unlock the mutex
	// 4. Return the size
	size := len(pq.tasks)
	pq.mu.Lock()
	size = len(pq.tasks)
	pq.mu.Unlock()
	return size
}

// ============================================
// 🧪 Priority Queue Test Suite
// ============================================
// This file tests the PriorityQueue implementation
// to verify it works correctly before using it in the scheduler.

func TestPQ() {
	fmt.Println("🧪 Testing Priority Queue Implementation")
	fmt.Println("==========================================")
	fmt.Println()

	// Track test results
	passed := 0
	failed := 0

	// Test 1: Create new priority queue
	fmt.Println("Test 1: Creating new priority queue...")
	pq := NewPriorityQueue()
	if pq == nil {
		fmt.Println("❌ FAILED: NewPriorityQueue returned nil")
		failed++
	} else {
		fmt.Println("✅ PASSED: Priority queue created")
		passed++
	}
	fmt.Println()

	// Test 2: IsEmpty on new queue
	fmt.Println("Test 2: IsEmpty() on new queue...")
	if pq.IsEmpty() {
		fmt.Println("✅ PASSED: New queue is empty")
		passed++
	} else {
		fmt.Println("❌ FAILED: New queue should be empty")
		failed++
	}
	fmt.Println()

	// Test 3: Size on new queue
	fmt.Println("Test 3: Size() on new queue...")
	size := pq.Size()
	if size == 0 {
		fmt.Println("✅ PASSED: New queue size is 0")
		passed++
	} else {
		fmt.Printf("❌ FAILED: Expected size 0, got %d\n", size)
		failed++
	}
	fmt.Println()

	// Test 4: Pop from empty queue
	fmt.Println("Test 4: Pop() from empty queue...")
	task, ok := pq.Pop()
	if !ok {
		fmt.Println("✅ PASSED: Pop from empty queue returns false")
		passed++
	} else {
		fmt.Printf("❌ FAILED: Pop from empty queue should return false, got task: %+v\n", task)
		failed++
	}
	fmt.Println()

	// Test 5: Push single task
	fmt.Println("Test 5: Push() single task...")
	task1 := Task{ID: 1, Priority: 3, Name: "Task 1", Duration: 100 * time.Millisecond}
	pq.Push(task1)
	size = pq.Size()
	if size == 1 {
		fmt.Println("✅ PASSED: Size is 1 after pushing one task")
		passed++
	} else {
		fmt.Printf("❌ FAILED: Expected size 1, got %d\n", size)
		failed++
	}
	if !pq.IsEmpty() {
		fmt.Println("✅ PASSED: Queue is not empty after push")
		passed++
	} else {
		fmt.Println("❌ FAILED: Queue should not be empty after push")
		failed++
	}
	fmt.Println()

	// Test 6: Pop single task
	fmt.Println("Test 6: Pop() single task...")
	popped, ok := pq.Pop()
	if ok && popped.ID == task1.ID && popped.Priority == task1.Priority {
		fmt.Println("✅ PASSED: Popped task matches pushed task")
		passed++
	} else {
		fmt.Printf("❌ FAILED: Expected task %+v, got %+v\n", task1, popped)
		failed++
	}
	if pq.IsEmpty() {
		fmt.Println("✅ PASSED: Queue is empty after popping last task")
		passed++
	} else {
		fmt.Println("❌ FAILED: Queue should be empty after popping last task")
		failed++
	}
	fmt.Println()

	// Test 7: Push multiple tasks and verify priority order
	fmt.Println("Test 7: Push multiple tasks and verify priority order...")
	tasks := []Task{
		{ID: 1, Priority: 5, Name: "Low Priority", Duration: 100 * time.Millisecond},
		{ID: 2, Priority: 1, Name: "High Priority", Duration: 100 * time.Millisecond},
		{ID: 3, Priority: 3, Name: "Medium Priority", Duration: 100 * time.Millisecond},
		{ID: 4, Priority: 2, Name: "Very High Priority", Duration: 100 * time.Millisecond},
		{ID: 5, Priority: 4, Name: "Low-Medium Priority", Duration: 100 * time.Millisecond},
		{ID: 6, Priority: 1, Name: "Another High Priority", Duration: 100 * time.Millisecond},
	}

	// Push all tasks
	for _, task := range tasks {
		pq.Push(task)
	}

	// Expected order when popping (lowest priority number = highest priority)
	expectedOrder := []int{1, 1, 2, 3, 4, 5} // Priority values

	fmt.Printf("   Pushed %d tasks\n", len(tasks))
	fmt.Println("   Expected pop order (by priority): 1, 1, 2, 3, 4, 5")

	// Pop all tasks and verify order
	allCorrect := true
	for i, expectedPriority := range expectedOrder {
		popped, ok := pq.Pop()
		if !ok {
			fmt.Printf("❌ FAILED: Pop returned false at index %d\n", i)
			allCorrect = false
			break
		}
		if popped.Priority != expectedPriority {
			fmt.Printf("❌ FAILED: At position %d, expected priority %d, got priority %d (Task ID: %d)\n",
				i, expectedPriority, popped.Priority, popped.ID)
			allCorrect = false
		} else {
			fmt.Printf("   ✓ Popped: Task %d (Priority: %d) - %s\n", popped.ID, popped.Priority, popped.Name)
		}
	}

	if allCorrect {
		fmt.Println("✅ PASSED: All tasks popped in correct priority order")
		passed++
	} else {
		fmt.Println("❌ FAILED: Tasks not in correct priority order")
		failed++
	}

	if pq.IsEmpty() {
		fmt.Println("✅ PASSED: Queue is empty after popping all tasks")
		passed++
	} else {
		fmt.Printf("❌ FAILED: Queue should be empty, but size is %d\n", pq.Size())
		failed++
	}
	fmt.Println()

	// Test 8: Push in different order, verify priority is maintained
	fmt.Println("Test 8: Push tasks in different order, verify priority order...")
	pq2 := NewPriorityQueue()
	// Push in reverse order
	for i := len(tasks) - 1; i >= 0; i-- {
		pq2.Push(tasks[i])
	}

	allCorrect = true
	for i, expectedPriority := range expectedOrder {
		popped, ok := pq2.Pop()
		if !ok || popped.Priority != expectedPriority {
			fmt.Printf("❌ FAILED: At position %d, expected priority %d, got priority %d\n",
				i, expectedPriority, popped.Priority)
			allCorrect = false
		}
	}

	if allCorrect {
		fmt.Println("✅ PASSED: Priority order maintained regardless of push order")
		passed++
	} else {
		fmt.Println("❌ FAILED: Priority order not maintained")
		failed++
	}
	fmt.Println()

	// Test 9: Concurrent access (basic test)
	fmt.Println("Test 9: Basic concurrent access test...")
	pq3 := NewPriorityQueue()
	// Push some tasks
	for _, task := range tasks[:3] {
		pq3.Push(task)
	}
	size = pq3.Size()
	if size == 3 {
		fmt.Println("✅ PASSED: Concurrent-safe operations work")
		passed++
	} else {
		fmt.Printf("❌ FAILED: Expected size 3, got %d\n", size)
		failed++
	}
	fmt.Println()

	// Summary
	fmt.Println("==========================================")
	fmt.Printf("Test Summary: %d passed, %d failed\n", passed, failed)
	if failed == 0 {
		fmt.Println("🎉 All tests passed! Priority queue is working correctly.")
	} else {
		fmt.Println("⚠️  Some tests failed. Review the implementation.")
	}
}
