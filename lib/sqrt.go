package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not sqrt negative number %v\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	var err error
	if x < 0 {
		err = ErrNegativeSqrt(x)
		return 0, err
	}
	z := x / 2
	for i := 0; i < 10000 && math.Abs(z*z-x) > 0.00001; i++ {
		z -= (z*z - x) / 2 * z
	}
	return z, err
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println(math.Sqrt(2))
}
