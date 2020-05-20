package main

import "fmt"

type Car struct {
	Name string
	Year int
	Color string
}

//type SuperCar struct {
//	Name string
//	Year int
//	Color string
//	CanFly bool
//}

type SuperCar struct {
	Car
	CanFly bool
}


// Exemplo de função
func (c Car) info() string {
	return fmt.Sprintf("CAr: %s\n Year: %d\n Color: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"Corolla", 2017, "White"}

	//sCar := SuperCar{"Fusca", 2030, "Blue", true}
	sCar := SuperCar{
		Car: Car{
			"Fusca",
			2030,
			"Blue",
		},
		CanFly: true,
	}
	fmt.Println(sCar)
	fmt.Println(sCar.Name) // Posso usar direto
	fmt.Println(sCar.Color)
	fmt.Println(sCar.Car.Name) // Posso usar a referencia tb

	fmt.Println(car1.info())

	sCar2 := SuperCar{
		Car: car1,
		CanFly: true,
	}
	fmt.Println(sCar2)

}
