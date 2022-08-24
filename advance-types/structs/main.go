package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	person := person{
		name: "Nabiel",
		age:  18,
	}
	fmt.Println(person)
	fmt.Println(person.name)
	
	pointerToPerson := &person
	fmt.Println(pointerToPerson.name)
}
