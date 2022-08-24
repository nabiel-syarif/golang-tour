package main

import (
	"fmt"
	"math/rand"
)

func main() {
	temp := rand.Intn(2)
	switch temp {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("Unknown")
	}

	temp = 1
	switch temp {
	case 1:
		fmt.Println("got 1")
		fallthrough
	case 0:
		fmt.Println("got 0")
	}

	// switch without condition is same as `switch true`
	switch {
	// all cases here should returning boolean
	case 1 > 2:
		fmt.Println("1 greater than 2")
	case 1 < 2:
		fmt.Println("1 less than 2")
	}
}
