package main

import "fmt"

func main() {
	var number = 100
	var pointerToInt *int = &number
	fmt.Printf("%v\n", pointerToInt)
	fmt.Printf("%v\n", *pointerToInt)
	
	*pointerToInt = 50
	fmt.Printf("%v\n", number)
}
