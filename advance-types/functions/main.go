package main

import "fmt"

func main() {

	functionAsValue := func() string {
		return "Woww, amazing !!"
	}

	var mainScopedVar = 100
	closure := func() {
		fmt.Println("mainScopedVar =", mainScopedVar)
	}

	fmt.Println(functionAsValue())
	closure()
}
