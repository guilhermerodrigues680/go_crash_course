package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 15

	if x > y {
		fmt.Printf("%d é maior que %d\n", x, y)
	} else if x < y {
		fmt.Printf("%d é menor que %d\n", x, y)
	} else {
		fmt.Printf("%d é igual %d\n", x, y)
	}

	color := "red"

	if color == "red" {
		fmt.Println("A cor é vermelha")
	} else if color == "blue" {
		fmt.Println("A cor é azul")
	} else {
		fmt.Println("A cor não é azul e nem vermelha")
	}

	color = "blue"

	// switch
	switch color {
	case "red":
		fmt.Println("A cor é vermelha")
	case "blue":
		fmt.Println("A cor é azul")
	default:
		fmt.Println("A cor não é azul e nem vermelha")
	}

}
