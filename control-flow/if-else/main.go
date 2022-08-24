package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	circle1 := math.Pow(math.Pi*float64(rand.Intn(10)), 2)
	circle2 := math.Pow(math.Pi*float64(rand.Intn(10)), 2)

	if circle1 > circle2 {
		fmt.Println("Circle 1")
	} else {
		fmt.Println("Circle 2")
	}

	// if else with statement
	if temp := rand.Intn(2); temp == 0 {
		fmt.Println("Temp value is zero")
	}
}
