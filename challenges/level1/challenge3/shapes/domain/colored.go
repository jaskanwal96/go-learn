package domain

type Color struct {
	ColorName string
}

type Colored interface {
	Color() string
}

type ColoredShape interface {
	Shape
	Colored
}

func (c Color) Color() string {
	return c.ColorName
}

