package main

import "fmt"

type person struct {
	name string
}

func (p person) SayHello() {
	fmt.Printf("Hello, %s\n", p.name)
}

func (p *person) SetName(newName string) {
	p.name = newName
}

func (person person) SetName2(newName string) {
	person.name = newName
}

func main() {
	person := &person{
		name: "Nabiel",
	}
	person.SayHello()
	person.SetName("Nabiel Omar")
	person.SayHello()

	person.SetName2("Coba")
	person.SayHello()
}
