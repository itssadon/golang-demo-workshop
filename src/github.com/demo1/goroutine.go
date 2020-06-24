package main

import "fmt"

func goroutinedemo() {
	count(5)
}

func count(num int) {
	for i := 0; i <= num; i++ {
		fmt.Println(i)
	}
}
