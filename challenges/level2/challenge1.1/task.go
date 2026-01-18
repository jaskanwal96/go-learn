package main

import "time"

// ============================================
// Task Definition
// ============================================
// This file defines the Task struct that represents
// a task to be processed by the scheduler.
//
// A Task has:
// - ID: Unique identifier
// - Priority: Lower number = higher priority (1 is highest)
// - Name: Human-readable task name
// - Duration: How long the task takes to process

// Task represents a task to be processed
type Task struct {
	ID       int
	Priority int           // Lower number = higher priority
	Name     string
	Duration time.Duration
}
