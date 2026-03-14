package main

import (
	"fmt"
)

type Car string
func (c Car) Accelerate() {
	fmt.Println("Speeding up!")
} 
func (c Car) Brake() {
	fmt.Println("Stoping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string
func (t Truck) Accelerate() {
	fmt.Println("Speeding up!")
} 
func (t Truck) Brake() {
	fmt.Println("Stoping")
}
func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}
func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Steer(string)
	Brake()
}

func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()
	truck, ok := vehicle.(Truck)
	if ok {
		truck.LoadCargo("test cargo")
	}
}

func main() {
	var vehicle Vehicle = Car("Toyota Yarvic")
	vehicle.Accelerate()
	vehicle.Steer("Left")

	vehicle = Truck("Ford F180")
	vehicle.Brake()
	vehicle.Steer("right")
	fmt.Println("=========================")

	TryVehicle(Truck("test"))



}
