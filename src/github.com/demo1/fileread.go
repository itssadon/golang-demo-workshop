package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readfiledemo() {
	// read()
	// read1()
	read2()
}

func read() {
	content, err := ioutil.ReadFile("src/github.com/demo1/dunkirk.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
}

func read1() {
	dat, _ := ioutil.ReadFile("src/github.com/demo1/dunkirk.txt")
	fmt.Println(string(dat))
}

func read2() {
	f, _ := os.Open("src/github.com/demo1/dunkirk.txt")
	b := make([]byte, 100)

	for {
		r, _ := f.Read(b)
		if r == 0 {
			break
		}
		fmt.Println(r)
		fmt.Println(string(b))
	}
}
