// Implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.



package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z:= 1.0
	for i:=0; i<10; i++ {
		z -= ((z*z - x) / (2*z))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}