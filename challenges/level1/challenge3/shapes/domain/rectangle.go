package domain

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}

