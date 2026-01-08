package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// ============================================
// 🎯 CHALLENGE 3.2: REST API Server
// ============================================
// This challenge teaches:
// - Creating HTTP servers
// - REST endpoints
// - JSON encoding
// - Middleware
//
// TASK: Build a simple REST API server
//
// Requirements:
// 1. Create a UserStore (in-memory map)
// 2. Implement GET /users - list all users
// 3. Implement GET /users/{id} - get user by ID
// 4. Implement POST /users - create user
// 5. Implement DELETE /users/{id} - delete user
// 6. Add logging middleware
//
// Concepts you'll practice:
// - http.HandleFunc()
// - http.ListenAndServe()
// - JSON encoding (json.Marshal)
// - URL path parsing
// - HTTP methods (GET, POST, DELETE)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserStore struct {
	users map[int]*User
	nextID int
	mu    sync.RWMutex
}

func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

// TODO: Implement handler functions
// - handleGetUsers(w http.ResponseWriter, r *http.Request)
// - handleGetUser(w http.ResponseWriter, r *http.Request)
// - handleCreateUser(w http.ResponseWriter, r *http.Request)
// - handleDeleteUser(w http.ResponseWriter, r *http.Request)
// - loggingMiddleware(next http.HandlerFunc) http.HandlerFunc

func main() {
	fmt.Println("🎮 Challenge 3.2: REST API Server")
	fmt.Println("================================\n")

	_ = NewUserStore() // TODO: Use store to register routes

	// TODO: Register routes
	// http.HandleFunc("/users", ...)
	// http.HandleFunc("/users/", ...)

	// TODO: Start server
	// http.ListenAndServe(":8080", nil)

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Test with:")
	fmt.Println("  curl http://localhost:8080/users")
	fmt.Println("  curl -X POST http://localhost:8080/users -d '{\"name\":\"Alice\",\"email\":\"alice@example.com\"}'")
}

