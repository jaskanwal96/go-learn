package main

import (
	"database/sql"
	"fmt"
)

// ============================================
// 🎯 CHALLENGE 4.1: Database Operations
// ============================================
// This challenge teaches:
// - Database connections
// - SQL queries (SELECT, INSERT, UPDATE, DELETE)
// - Prepared statements
// - Transactions
// - Error handling
//
// TASK: Build a database-backed user management system
//
// Requirements:
// 1. Connect to SQLite database
// 2. Create users table
// 3. Implement CreateUser(user User) error
// 4. Implement GetUser(id int) (*User, error)
// 5. Implement GetAllUsers() ([]User, error)
// 6. Implement UpdateUser(id int, user User) error
// 7. Implement DeleteUser(id int) error
// 8. Use transactions for batch operations
//
// Concepts you'll practice:
// - database/sql package
// - sql.Open()
// - db.Exec() for CREATE, INSERT, UPDATE, DELETE
// - db.Query() and db.QueryRow() for SELECT
// - Prepared statements (db.Prepare())
// - Transactions (db.Begin(), tx.Commit(), tx.Rollback())
//
// Setup:
//   go get github.com/mattn/go-sqlite3

type User struct {
	ID    int
	Name  string
	Email string
}

type UserDB struct {
	db *sql.DB
}

// NewUserDB creates a new database connection
func NewUserDB(dbPath string) (*UserDB, error) {
	// TODO: Open database connection
	// db, err := sql.Open("sqlite3", dbPath)
	// return &UserDB{db: db}, err
	return nil, nil
}

// InitSchema creates the users table
func (udb *UserDB) InitSchema() error {
	// TODO: Create table if not exists
	// CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, email TEXT)
	return nil
}

// CreateUser inserts a new user
func (udb *UserDB) CreateUser(user User) error {
	// TODO: Insert user
	// Use prepared statement or db.Exec()
	return nil
}

// GetUser retrieves a user by ID
func (udb *UserDB) GetUser(id int) (*User, error) {
	// TODO: Query single user
	// Use db.QueryRow()
	return nil, nil
}

// GetAllUsers retrieves all users
func (udb *UserDB) GetAllUsers() ([]User, error) {
	// TODO: Query all users
	// Use db.Query() and iterate rows
	return nil, nil
}

// UpdateUser updates a user
func (udb *UserDB) UpdateUser(id int, user User) error {
	// TODO: Update user
	return nil
}

// DeleteUser deletes a user
func (udb *UserDB) DeleteUser(id int) error {
	// TODO: Delete user
	return nil
}

// BatchCreateUsers creates multiple users in a transaction
func (udb *UserDB) BatchCreateUsers(users []User) error {
	// TODO: Bonus - use transaction
	// tx, err := udb.db.Begin()
	// ... insert users ...
	// tx.Commit() or tx.Rollback()
	return nil
}

// Close closes the database connection
func (udb *UserDB) Close() error {
	return udb.db.Close()
}

func main() {
	fmt.Println("🎮 Challenge 4.1: Database Operations")
	fmt.Println("====================================\n")

	// TODO: Initialize database
	// userDB, err := NewUserDB("users.db")
	// if err != nil { log.Fatal(err) }
	// defer userDB.Close()
	// userDB.InitSchema()

	// TODO: Test operations
	// userDB.CreateUser(User{Name: "Alice", Email: "alice@example.com"})
	// user, _ := userDB.GetUser(1)
	// fmt.Printf("User: %+v\n", user)

	fmt.Println("✅ Challenge complete!")
}
