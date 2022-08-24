package main

import "fmt"

// package level variables
var a, b string = "a", "b" 

func main() {
	// function (local) level variables
	// this variable will override outer scope variable
	var a string = "LOCAL"
	fmt.Println(a)
	fmt.Println(b)

	// inferred variable data type
	var name, age = "Nabiel", 18
	// name is string, and age is int
	fmt.Println(name, age)

	// short variable declarations
	helloWorld := "Halo, Dunia!"
	fmt.Println(helloWorld)
}