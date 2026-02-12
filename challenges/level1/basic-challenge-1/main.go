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

// AddEdge adds a directed edge from 'from' to 'to'
func (g *Graph) AddEdge(from, to string) {
	if _, exists := g.adj[from]; !exists {
		g.adj[from] = make([]string, 0)
	}
	if _, exists := g.adj[to]; !exists {
		g.adj[to] = make([]string, 0)
	}
	g.adj[from] = append(g.adj[from], to)

}

// Returns the order of visited nodes
func (g *Graph) BFS(start string) []string {
	// TODO: Implement BFS
	visited := make(map[string]bool, len(g.adj))
	queue := []string{start}
	result := make([]string, 0)

	visited[start] = true
	for len(queue) != 0 {
		var node = queue[0]
		queue = queue[1:]
		result = append(result, node)
		for _, neighbor := range g.adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return result
}

// DFS performs Depth-First Search starting from 'start'
// Returns the order of visited nodes
func (g *Graph) DFS(start string) []string {
	// TODO: Implement DFS
	visited := make(map[string]bool, len(g.adj))
	result := []string{}
	g.dfsHelper(start, visited, &result)
	return result
}

// Helper function for DFS recursion
func (g *Graph) dfsHelper(node string, visited map[string]bool, result *[]string) {
	// TODO: Implement recursive DFS helper
	visited[node] = true
	*result = append(*result, node)
	var connectedNodes = g.adj[node]
	for _, value := range connectedNodes {
		if !visited[value] {
			g.dfsHelper(value, visited, result)
		}
	}
}

func main() {
	fmt.Println("🎮 Challenge 1.1: Graph Traversal")
	fmt.Println("===================================")

	// Create a graph
	graph := NewGraph()

	// Build this graph:
	//     A
	//    / \
	//   B   C
	//  / \   \
	// D   E   F

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
