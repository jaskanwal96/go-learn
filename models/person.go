package models

import "fmt"

// Person represents a person with name and age
// Capitalized struct name = exported
type Person struct {
	Name string // Capitalized field = exported
	Age  int    // Capitalized field = exported
	// email string // lowercase = unexported (private to package)
}

// NewPerson creates a new Person instance
// Constructor function (convention: NewXxx)
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// GetInfo returns a formatted string with person info
func (p *Person) GetInfo() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// Birthday increments the person's age
func (p *Person) Birthday() {
	p.Age++
}
