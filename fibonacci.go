// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).



package main

import "fmt"

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	f0, f1 := 0, 1
	return func () int {
		ans := f0 + f1
		f0, f1 = f1, ans
		return ans
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
