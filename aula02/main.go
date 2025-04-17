package main

import (
	"fmt"
)

func main() {
	alunoIdade := make(map[string]int)
	alunoIdade["Maite"] = 26
	alunoIdade["Bruno"] = 38
	alunoIdade["Bernardo"] = 48
	alunoIdade["Giovanna"] = 16
	fmt.Println("A idade de Maite é", alunoIdade["Maite"])
	notasMedic := map[string]float64{
		"Maite":    1000,
		"Bruno":    10000,
		"Bernardo": 0,
		"Giovanna": 8,
	}

	for k, v := range notasMedic {
		fmt.Printf("%s tirou a nota %.1f  \n", k, v)
	}
}
