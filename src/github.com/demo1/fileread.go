package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readfiledemo() {
	read1()
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
