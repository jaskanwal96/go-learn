package domain

type ShapeCollection struct {
	shapes []Shape
}

func (sc *ShapeCollection) Add(s Shape) {
	sc.shapes = append(sc.shapes, s)
}

func (sc ShapeCollection) TotalArea() float64 {
	totalArea := 0.00
	for _, s := range sc.shapes {
		totalArea += s.Area()
	}
	return totalArea
}

func (sc ShapeCollection) Largest() Shape {
	maxArea := 0.00
	var largestShape Shape
	for _, s := range sc.shapes {
		if maxArea < s.Area() {
			maxArea = s.Area()
			largestShape = s
		}
	}
	return largestShape
}

