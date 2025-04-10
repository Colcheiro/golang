package main

import (
	"fmt"
)

func main() {
	var numeros [5]int
	var soma int

	fmt.Println("Digite cinco numeros inteiros")
	for i := 0; i < 5; i++ {
		fmt.Printf("Numero %d: ", i+1)
		fmt.Scan(&numeros[i])
		soma += numeros[i]
	}

	fmt.Printf("A soma dos números é: %d\n", soma)
}
