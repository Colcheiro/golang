package main

import (
	"fmt"
)

func sayMyName(name string) {
	fmt.Println("Olá", name)
}

func getSaldo() int {
	var saldo int
	fmt.Println("Digite seu saldo:")
	fmt.Scan(&saldo)
	return saldo
}

func getSaque() int {
	var saque int
	fmt.Println("Digite quanto você deseja sacar:")
	fmt.Scan(&saque)
	return saque
}

func atualizaSaldo(saldo int, saque int) int {
	return saldo - saque
}

func main() {
	sayMyName("Juvelino")
	saldo := getSaldo()
	saque := getSaque()
	resultado := atualizaSaldo(saldo, saque)
	fmt.Println("Saldo restante:", resultado)
}
