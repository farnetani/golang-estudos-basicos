package main

import "fmt"

func main() {
	slice := make([]int, 0)
	fmt.Println(slice)

	slice2 := make([]int, 5)
	slice2 = append(slice2, 1) // para adicionar mais 1
	fmt.Println(slice2)
	fmt.Println(len(slice2)) // = 6

	slice3 := make([]int, 2, 5) // array de 5 posições, porém de tamanho 2
	slice3 = append(slice3, 10, 2, 5, 4) // se ultrapassar, ele dobra = 10
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slicex := make([]int, 2, 5)
	for i := 0; i < 20; i++ {
		slicex = append(slicex, i)
		fmt.Println("Len: ", len(slicex), " Cap: ", cap(slicex))
	}

	sliceTest := slicex
	slicex[0] = 10
	slicex = append(slicex, 9)
	fmt.Println(slicex)
	fmt.Println(sliceTest)

	slicex2 := make([]int, 5)
	sliceTest2 := slicex2

	fmt.Println(slicex2)
	fmt.Println(sliceTest2)

	slicex2 = append(slicex2, 1, 2, 3, 5)
	slicex2[0] = 10
	fmt.Println(slicex2)
	fmt.Println(sliceTest2)

	// Criar um slice sem make
	// []   (sem valor = slice)
	// [10] (com valor = igual a array)
	sliceString := []string {  // Cria o slice sem a necessidade do comando make
		"Hello",
		"World",
		"Much",
		"Better",
	}
	fmt.Println(sliceString)
	fmt.Println(sliceString[0]) // Hello
	fmt.Println(sliceString[:2]) // Hello World (pega até a posição 2
	fmt.Println(sliceString[1:2]) // World (pega 2 posições depois da primeira = World
	fmt.Println(sliceString[2:4]) // Much Better
	fmt.Println(sliceString[2:]) // Much Better (pega da posição 2 até o infinito
}
