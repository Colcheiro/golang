package main

import (
	"fmt"
)

func main() {
	estoque := make(map[string]int)

	estoque["Coxinha"] = 10
	estoque["Pao de queijo"] = 15
	estoque["Refrigerante"] = 20

	estoque["Coxinha"] -= 2
	estoque["Pao de queijo"] -= 1

	for item, total := range estoque {
		fmt.Printf("%s sobrou %d unidades\n", item, total)
	}
}
