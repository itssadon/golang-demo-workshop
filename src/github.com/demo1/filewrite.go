package main

import (
	"fmt"
	"io/ioutil"
)

func writefiledemo() {
	write1()
}

func write1() {
	b := []byte("How did you like the review?\nHope you liked it.\n")
	ioutil.WriteFile("src/github.com/demo1/demo.txt", b, 0644)
	fmt.Println("Done")
}
