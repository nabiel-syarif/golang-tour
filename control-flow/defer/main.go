package main

import "fmt"

func main() {
	fmt.Println("Mulai")

	// do something here
	count()

	defer fmt.Println("Selesai")
}

func count() {
	// stacking defers, (LIFO - Last in first out)
	for i := 0; i < 5; i++ {
		defer fmt.Println(i + 1)
	}
}
