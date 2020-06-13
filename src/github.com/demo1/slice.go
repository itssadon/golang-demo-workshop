package main

import "fmt"

func slice1() {
	x := make([]int, 0)

	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)

	fmt.Println(x)
}
