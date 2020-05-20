package main

import "fmt"

type SuperNumber int // Siginifica que eu tenho um tipo chamado SuperNumber e conseguirá ser utilizado dentro de qualquer função

type CarName string
type CarYear int
type Color string

type Car struct {
	Name string
	Year int
	Color string
}

// Exemplo de função
func (c Car) info() string {
	return fmt.Sprintf("CAr: %s\n Year: %d\n Color: %s", c.Name, c.Year, c.Color)
}

func main() {
	x := []SuperNumber{}
	x = append(x, 1)
	fmt.Println(x)

	car1 := Car{"Corolla", 2017, "White"}
	car2 := Car{"BMW 320i", 2017, "Black"}
	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car1.Name)
	fmt.Println(car2.Color)

	// Exemplo do uso da função criada com o tipo struct
	fmt.Println(car1.info())

}
