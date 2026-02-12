#!/bin/bash

# Script to rename challenge folders with descriptive names
# This will make it easier to identify challenges at a glance

cd /Users/sunnysingh/Learning/go-learn/challenges

echo "Starting folder renaming process..."

# Level 1 Challenges
echo "Renaming Level 1 challenges..."

# Based on README content and file analysis:
# challenge1 - appears to be empty/basic
# challenge1.1 - leetcode 332 (Reconstruct Itinerary)
# challenge2 - appears to be empty/basic
# challenge3 - Interfaces & Polymorphism
# challenge3.1 - appears to be a variant
# challenge3.4 - appears to be a variant
# scanner - Input scanning basics

if [ -d "level1/challenge1.1" ]; then
    mv level1/challenge1.1 level1/leetcode-332-reconstruct-itinerary
    echo "✓ Renamed challenge1.1 → leetcode-332-reconstruct-itinerary"
fi

if [ -d "level1/challenge3" ]; then
    mv level1/challenge3 level1/interfaces-polymorphism
    echo "✓ Renamed challenge3 → interfaces-polymorphism"
fi

if [ -d "level1/challenge3.1" ]; then
    mv level1/challenge3.1 level1/interfaces-variant-1
    echo "✓ Renamed challenge3.1 → interfaces-variant-1"
fi

if [ -d "level1/challenge3.4" ]; then
    mv level1/challenge3.4 level1/interfaces-variant-4
    echo "✓ Renamed challenge3.4 → interfaces-variant-4"
fi

if [ -d "level1/scanner" ]; then
    mv level1/scanner level1/input-scanner-basics
    echo "✓ Renamed scanner → input-scanner-basics"
fi

if [ -d "level1/challenge1" ]; then
    mv level1/challenge1 level1/basic-challenge-1
    echo "✓ Renamed challenge1 → basic-challenge-1"
fi

if [ -d "level1/challenge2" ]; then
    mv level1/challenge2 level1/basic-challenge-2
    echo "✓ Renamed challenge2 → basic-challenge-2"
fi

# Level 2 Challenges
echo "Renaming Level 2 challenges..."

# Based on README content:
# challenge0 - Goroutines basics (say hello/world)
# challenge1 - appears to be empty/basic
# challenge1.1 - Concurrent Task Scheduler with Priority Queue
# challenge1.2 - RPC using GO
# challenge1.3 - Concurrent URL Fetcher with Rate Limiting
# challenge2 - appears to be empty/basic

if [ -d "level2/challenge0" ]; then
    mv level2/challenge0 level2/goroutines-basics
    echo "✓ Renamed challenge0 → goroutines-basics"
fi

if [ -d "level2/challenge1.1" ]; then
    mv level2/challenge1.1 level2/concurrent-task-scheduler
    echo "✓ Renamed challenge1.1 → concurrent-task-scheduler"
fi

if [ -d "level2/challenge1.2" ]; then
    mv level2/challenge1.2 level2/rpc-implementation
    echo "✓ Renamed challenge1.2 → rpc-implementation"
fi

if [ -d "level2/challenge1.3" ]; then
    mv level2/challenge1.3 level2/concurrent-url-fetcher
    echo "✓ Renamed challenge1.3 → concurrent-url-fetcher"
fi

if [ -d "level2/challenge1" ]; then
    mv level2/challenge1 level2/concurrency-challenge-1
    echo "✓ Renamed challenge1 → concurrency-challenge-1"
fi

if [ -d "level2/challenge2" ]; then
    mv level2/challenge2 level2/concurrency-challenge-2
    echo "✓ Renamed challenge2 → concurrency-challenge-2"
fi

# Level 3 Challenges
echo "Renaming Level 3 challenges..."

if [ -d "level3/challenge1" ]; then
    mv level3/challenge1 level3/advanced-challenge-1
    echo "✓ Renamed challenge1 → advanced-challenge-1"
fi

if [ -d "level3/challenge2" ]; then
    mv level3/challenge2 level3/advanced-challenge-2
    echo "✓ Renamed challenge2 → advanced-challenge-2"
fi

# Level 4 Challenges
echo "Renaming Level 4 challenges..."

if [ -d "level4/challenge1" ]; then
    mv level4/challenge1 level4/expert-challenge-1
    echo "✓ Renamed challenge1 → expert-challenge-1"
fi

# Level 5 Challenges
echo "Renaming Level 5 challenges..."

if [ -d "level5/challenge1" ]; then
    mv level5/challenge1 level5/master-challenge-1
    echo "✓ Renamed challenge1 → master-challenge-1"
fi

echo ""
echo "✅ Folder renaming complete!"
echo ""
echo "Summary of changes:"
echo "==================="
echo "Level 1:"
echo "  - challenge1.1 → leetcode-332-reconstruct-itinerary"
echo "  - challenge3 → interfaces-polymorphism"
echo "  - challenge3.1 → interfaces-variant-1"
echo "  - challenge3.4 → interfaces-variant-4"
echo "  - scanner → input-scanner-basics"
echo "  - challenge1 → basic-challenge-1"
echo "  - challenge2 → basic-challenge-2"
echo ""
echo "Level 2:"
echo "  - challenge0 → goroutines-basics"
echo "  - challenge1.1 → concurrent-task-scheduler"
echo "  - challenge1.2 → rpc-implementation"
echo "  - challenge1.3 → concurrent-url-fetcher"
echo "  - challenge1 → concurrency-challenge-1"
echo "  - challenge2 → concurrency-challenge-2"
echo ""
echo "Level 3:"
echo "  - challenge1 → advanced-challenge-1"
echo "  - challenge2 → advanced-challenge-2"
echo ""
echo "Level 4:"
echo "  - challenge1 → expert-challenge-1"
echo ""
echo "Level 5:"
echo "  - challenge1 → master-challenge-1"
