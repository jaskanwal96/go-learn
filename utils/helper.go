package utils

import "fmt"

// Add adds two integers and returns the result
// Capitalized function name = exported (can be used outside package)
func Add(a, b int) int {
	return a + b
}

// Multiply multiplies two integers
func Multiply(a, b int) int {
	return a * b
}

// Greet prints a greeting message
func Greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// subtract is not exported (lowercase) - can only be used within this package
func subtract(a, b int) int {
	return a - b
}
