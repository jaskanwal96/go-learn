# Challenge 2.1: Concurrent Task Scheduler with Priority Queue

## Learning Objectives

After completing this challenge, you'll understand:

1. **Goroutines** - How to run functions concurrently
2. **Channels** - How goroutines communicate with each other
3. **WaitGroups** - How to synchronize multiple goroutines
4. **Priority Queue (Heap)** - A fundamental data structure for scheduling
5. **Concurrent Programming Patterns** - Worker pool and dispatcher patterns

## File Structure

This challenge is split into multiple files for easier understanding:

- **`task.go`** - Defines the Task struct (what we're processing)
- **`priority_queue.go`** - Implements the PriorityQueue data structure (min-heap)
- **`scheduler.go`** - Implements the Scheduler with goroutines, channels, and WaitGroups
- **`main.go`** - Sets up and runs the challenge

## Key Concepts

### What Challenge 1 Taught Us

Challenge 1 demonstrated:
- **Goroutines**: Using `go` keyword to run functions concurrently
- **Channels**: `chan<-` for sending, `<-chan` for receiving
- **WaitGroups**: `wg.Add()`, `wg.Done()`, `wg.Wait()` for synchronization
- **Buffered Channels**: Channels with capacity to prevent blocking

### Priority Queue (Min-Heap)

A priority queue is a data structure where:
- Elements are stored based on priority
- Highest priority element is always at the root
- In a **min-heap**, lower numbers = higher priority
- Operations: `Push()` adds element, `Pop()` removes highest priority element

#### Heap Properties:
- **Complete Binary Tree**: All levels filled except possibly the last
- **Heap Property**: Parent priority ≤ child priority (min-heap)
- **Parent Index**: `(child_index - 1) / 2`
- **Left Child**: `2 * index + 1`
- **Right Child**: `2 * index + 2`

#### Bubble Up (for Push):
```
1. Add element to end of array
2. Compare with parent
3. If element.priority < parent.priority, swap
4. Repeat until heap property satisfied
```

#### Bubble Down (for Pop):
```
1. Remove root element
2. Move last element to root
3. Compare with children
4. Swap with smaller priority child
5. Repeat until heap property satisfied
```

### Concurrent Task Processing

The scheduler uses:
- **Worker Goroutines**: Process tasks concurrently
- **Dispatcher Goroutine**: Sends tasks from priority queue to workers
- **Channel**: Communication between dispatcher and workers
- **WaitGroup**: Ensures all tasks complete before program ends

## Step-by-Step Implementation Guide

### Step 1: Implement Priority Queue (`priority_queue.go`)

1. **NewPriorityQueue()**: Return `&PriorityQueue{tasks: []Task{}}`
2. **Push()**: 
   - Lock mutex
   - Append task to slice
   - Bubble up: while `index > 0` and `task.Priority < parent.Priority`, swap
   - Unlock mutex
3. **Pop()**:
   - Lock mutex
   - If empty, return `Task{}, false`
   - Save root task
   - Move last element to root
   - Bubble down: compare with children, swap with smaller
   - Unlock mutex
   - Return saved task and `true`
4. **IsEmpty()** and **Size()**: Simple checks with mutex

### Step 2: Implement Scheduler (`scheduler.go`)

1. **NewScheduler()**: 
   - Create PriorityQueue
   - Create buffered channel: `make(chan Task, workers * 2)`
   - Initialize struct fields
2. **AddTask()**: Call `pq.Push(task)`
3. **processTask()**: 
   - Print start message
   - `time.Sleep(task.Duration)`
   - Print completion message
   - Increment completed (with mutex)
4. **Start()**:
   - Start worker goroutines (loop receiving from channel)
   - Start dispatcher goroutine (pop from queue, send to channel, close when empty)
5. **Wait()**: Call `wg.Wait()`
6. **GetCompleted()**: Return completed count (with mutex)

## Implementation Order

Recommended order to implement:

1. **Start with `task.go`** - Already done! Just understand the Task struct.
2. **Then `priority_queue.go`**:
   - NewPriorityQueue()
   - IsEmpty() and Size() (easiest)
   - Push() (bubble up)
   - Pop() (bubble down)
3. **Then `scheduler.go`**:
   - NewScheduler()
   - AddTask()
   - processTask()
   - Wait() and GetCompleted()
   - Start() (most complex - workers and dispatcher)

## Hints

### Priority Queue Implementation:

```go
// Bubble up helper (you can add this to priority_queue.go)
func (pq *PriorityQueue) bubbleUp(index int) {
    for index > 0 {
        parent := (index - 1) / 2
        if pq.tasks[index].Priority >= pq.tasks[parent].Priority {
            break
        }
        pq.tasks[index], pq.tasks[parent] = pq.tasks[parent], pq.tasks[index]
        index = parent
    }
}

// Bubble down helper
func (pq *PriorityQueue) bubbleDown(index int) {
    for {
        smallest := index
        left := 2*index + 1
        right := 2*index + 2
        
        if left < len(pq.tasks) && pq.tasks[left].Priority < pq.tasks[smallest].Priority {
            smallest = left
        }
        if right < len(pq.tasks) && pq.tasks[right].Priority < pq.tasks[smallest].Priority {
            smallest = right
        }
        if smallest == index {
            break
        }
        pq.tasks[index], pq.tasks[smallest] = pq.tasks[smallest], pq.tasks[index]
        index = smallest
    }
}
```

### Worker Pattern (from Challenge 1):

```go
// Worker goroutine pattern
go func() {
    defer wg.Done()
    for task := range taskChan {
        processTask(task)
    }
}()
```

### Dispatcher Pattern:

```go
// Dispatcher goroutine
go func() {
    for !pq.IsEmpty() {
        task, ok := pq.Pop()
        if ok {
            taskChan <- task
        }
    }
    close(taskChan)
}()
```

## Testing Your Solution

Run the challenge:
```bash
go run challenges/level2/challenge2.1/*.go
```

Or if you're in the directory:
```bash
cd challenges/level2/challenge2.1
go run *.go
```

Expected output:
- Tasks should be processed in priority order (1, 1, 2, 3, 4, 5)
- Multiple tasks can run concurrently (3 workers)
- All tasks should complete successfully

## Common Mistakes

1. **Forgetting mutex locks**: Priority queue operations need mutex protection
2. **Wrong heap property**: Remember min-heap means smaller number = higher priority
3. **Not closing channel**: Dispatcher must close channel when done
4. **Race conditions**: Use mutex for shared state (completed counter)
5. **Deadlock**: Make sure WaitGroup is used correctly
6. **Index out of bounds**: Check array bounds before accessing children in heap

## What You're Learning

- **Data Structures**: Priority queue (heap) implementation
- **Algorithms**: Heap operations (bubble up/down)
- **Concurrency**: Goroutines, channels, WaitGroups
- **Synchronization**: Mutex for thread-safe operations
- **Design Patterns**: Worker pool, dispatcher pattern
- **Code Organization**: Splitting code into logical files

## Next Steps

After completing this challenge:
- Try implementing a max-heap (reverse the comparisons)
- Add task cancellation with context.Context
- Implement task retry logic
- Add metrics (average wait time, throughput)
- Experiment with different numbers of workers
