package main

import (
	"fmt"
)

func sayGreeting(nome string) {
	fmt.Println("Olá", nome)
}
func addnumber(numero1 int, numero2 int) int {
	return numero1 + numero2
}
func main() {
	sayGreeting("juvelino")
	resultado := addnumber(180, 360)
	fmt.Println(resultado)

}
