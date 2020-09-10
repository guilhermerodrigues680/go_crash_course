package main

import "fmt"

func main() {
	// For is Go's "while"
	// var i int8 = -128

	// for i < 127 {
	// 	fmt.Print(i, "\t")
	// 	i++
	// }

	// Forever
	// for {
	// }

	// for j := 0; j < 15; j++ {
	// 	fmt.Println(j)
	// }

	// Fizz Buzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 { // ou i % 15
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
