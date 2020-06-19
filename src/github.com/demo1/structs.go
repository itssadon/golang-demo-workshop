package main

import "fmt"

type Book struct {
	name   string
	author string
	pages  int
}

type Shelf struct {
	position int
	book     Book
}

func structdemo() {
	book := Book{name: "Harry Potter", author: "J.K Rowling", pages: 800}
	fmt.Println(book)
}
