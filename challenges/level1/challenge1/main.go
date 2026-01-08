package main

import "fmt"

// ============================================
// 🎯 CHALLENGE 1.1: Graph Traversal (BFS/DFS)
// ============================================
// This challenge combines multiple concepts:
// - Structs and methods
// - Slices and maps
// - Recursion
// - Graph algorithms
//
// TASK: Implement a Graph with BFS and DFS traversal
//
// Requirements:
// 1. Create a Graph struct that stores nodes and edges
// 2. Implement AddNode(node string) - adds a node
// 3. Implement AddEdge(from, to string) - adds an edge
// 4. Implement BFS(start string) []string - Breadth-First Search
// 5. Implement DFS(start string) []string - Depth-First Search
// 6. Use a map[string][]string to store adjacency list
// 7. Use a map[string]bool to track visited nodes (hashing!)
//
// Concepts you'll practice:
// - Maps for adjacency list and visited tracking
// - Slices for queues and results
// - Recursion for DFS
// - Iteration for BFS
// - Error handling for invalid nodes

type Graph struct {
	// TODO: Add fields here
	// Hint: You need an adjacency list (map[string][]string)
	adj map[string][]string
}

// NewGraph creates a new empty graph
func NewGraph() *Graph {
	return &Graph{
		adj: make(map[string][]string),
	}
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(node string) {
	g.adj[node] = make([]string, 0)
}

// AddEdge adds a directed edge from 'from' to 'to'
func (g *Graph) AddEdge(from, to string) {
	if _, exists := g.adj[from]; !exists {
		g.adj[from] = make([]string, 0)
	} else {
		g.adj[from] = append(g.adj[from], to)
	}
}

// Returns the order of visited nodes
func (g *Graph) BFS(start string) []string {
	// TODO: Implement BFS
	// Hint: Use a queue (slice), visited map, and iteration
	return nil
}

// DFS performs Depth-First Search starting from 'start'
// Returns the order of visited nodes
func (g *Graph) DFS(start string) []string {
	// TODO: Implement DFS
	// Hint: Use recursion and a visited map
	return nil
}

// Helper function for DFS recursion
func (g *Graph) dfsHelper(node string, visited map[string]bool, result *[]string) {
	// TODO: Implement recursive DFS helper
}

func main() {
	fmt.Println("🎮 Challenge 1.1: Graph Traversal")
	fmt.Println("===================================\n")

	// Create a graph
	graph := NewGraph()

	// Build this graph:
	//     A
	//    / \
	//   B   C
	//  / \   \
	// D   E   F
	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddNode("E")
	graph.AddNode("F")

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("B", "E")
	graph.AddEdge("C", "F")

	// Test BFS
	fmt.Println("BFS starting from A:")
	bfsResult := graph.BFS("A")
	fmt.Println(bfsResult)
	// Expected: [A B C D E F] or [A C B F D E] (order may vary slightly)

	// Test DFS
	fmt.Println("\nDFS starting from A:")
	dfsResult := graph.DFS("A")
	fmt.Println(dfsResult)
	// Expected: [A B D E C F] or similar (depends on implementation)

	fmt.Println("✅ Challenge complete! Check your output matches the expected results.")
}
