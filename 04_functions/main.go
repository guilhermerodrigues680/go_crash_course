package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func sum(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(greeting("Guilherme"))
	fmt.Println(sum(10, 5))
}
