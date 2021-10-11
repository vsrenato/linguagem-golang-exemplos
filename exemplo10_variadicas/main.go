package main

import "fmt"

var i = 10

func main() {
	fakeLoop()
	fmt.Println(soma(12, 13, 30, 100))

}

func soma(numbers ...int) int {
	var total int

	for _, number := range numbers {
		total += number
	}

	return total

}

func fakeLoop() {
	fmt.Println(i)
	i--
	if i > 0 {
		fakeLoop()
	}

}
