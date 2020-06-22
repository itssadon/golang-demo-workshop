package main

import "fmt"

func mapdemo() {
	nm := map[string]int{"first": 1, "second": 2}
	fmt.Println("Map nm :", nm)

	m := make(map[string]int)
	m["a"] = 12
	m["b"] = 13
	m["c"] = 14
	fmt.Println("Map m :", m)

	A := m["a"]
	fmt.Println("Value of m['a'] is ", A)

	delete(m, "b")
	fmt.Println("Deleted m['b'] from Map m :", m)
}
