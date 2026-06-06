package render

import (
	"fmt"

	"challenge3/shapes/domain"
)

func GetShapeInfo(s domain.Shape) string {
	return fmt.Sprintf("Type: %T, Area: %.2f, Perimeter: %.2f\n", s, s.Area(), s.Perimeter())
}

func Draw(s domain.ColoredShape) {
	fmt.Printf(
		"Drawing %T with perimeter %.2f and color %s\n",
		s,
		s.Perimeter(),
		s.Area(),
		s.Color(),
	)
}
