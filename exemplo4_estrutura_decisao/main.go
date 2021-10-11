package main

import "fmt"

func main() {
	number := 20

	if number > 0 {
		fmt.Println("maior que 0")
	} else {
		fmt.Println("menor que 0")
	}

	switch number {
	case 2:
		fmt.Println("number 2")
	case 5:
		fmt.Println("number 5")
	default:
		fmt.Println("default")
	}
}
