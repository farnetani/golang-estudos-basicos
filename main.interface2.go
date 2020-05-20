package main

import "fmt"

type Names []interface{}

// Tem que ser um ponteiro para não fazer a cópia
func (n *Names) load() {
	*n = Names {
		"Junior",
		"Lucca",
		"Virginia",
	}
}

func (n Names) printNames() {
	fmt.Println(n)
}

func main () {
	var names Names
	names.load()
	names.printNames()
}
