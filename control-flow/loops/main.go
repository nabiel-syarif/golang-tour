package main

import "fmt"

func main() {
	basicLoop()
	loopWithBreak()
	infiniteLoop()
}

func basicLoop()  {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func loopWithBreak()  {
	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Println("BREAK LOOP")
			break
		}
	}
}

func infiniteLoop()  {
	i := 1
	for ;; i++{
		if i == 1001 {
			break
		}
		fmt.Println(i)
	}
}