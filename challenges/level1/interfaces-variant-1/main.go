package challenge3_1

// Custom Errors and Printing

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64
type SqrtNumber float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func (x SqrtNumber) String() string {
	return fmt.Sprintf("Sqrt of the number is: %v", float64(x))
}

func Sqrt(x float64) (SqrtNumber, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return SqrtNumber(math.Sqrt(x)), nil
}

func main() {
	if result, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	if result, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
