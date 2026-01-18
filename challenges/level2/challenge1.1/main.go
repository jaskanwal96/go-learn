package main

// ============================================
// 🎯 CHALLENGE 2.1: Concurrent Task Scheduler with Priority Queue
// ============================================
// This challenge teaches:
// - Goroutines (concurrent task execution)
// - Channels (task queue communication)
// - WaitGroups (synchronization)
// - Priority Queue data structure (heap-based)
// - Scheduling algorithms (priority-based)
//
// TASK: Build a concurrent task scheduler that processes tasks based on priority
//
// File Structure:
// - task.go: Defines the Task struct
// - priority_queue.go: Implements the PriorityQueue (min-heap) data structure
// - scheduler.go: Implements the Scheduler with goroutines, channels, and WaitGroups
// - main.go: This file - sets up and runs the challenge
//
// Requirements:
// 1. Implement a PriorityQueue data structure (min-heap) in priority_queue.go
// 2. Create a Scheduler that uses goroutines to process tasks concurrently in scheduler.go
// 3. Use channels to send tasks to workers
// 4. Use WaitGroup to wait for all tasks to complete
// 5. Process tasks in priority order (lower number = higher priority)
//
// Concepts you'll practice:
// - go keyword for goroutines
// - Channel communication (buffered channels)
// - sync.WaitGroup for synchronization
// - Heap data structure (priority queue)
// - Concurrent task processing
//
// HINT: Review challenge1 to understand:
//   - How goroutines work with channels
//   - How WaitGroups synchronize goroutines
//   - How to send/receive from channels

func main() {
	// fmt.Println("🎮 Challenge 2.1: Concurrent Task Scheduler with Priority Queue")
	// fmt.Println("===============================================================")
	// fmt.Println()

	// // Create scheduler with 3 workers
	// scheduler := NewScheduler(3)

	// // Add tasks with different priorities
	// tasks := []Task{
	// 	{ID: 1, Priority: 5, Name: "Low Priority Task", Duration: 200 * time.Millisecond},
	// 	{ID: 2, Priority: 1, Name: "High Priority Task", Duration: 150 * time.Millisecond},
	// 	{ID: 3, Priority: 3, Name: "Medium Priority Task", Duration: 100 * time.Millisecond},
	// 	{ID: 4, Priority: 2, Name: "Very High Priority Task", Duration: 120 * time.Millisecond},
	// 	{ID: 5, Priority: 4, Name: "Low-Medium Priority Task", Duration: 180 * time.Millisecond},
	// 	{ID: 6, Priority: 1, Name: "Another High Priority Task", Duration: 130 * time.Millisecond},
	// }

	// fmt.Println("Adding tasks to scheduler...")
	// for _, task := range tasks {
	// 	scheduler.AddTask(task)
	// 	fmt.Printf("  Added: %s (Priority: %d)\n", task.Name, task.Priority)
	// }

	// fmt.Println("\nStarting scheduler...")
	// scheduler.Start()

	// fmt.Println("\nWaiting for all tasks to complete...")
	// scheduler.Wait()

	// fmt.Printf("\n✅ All tasks completed! Total: %d\n", scheduler.GetCompleted())
	// fmt.Println("\nExpected behavior:")
	// fmt.Println("- Tasks with Priority 1 should complete first")
	// fmt.Println("- Tasks with Priority 2 should complete next")
	// fmt.Println("- And so on...")
	// fmt.Println("- Multiple tasks can run concurrently (3 workers)")
	TestPQ()
}
