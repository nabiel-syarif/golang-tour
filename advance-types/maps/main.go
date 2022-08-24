package main

import "fmt"

func main() {

	dict := map[string]int{
		"Januari":  31,
		"Februari": 28,
	}

	val, isExists := dict["Januari"]
	if isExists {
		fmt.Println(val)
	}

	_, isExists = dict["blabla"]
	fmt.Println(isExists)

	dict["Maret"] = 30
	delete(dict, "Januari")
	fmt.Println(dict)
}
