package main

import "fmt"

func recursion() {
	z := fact(5)
	fmt.Printf("Factorial : %d\n", z)
}

func fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}
