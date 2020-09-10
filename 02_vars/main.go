package main

import "fmt"

func main() {

	//var nome string = "Guilherme"
	var nome = "Guilherme"
	var idade = 12
	isCool := false

	pi := 3.14

	fmt.Println(nome, idade, isCool, pi)
	fmt.Printf("%v %#v %T %%\n", nome, nome, nome)
	fmt.Printf("%v %#v %T %%\n", idade, idade, idade)
}
