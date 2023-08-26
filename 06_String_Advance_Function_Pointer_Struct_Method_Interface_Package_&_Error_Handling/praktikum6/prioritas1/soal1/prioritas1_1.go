package main

import (
	"fmt"
)

type Car struct {
	carType string
	fuelin  float64
}

func (c Car) EstimateDistance() float64 {
	fuelEfficiency := 1.5 // L/mill
	estimatedDistance := c.fuelin * fuelEfficiency
	return estimatedDistance
}

func main() {
	car := Car{
		carType: "SUV",
		fuelin:  10.5,
	}

	estimatedDistance := car.EstimateDistance()
	fmt.Printf("Car type: %s, est. distance: %.2f\n", car.carType, estimatedDistance)
}
