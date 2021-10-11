package main

import "fmt"

func main() {
	resultado, msg := soma(10, 20)
	if msg != "" {
		fmt.Println(msg)
	}

	fmt.Println("O resultado da soma é:", resultado)
}

func soma(n1, n2 int) (int, string) {
	fmt.Println("Iniciando soma")
	resultado := n1 + n2
	if resultado > 10 {
		return resultado, "esse valor é maior que 10"
	}
	return resultado, ""
}
