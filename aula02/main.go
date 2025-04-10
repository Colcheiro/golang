package main

import (
	"fmt"
)

func main() {
	age := 45
	fmt.Println(age <=50)
	fmt.Println(age >=50)
	fmt.Println(age ==50)
	fmt.Println(age !=50)

	if age < 30 {
		fmt.Println("Menor que 30 anos")
	} else {
		fmt.Println("Menor que 40 anos")
	} else {
		fmt.Println("Não é menor que 40 anos")
	}
	names := []string{"malu","aidan","francesco"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continue após a posição",index,"e valor", value)
			continue
		}
		if index > 2 {
			fmt.Println("sair após", index)
			break
		}
		fmt.Println("Valor",value)
	}
}
