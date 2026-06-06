package domain

type ColoredCircle struct {
	Circle
	ColorValue Color
}

// Color method delegates to embedded ColorValue struct
// This ensures ColoredCircle implements Colored interface
func (cc ColoredCircle) Color() string {
	return cc.ColorValue.Color()
}

