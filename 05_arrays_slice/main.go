package main

import "fmt"

func main() {
	// var fruitsArr [2]string
	// fruitsArr[0] = "Apple"
	// fruitsArr[1] = "Orange"

	// fruitsArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitsArr)
	// fmt.Println(fruitsArr[1])

	fruitsSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitsSlice)
	fmt.Println(fruitsSlice[1])
	fmt.Println(len(fruitsSlice))
	fmt.Println(fruitsSlice[1:3])
}
