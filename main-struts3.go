package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name string
	Year int
	Color string
}
func main() {
	car := Car{"Gol", 2017, "Yellow"}

	result, _ := json.Marshal(car)
	fmt.Println("ASCII:", result)

	// Convertendo o Json Ascii para String
	// Para vir os atributos, temos que ter os atributos iniciando em maíusculo, caso queiramos proteger um field/coluna devemos deixar minusculo
	fmt.Println(string(result))

}
