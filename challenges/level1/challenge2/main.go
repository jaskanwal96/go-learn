package main

import "fmt"

// ============================================
// 🎯 CHALLENGE 1.2: Hash Table Implementation
// ============================================
// This challenge combines:
// - Structs and methods
// - Slices and arrays
// - Hashing functions
// - Collision handling
//
// TASK: Implement a HashTable with basic operations
//
// Requirements:
// 1. Create a HashTable struct with buckets (slice of slices)
// 2. Implement Set(key, value string) - stores key-value pair
// 3. Implement Get(key string) (string, bool) - retrieves value
// 4. Implement Delete(key string) - removes key-value pair
// 5. Implement hash function: hash(key string) int
// 6. Handle collisions using chaining (slice in each bucket)
//
// Concepts you'll practice:
// - Hash functions
// - Collision resolution
// - Slices of slices
// - Error handling

type KeyValue struct {
	Key   string
	Value string
}

type HashTable struct {
	// TODO: Add fields
	// Hint: buckets []([]KeyValue) - slice of slices
	buckets [][]KeyValue
	size    int
}

// NewHashTable creates a new hash table with given size
func NewHashTable(size int) *HashTable {
	// TODO: Initialize buckets
	return &HashTable{
		buckets: make([][]KeyValue, size),
		size:    size,
	}
}

// hash calculates hash value for a key
func (ht *HashTable) hash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % ht.size
}

// Set stores a key-value pair
func (ht *HashTable) Set(key, value string) {
	// TODO: Implement Set
	// 1. Calculate hash
	// 2. Find bucket
	// 3. Check if key exists, update if yes
	// 4. Otherwise append new KeyValue
	bucket := ht.hash(key)
	if len(ht.buckets[bucket]) == 0 {
		ht.buckets[bucket] = []KeyValue{{Key: key, Value: value}}
	} else {
		for index, existingKeyValues := range ht.buckets[bucket] {
			if key == existingKeyValues.Key {
				ht.buckets[bucket][index].Value = value
				return
			}
		}
		ht.buckets[bucket] = append(ht.buckets[bucket], KeyValue{Key: key, Value: value})
	}
}

// Get retrieves a value by key
func (ht *HashTable) Get(key string) (string, bool) {
	// TODO: Implement Get
	// 1. Calculate hash
	// 2. Search in bucket
	// 3. Return value and true if found
	bucket := ht.hash(key)
	if len(ht.buckets[bucket]) == 0 {
		return "", false
	}
	for index, kv := range ht.buckets[bucket] {
		if key == kv.Key {
			return ht.buckets[bucket][index].Value, true
		}
	}
	return "", false
}

// Delete removes a key-value pair
func (ht *HashTable) Delete(key string) bool {
	// TODO: Implement Delete
	// 1. Calculate hash
	// 2. Find and remove from bucket
	bucketIndex := ht.hash(key)
	bucket := ht.buckets[bucketIndex]

	// Find the index of the key
	for i, kv := range bucket {
		if kv.Key == key {
			// Remove element at index i using slicing
			lastIndex := len(bucket) - 1
			bucket[i] = bucket[lastIndex]
			// Truncate slice (remove last element)
			ht.buckets[bucketIndex] = bucket[:lastIndex]
			return true
		}
	}
	return false
}

// Keys returns all keys in the hash table
func (ht *HashTable) Keys() []string {
	// TODO: Bonus - return all keys
	return nil
}

func main() {
	fmt.Println("🎮 Challenge 1.2: Hash Table Implementation")
	fmt.Println("==========================================")

	ht := NewHashTable(10)

	// Test Set and Get
	fmt.Println("Test 1: Set and Get")
	ht.Set("name", "Alice")
	ht.Set("age", "25")
	ht.Set("city", "New York")

	if value, ok := ht.Get("name"); ok && value == "Alice" {
		fmt.Printf("  Get('name'): %s ✅\n", value)
	} else {
		fmt.Printf("  Get('name'): FAILED ❌ (got: %s, ok: %v)\n", value, ok)
	}

	// Test Update existing key
	fmt.Println("\nTest 2: Update existing key")
	ht.Set("name", "Sunny") // Update "name" from "Alice" to "Sunny"
	if value, ok := ht.Get("name"); ok && value == "Sunny" {
		fmt.Printf("  Get('name') after update: %s ✅\n", value)
	} else {
		fmt.Printf("  Update FAILED ❌ (expected: 'Sunny', got: %s)\n", value)
	}

	// Test Update city
	fmt.Println("\nTest 3: Update city")
	ht.Set("city", "San Francisco") // Update existing
	if value, ok := ht.Get("city"); ok && value == "San Francisco" {
		fmt.Printf("  Get('city'): %s ✅\n", value)
	} else {
		fmt.Printf("  Update FAILED ❌ (expected: 'San Francisco', got: %s)\n", value)
	}

	// Test Delete
	fmt.Println("\nTest 4: Delete")
	if ht.Delete("age") {
		fmt.Println("  Delete('age'): SUCCESS ✅")
		if _, ok := ht.Get("age"); !ok {
			fmt.Println("  Get('age') after delete: Not found ✅")
		} else {
			fmt.Println("  Get('age') after delete: Still exists ❌")
		}
	} else {
		fmt.Println("  Delete('age'): FAILED ❌ (Delete not implemented)")
	}

	fmt.Println("\n✅ Challenge complete!")
}
