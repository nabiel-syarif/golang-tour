package main

import "fmt"

// constants package level
const Pi = 3.14

func main() {
	// constants function level
	const lightSpeed = 3e8

	// cannot reassign constant
	// lightSpeed = 200

	// constant can be character, string, or number
	const name = "Nabiel"
	const grade = 'A'

	// constant can't be declared using := syntax
	// const a := 1

	fmt.Println(Pi, lightSpeed, name, grade)
}
