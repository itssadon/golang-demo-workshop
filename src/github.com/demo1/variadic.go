package main

import "fmt"

func variadic() {
	sum(1, 2, 3)
}

func sum(nums ...int) {
	fmt.Println("Nums Received", nums)
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println("Total :", total)
}
