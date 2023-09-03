package main

import (
	"fmt"
)

// create struct
type Car struct {
	Type   string
	FuelIn float64
}

// create method
func (c *Car) CalculateEstimate() float64 {
	// calculate the distance that can be traveled
	return c.FuelIn * 1.5
}

func main() {

	// declare struct
	car := Car{"Alpard", 10.5}

	fmt.Println("Type mobil : ", car.Type)
	fmt.Println("Bahan Bakar yang tersedia : ", car.FuelIn)
	fmt.Println("Perkiraan jarak yang dapat ditempuh : ", car.CalculateEstimate(), "km")

}
