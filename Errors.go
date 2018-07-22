package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	z := 1.0

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		for i := 0.0; i < 10.0; i++ {
			z -= ((z*z - x) / (2 * z))

		}

		return z, nil
	}

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
