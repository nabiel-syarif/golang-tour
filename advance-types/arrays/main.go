package main

import "fmt"

func main() {
	var array = [10]int{1, 2, 3}
	fmt.Println(array)

	array[0] = 100
	fmt.Println(array)
}
