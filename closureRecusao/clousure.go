//closure: funcao "chamar" uma variável que está ewm outra função

package main

import "fmt"

func main() {
	x := 0

	numero := func() int {
		x++
		return x
	}

	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
}