package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func writefiledemo() {
	// write1()
	write2()
}

func write1() {
	b := []byte("How did you like the review?\nHope you liked it.\n")
	ioutil.WriteFile("src/github.com/demo1/demo.txt", b, 0644)
	fmt.Println("Done")
}

func write2() {
	f, _ := os.OpenFile("src/github.com/demo1/demo.txt", os.O_APPEND|os.O_WRONLY, 0777)
	b := []byte("Adding new line.")
	defer f.Close()

	f.Write(b)
	fmt.Println(f.Stat())
}
