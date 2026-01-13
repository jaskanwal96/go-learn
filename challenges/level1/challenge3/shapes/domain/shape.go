package domain

// Shape interface defines the contract for all shapes
// Any type that implements Area() and Perimeter() automatically satisfies this interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

