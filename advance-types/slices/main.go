package main

import "fmt"

func main() {

	sliceOfInt := []int{1, 2}
	sliceOfInt = append(sliceOfInt, 1, 2, 3)

	// passed by reference
	newSlice := sliceOfInt[:2]
	fmt.Println(sliceOfInt)
	fmt.Println(newSlice)

	newSlice[0] = 100
	fmt.Println(sliceOfInt)

	var nilSlices []int
	fmt.Println(nilSlices == nil)

	// The length of a slice is the number of elements it contains.
	// The length of a slice is the number of elements it contains.
	sliceWithMake := make([]int, 100, 101)

	fmt.Println(len(sliceWithMake))
	fmt.Println(cap(sliceWithMake))

	twoDimensionalSlice := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(twoDimensionalSlice)
}
