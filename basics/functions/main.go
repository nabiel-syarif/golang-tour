package main

import "fmt"

// basic function without return
func sayHello() {
	fmt.Println("Hello, Tokopedia!!!")
}

// basic function with return
func add(x int, y int) int {
	return x + y
}

// shorten version of function parameters when the data type is same
func subtract(x, y int) int {
	return x - y
}

// named return
func multiply(x, y int) (res int) {
	res = x * y

	return res
}

// multiple return value
func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	sayHello()
	fmt.Println(add(1, 2))
	fmt.Println(subtract(1, 2))
	fmt.Println(multiply(2, 2))

	a, b := "A", "B"
	a, b = swap(a, b)
	fmt.Println(a, b)

}
