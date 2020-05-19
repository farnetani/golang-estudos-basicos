package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	fmt.Println(m)

	delete(m,"c")
	fmt.Println(m["c"])

	_, exists := m["c"]
	fmt.Println(exists)

	_, exists2 := m["b"]
	fmt.Println(exists2)

	value, exists := m["a"]
	println(value)
}
