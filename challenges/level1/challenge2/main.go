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
	return nil
}

// hash calculates hash value for a key
func (ht *HashTable) hash(key string) int {
	// TODO: Implement hash function
	// Hint: Sum ASCII values and use modulo
	return 0
}

// Set stores a key-value pair
func (ht *HashTable) Set(key, value string) {
	// TODO: Implement Set
	// 1. Calculate hash
	// 2. Find bucket
	// 3. Check if key exists, update if yes
	// 4. Otherwise append new KeyValue
}

// Get retrieves a value by key
func (ht *HashTable) Get(key string) (string, bool) {
	// TODO: Implement Get
	// 1. Calculate hash
	// 2. Search in bucket
	// 3. Return value and true if found
	return "", false
}

// Delete removes a key-value pair
func (ht *HashTable) Delete(key string) bool {
	// TODO: Implement Delete
	// 1. Calculate hash
	// 2. Find and remove from bucket
	return false
}

// Keys returns all keys in the hash table
func (ht *HashTable) Keys() []string {
	// TODO: Bonus - return all keys
	return nil
}

func main() {
	fmt.Println("🎮 Challenge 1.2: Hash Table Implementation")
	fmt.Println("==========================================\n")

	ht := NewHashTable(10)

	// Test Set and Get
	ht.Set("name", "Alice")
	ht.Set("age", "25")
	ht.Set("city", "New York")

	if value, ok := ht.Get("name"); ok {
		fmt.Printf("Get('name'): %s ✅\n", value)
	} else {
		fmt.Println("Get('name'): FAILED ❌")
	}

	// Test collision (if hash function causes it)
	ht.Set("city", "San Francisco") // Update existing
	if value, ok := ht.Get("city"); ok {
		fmt.Printf("Get('city'): %s ✅\n", value)
	}

	// Test Delete
	if ht.Delete("age") {
		fmt.Println("Delete('age'): SUCCESS ✅")
	}
	if _, ok := ht.Get("age"); !ok {
		fmt.Println("Get('age') after delete: Not found ✅")
	}

	fmt.Println("\n✅ Challenge complete!")
}

