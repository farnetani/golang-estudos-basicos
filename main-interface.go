package main

import "fmt"

type vehicle interface {
	start () string
}

type car struct {
	name string
}

type motorcycle struct {
	name string
}

func (m motorcycle) startCar() string {
	return "The motorcycle " + m.name + "has been started"
}

func (c car) startMotocycle() string {
	return "The car " + c.name + "has been started"
}

func (m motorcycle) start() string {
	return "The motorcycle " + m.name + "has been started"
}

func (c car) start() string {
	return "The car " + c.name + "has been started"
}

// Para ser usado com a interface
func startVehicle(v vehicle) string {
	return v.start()
}

func main() {
	c := car{"Gol"}
	m := motorcycle{("XPTO")}
	//fmt.Println(c.startCar())
	//fmt.Println(m.startMotocycle())

	fmt.Println(startVehicle(c)) // Resulta em erro pq estava usando startCar
	fmt.Println(startVehicle(m)) // Resulta em erro pq estava usando startCar

}