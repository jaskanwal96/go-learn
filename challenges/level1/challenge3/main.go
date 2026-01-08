package main

import (
	"fmt"
)

// ============================================
// 🎯 CHALLENGE 1.3: Interfaces & Polymorphism
// ============================================
// This challenge combines multiple concepts:
// - Interfaces (Go's way of achieving polymorphism)
// - Method sets (value vs pointer receivers)
// - Type assertions
// - Interface composition
//
// TASK: Build a shape calculator using interfaces
//
// Requirements:
// 1. Create a Shape interface with Area() and Perimeter() methods
// 2. Implement Circle, Rectangle, and Triangle structs
// 3. Make all shapes implement the Shape interface
// 4. Create a function that accepts any Shape (polymorphism)
// 5. Use type assertions to handle specific types
// 6. Create a ShapeCollection that can hold different shapes
//
// Concepts you'll practice:
// - Defining interfaces
// - Implementing interfaces (implicit in Go!)
// - Polymorphism (one interface, multiple types)
// - Type assertions (value, ok := shape.(Circle))
// - Interface composition
// - Method sets (when to use value vs pointer receivers)

// Shape interface defines what a shape must do
// Any type with Area() and Perimeter() methods implements this interface
type Shape interface {
	// TODO: Define interface methods
	// Area() float64
	// Perimeter() float64
}

// Circle represents a circle
type Circle struct {
	// TODO: Add fields (radius float64)
}

// TODO: Implement Shape interface for Circle
// func (c Circle) Area() float64 { ... }
// func (c Circle) Perimeter() float64 { ... }

// Rectangle represents a rectangle
type Rectangle struct {
	// TODO: Add fields (width, height float64)
}

// TODO: Implement Shape interface for Rectangle
// func (r Rectangle) Area() float64 { ... }
// func (r Rectangle) Perimeter() float64 { ... }

// Triangle represents a triangle (assume right triangle for simplicity)
type Triangle struct {
	// TODO: Add fields (base, height float64)
}

// TODO: Implement Shape interface for Triangle
// func (t Triangle) Area() float64 { ... }
// func (t Triangle) Perimeter() float64 { ... }
// Hint: For perimeter, use Pythagorean theorem: sqrt(base^2 + height^2) for hypotenuse

// PrintShapeInfo prints information about any shape
// This demonstrates polymorphism - one function works with any Shape
func PrintShapeInfo(s Shape) {
	// TODO: Print area and perimeter
	// Use fmt.Printf to display the information
}

// GetShapeType returns the type name of a shape using type assertion
func GetShapeType(s Shape) string {
	// TODO: Use type assertion to determine the type
	// Example: if _, ok := s.(Circle); ok { return "Circle" }
	// Try: Circle, Rectangle, Triangle
	return "Unknown"
}

// ShapeCollection holds multiple shapes
type ShapeCollection struct {
	// TODO: Add field to store shapes ([]Shape)
}

// Add adds a shape to the collection
func (sc *ShapeCollection) Add(s Shape) {
	// TODO: Append shape to collection
}

// TotalArea calculates the sum of all shape areas
func (sc *ShapeCollection) TotalArea() float64 {
	// TODO: Iterate through shapes and sum their areas
	return 0.0
}

// FindLargest returns the shape with the largest area
func (sc *ShapeCollection) FindLargest() Shape {
	// TODO: Find and return the shape with maximum area
	return nil
}

// ============================================
// BONUS: Interface Composition
// ============================================

// Drawable interface for shapes that can be drawn
type Drawable interface {
	Draw() string
}

// Colored interface for shapes with color
type Colored interface {
	GetColor() string
	SetColor(string)
}

// ColoredShape combines Shape and Colored interfaces
type ColoredShape interface {
	Shape
	Colored
}

// ColoredCircle is a circle with color
type ColoredCircle struct {
	Circle
	color string
}

// TODO: Implement Colored interface for ColoredCircle
// func (cc *ColoredCircle) GetColor() string { ... }
// func (cc *ColoredCircle) SetColor(color string) { ... }

// TODO: Implement Drawable interface for ColoredCircle
// func (cc ColoredCircle) Draw() string { ... }

func main() {
	fmt.Println("🎮 Challenge 1.3: Interfaces & Polymorphism")
	fmt.Println("==========================================\n")

	// Create different shapes
	circle := Circle{radius: 5.0}
	rectangle := Rectangle{width: 4.0, height: 6.0}
	triangle := Triangle{base: 3.0, height: 4.0}

	// Test polymorphism - same function works with different types!
	fmt.Println("=== Testing Polymorphism ===")
	PrintShapeInfo(circle)
	PrintShapeInfo(rectangle)
	PrintShapeInfo(triangle)

	// Test type assertions
	fmt.Println("\n=== Testing Type Assertions ===")
	fmt.Printf("Circle type: %s\n", GetShapeType(circle))
	fmt.Printf("Rectangle type: %s\n", GetShapeType(rectangle))
	fmt.Printf("Triangle type: %s\n", GetShapeType(triangle))

	// Test ShapeCollection
	fmt.Println("\n=== Testing ShapeCollection ===")
	collection := &ShapeCollection{}
	collection.Add(circle)
	collection.Add(rectangle)
	collection.Add(triangle)

	fmt.Printf("Total area: %.2f\n", collection.TotalArea())
	largest := collection.FindLargest()
	if largest != nil {
		fmt.Printf("Largest shape: %s (area: %.2f)\n", GetShapeType(largest), largest.Area())
	}

	// Bonus: Test ColoredCircle
	fmt.Println("\n=== Bonus: Colored Shapes ===")
	coloredCircle := &ColoredCircle{
		Circle: Circle{radius: 3.0},
		color:  "red",
	}
	coloredCircle.SetColor("blue")
	fmt.Printf("Colored circle color: %s\n", coloredCircle.GetColor())
	fmt.Printf("Colored circle area: %.2f\n", coloredCircle.Area())

	// Test with ColoredShape interface
	var coloredShape ColoredShape = coloredCircle
	PrintShapeInfo(coloredShape)
	fmt.Printf("Color: %s\n", coloredShape.GetColor())

	fmt.Println("\n✅ Challenge complete!")
}
