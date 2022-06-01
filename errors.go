// Copy your Sqrt function from the earlier exercise and modify it to return an error value.
// Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.
// Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)).
// Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.



package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
		return fmt.Sprintf(", cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x>=0 {
		z:= 1.0
		for i:=0; i<10; i++ {
			z -= ((z*z - x) / (2*z))
		}
		return z, nil
	}
	return 0, ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
