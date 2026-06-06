# Challenge 2.3: Concurrent URL Fetcher with Rate Limiting

## Learning Objectives

After completing this challenge, you'll master:

1. **Goroutines** - Concurrent execution of multiple tasks
2. **Channels** - Communication between goroutines (buffered and unbuffered)
3. **WaitGroups** - Synchronizing multiple goroutines
4. **Mutex** - Protecting shared data from race conditions
5. **Real-world patterns** - Worker pools, rate limiting, result aggregation

## What You're Building

A **concurrent URL fetcher** that:
- Fetches multiple URLs concurrently using worker goroutines
- Limits concurrent requests (rate limiting)
- Aggregates results safely using mutex
- Reports success/failure statistics
- Handles errors gracefully

## File Structure

- **`url.go`** - Defines the URL and Result structs
- **`fetcher.go`** - Implements the concurrent fetcher with workers
- **`stats.go`** - Thread-safe statistics collector using mutex
- **`main.go`** - Sets up and runs the challenge

## Key Concepts

### Worker Pool Pattern

```
URL Queue → [urlChan] → Worker 1 (goroutine)
                      → Worker 2 (goroutine)
                      → Worker 3 (goroutine)
                      → Worker N (goroutines)
                      
Each worker:
1. Receives URL from channel
2. Fetches the URL (HTTP request)
3. Sends result to results channel
4. Updates shared statistics (with mutex)
```

### Rate Limiting

By controlling the number of workers, we limit concurrent requests:
- 5 workers = max 5 concurrent HTTP requests
- Prevents overwhelming servers
- Real-world production pattern

### Result Aggregation

A separate goroutine collects results from all workers:
```
Workers → [resultsChan] → Collector Goroutine → Stats
```

## Step-by-Step Implementation Guide

### Step 1: Implement Stats Collector (`stats.go`)

The stats collector tracks:
- Total URLs processed
- Successful fetches
- Failed fetches
- Total bytes downloaded

**Key challenge**: Multiple workers update stats concurrently!
**Solution**: Use `sync.Mutex` to protect shared data.

```go
type Stats struct {
    mu          sync.Mutex
    total       int
    successful  int
    failed      int
    totalBytes  int64
}

func (s *Stats) RecordSuccess(bytes int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.total++
    s.successful++
    s.totalBytes += int64(bytes)
}
```

### Step 2: Implement Fetcher (`fetcher.go`)

The fetcher manages:
1. **URL Channel**: Sends URLs to workers
2. **Results Channel**: Receives results from workers
3. **Worker Goroutines**: Fetch URLs concurrently
4. **Collector Goroutine**: Aggregates results
5. **WaitGroup**: Ensures all workers finish

**Architecture**:
```
AddURL() → urlChan → Workers → resultsChan → Collector → Stats
```

**Key methods**:
- `NewFetcher(workers int)` - Creates fetcher with N workers
- `AddURL(url string)` - Adds URL to fetch queue
- `Start()` - Launches workers and collector
- `Wait()` - Waits for all work to complete
- `GetStats()` - Returns final statistics

### Step 3: Worker Implementation

Each worker:
```go
go func(id int) {
    defer wg.Done()
    for url := range urlChan {
        // Fetch URL
        resp, err := http.Get(url)
        
        // Create result
        result := Result{URL: url, ...}
        
        // Send to results channel
        resultsChan <- result
    }
}(workerID)
```

### Step 4: Collector Implementation

The collector aggregates results:
```go
go func() {
    for result := range resultsChan {
        if result.Success {
            stats.RecordSuccess(result.Size)
        } else {
            stats.RecordFailure()
        }
    }
    collectorDone <- true
}()
```

## Implementation Order

1. **`url.go`** - Define structs (easiest)
2. **`stats.go`** - Implement thread-safe stats collector
3. **`fetcher.go`** - Implement the concurrent fetcher:
   - `NewFetcher()` - Initialize channels and stats
   - `AddURL()` - Send URLs to channel
   - `fetchURL()` - HTTP GET logic
   - `Start()` - Launch workers and collector
   - `Wait()` - Synchronization
   - `GetStats()` - Return statistics

## Hints

### HTTP Fetching

```go
import "net/http"

resp, err := http.Get(url)
if err != nil {
    // Handle error
    return
}
defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)
if err != nil {
    // Handle error
    return
}

// body contains the response
size := len(body)
```

### Channel Patterns

```go
// Buffered channel (non-blocking sends)
urlChan := make(chan string, 100)

// Unbuffered channel (synchronous)
resultsChan := make(chan Result)

// Sending
urlChan <- "https://example.com"

// Receiving
url := <-urlChan

// Range over channel (exits when closed)
for url := range urlChan {
    // Process url
}

// Close channel (sender's responsibility)
close(urlChan)
```

### WaitGroup Pattern

```go
var wg sync.WaitGroup

// Before launching goroutine
wg.Add(1)

// Inside goroutine
go func() {
    defer wg.Done()
    // Do work
}()

// Wait for all goroutines
wg.Wait()
```

### Mutex Pattern

```go
var mu sync.Mutex

// Protect critical section
mu.Lock()
sharedData++
mu.Unlock()

// Or use defer
mu.Lock()
defer mu.Unlock()
sharedData++
```

## Testing Your Solution

Run the challenge:
```bash
cd challenges/level2/challenge1.3
go run *.go
```

Expected output:
- Workers fetch URLs concurrently
- Progress messages from each worker
- Final statistics (total, successful, failed, bytes)
- All URLs processed

## Common Mistakes

1. **Forgetting to close channels**: Causes goroutines to hang forever
2. **Not using mutex**: Race conditions when updating stats
3. **Forgetting WaitGroup.Add()**: Program exits before workers finish
4. **Closing channel from receiver**: Only sender should close
5. **Deadlock**: Sending to unbuffered channel with no receiver
6. **Not using defer for mutex**: Forgetting to unlock causes deadlock

## What You're Learning

- **Concurrency patterns**: Worker pools, fan-out/fan-in
- **Synchronization**: WaitGroups, channels, mutexes
- **Real-world skills**: HTTP requests, error handling, rate limiting
- **Best practices**: Proper channel closing, defer usage
- **Performance**: Concurrent vs sequential execution

## Challenge Yourself

After completing the basic version:
- Add timeout for slow URLs (use `context.Context`)
- Implement retry logic for failed requests
- Add a progress bar showing completion percentage
- Save results to a file
- Parse HTML and extract links (build a mini crawler!)
- Add configurable timeout per request
- Implement exponential backoff for retries

## Comparison to Challenge 1.1

| Aspect | Challenge 1.1 (Scheduler) | Challenge 1.3 (Fetcher) |
|--------|--------------------------|-------------------------|
| **Data Structure** | Priority Queue (heap) | Simple queue (channel) |
| **Workers** | Process tasks by priority | Fetch URLs concurrently |
| **Channels** | 1 channel (tasks) | 2 channels (URLs + results) |
| **Mutex Usage** | Completed counter | Full stats struct |
| **Complexity** | Medium | Medium-High |
| **Real-world** | Job scheduling | Web scraping, API calls |

Both challenges teach the same core concepts but in different contexts!
