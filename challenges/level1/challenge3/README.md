# Challenge 1.3: Interfaces & Polymorphism

## Learning Objectives

After completing this challenge, you'll understand:

1. **Interfaces in Go** - How to define and use interfaces
2. **Polymorphism** - One interface, multiple implementations
3. **Type Assertions** - Checking and converting interface types
4. **Interface Composition** - Combining multiple interfaces
5. **Method Sets** - When methods are available on interfaces

## Key Concepts

### What are Interfaces?

In Go, interfaces define **behavior** (what a type can do), not structure. Unlike traditional OOP languages:

- **No explicit "implements" keyword** - If a type has the required methods, it implements the interface automatically
- **Duck typing** - "If it walks like a duck and quacks like a duck, it's a duck"
- **Composition over inheritance** - Prefer interfaces over class hierarchies

### Example:

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    radius float64
}

// Circle automatically implements Shape because it has Area() method
func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}
```

### Polymorphism

You can use interfaces to write functions that work with multiple types:

```go
func PrintArea(s Shape) {
    fmt.Println(s.Area())  // Works with Circle, Rectangle, Triangle, etc.
}
```

### Type Assertions

Check if an interface value is a specific type:

```go
if circle, ok := shape.(Circle); ok {
    // shape is a Circle
    fmt.Println(circle.radius)
}
```

## Hints

1. **Area formulas:**
   - Circle: `π * r²`
   - Rectangle: `width * height`
   - Triangle: `(base * height) / 2`

2. **Perimeter formulas:**
   - Circle: `2 * π * r`
   - Rectangle: `2 * (width + height)`
   - Triangle: `base + height + hypotenuse` (use Pythagorean theorem)

3. **Type assertion syntax:**
   ```go
   value, ok := interfaceValue.(SpecificType)
   ```

4. **Interface composition:**
   ```go
   type Combined interface {
       Interface1
       Interface2
   }
   ```

## Common Mistakes

- Forgetting that interfaces are satisfied implicitly
- Using pointer receivers when value receivers would work
- Not handling the `ok` value in type assertions
- Confusing interface types with concrete types

## Next Steps

After completing this challenge:
- Try adding more shape types (Square, Ellipse, etc.)
- Experiment with interface composition
- Learn about the empty interface `interface{}` (or `any` in Go 1.18+)

