package main

import (
	"fmt"
)

func dividir(dividendo int, divisor int) (int, string) {
	if divisor == 0 {
		return 0, "erro na divisão por zero"
	}
	return dividendo / divisor, "sem erro"
}

func main() {
	resultado, erro := dividir(10, 0)
	if erro != "sem erro" {
		fmt.Println(erro)
	} else {
		fmt.Println("O resultado da divisão é", resultado, erro)
	}
}
