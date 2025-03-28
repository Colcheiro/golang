package main

import "fmt"

func main() {
	a, b := 10, 3
	fmt.Println("a soma é", a+b)
	fmt.Println("a subtração é", a-b)
	fmt.Println("a multiplicação", a*b)

	a++
	fmt.Println("incrementar a",a)

	if a > 0 && b > 0 {
		fmt.Println("numeros positivos")
	}

}
