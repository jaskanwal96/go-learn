package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ============================================
// 🎯 CHALLENGE 3.1: HTTP Client & API Calls
// ============================================
// This challenge teaches:
// - Making HTTP requests
// - Parsing JSON responses
// - Error handling
// - Struct tags for JSON
//
// TASK: Build an API client that fetches data
//
// Requirements:
// 1. Create a struct to represent API response
// 2. Implement FetchUser(userID int) function
// 3. Make GET request to JSONPlaceholder API
// 4. Parse JSON response
// 5. Handle errors properly
//
// API: https://jsonplaceholder.typicode.com/users/{id}
//
// Concepts you'll practice:
// - http.Get() or http.Client
// - json.Unmarshal()
// - Struct tags (`json:"field"`)
// - Error handling

type User struct {
	// TODO: Add fields with JSON tags
	// Fields: ID, Name, Email, Phone, Website
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

// FetchUser fetches user data from API
func FetchUser(userID int) (*User, error) {
	// TODO: Implement this
	// 1. Build URL: fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", userID)
	// 2. Make HTTP GET request
	// 3. Check for errors
	// 4. Read response body
	// 5. Unmarshal JSON into User struct
	// 6. Return user and error

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", userID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// FetchMultipleUsers fetches multiple users concurrently
func FetchMultipleUsers(userIDs []int) ([]*User, error) {
	// TODO: Bonus - fetch multiple users using goroutines
	// Use channels to collect results
	return nil, nil
}

func main() {
	fmt.Println("🎮 Challenge 3.1: HTTP Client & API Calls")
	fmt.Println("========================================\n")

	// Fetch a single user
	user, err := FetchUser(1)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	fmt.Printf("✅ User fetched:\n")
	fmt.Printf("  ID: %d\n", user.ID)
	fmt.Printf("  Name: %s\n", user.Name)
	fmt.Printf("  Email: %s\n", user.Email)
	fmt.Printf("  Phone: %s\n", user.Phone)

	// Bonus: Fetch multiple users
	fmt.Println("\n bonus: Fetching multiple users...")
	userIDs := []int{1, 2, 3}
	users, err := FetchMultipleUsers(userIDs)
	if err == nil {
		fmt.Printf("✅ Fetched %d users\n", len(users))
	}

	fmt.Println("\n✅ Challenge complete!")
}
