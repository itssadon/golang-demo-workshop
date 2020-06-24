package main

import "fmt"

func channeldemo() {
	ch := make(chan string)

	func(msg string) {
		ch <- msg
	}("hello")

	recv := <-ch

	fmt.Println(recv)
}
