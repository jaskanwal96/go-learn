package main

import (
	"challenge3/shapes/domain"
	"challenge3/shapes/render"
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

func main() {
	var sc domain.ShapeCollection
	sc.Add(domain.Circle{Radius: 5})
	sc.Add(domain.ColoredCircle{
		ColorValue: domain.Color{ColorName: "red"},
		Circle:     domain.Circle{Radius: 10},
	})
	sc.Add(domain.Rectangle{Width: 4, Length: 6})
	sc.Add(domain.Triangle{Base: 3, Height: 4})

	fmt.Println(render.GetShapeInfo(sc.Largest()))
	if c, ok := sc.Largest().(domain.Colored); ok {
		fmt.Println("Color:", c.Color())
	}
	fmt.Println(sc.TotalArea())

	render.Draw(domain.ColoredCircle{
		ColorValue: domain.Color{ColorName: "red"},
		Circle:     domain.Circle{Radius: 2},
	})
}
