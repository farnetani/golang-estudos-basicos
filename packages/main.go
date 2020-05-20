package main

import (
	"fmt"
	car2 "github.com/farnetani/iniciando-com-go-lang-oo/packages/car"
)

func main() {
	car := car2.Car{"Gol", "Black"}
	fmt.Println(car.Start())
}
