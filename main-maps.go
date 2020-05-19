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

	// Num MAP eu tenho 2 propriedades: value, exists (valor/se ele existe)
	_, exists := m["c"] // como não existe valor por ele ter sido deletado, eu uso um blank identifier _
	fmt.Println(exists)

	value, exists := m["a"]
	fmt.Println(exists)

	_, exists2 := m["b"]
	fmt.Println(exists2)


	value, exists3 := m["a"]
	fmt.Println(value)
	println(exists3)

	value_c, exists := m["c"]
	fmt.Println(value_c)

	var x = map[string]int{}
	fmt.Println(x)

	newMap := map[string]int{"x":5, "y":10}
	fmt.Println(newMap)

	// se o cara existir, entra na primeira condição
	if value, exists := m["a"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("ops")
	}

}
