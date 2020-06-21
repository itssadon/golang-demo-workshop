package main

import "fmt"

func adslice() {
	s := make([]int, 3)
	for i := 0; i < 3; i++ {
		s[i] = i + 1
	}
	s = append(s, 4, 5, 6, 7, 8, 9)

	fmt.Println("s =", s)

	// copy slice
	b := make([]int, len(s))
	copy(b, s)
	fmt.Println("b =", b)

	// cut out specific index
	s = append(s[:2], s[4:]...)
	fmt.Println("cut(3,4) s =", s)
}
