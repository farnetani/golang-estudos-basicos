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

}
