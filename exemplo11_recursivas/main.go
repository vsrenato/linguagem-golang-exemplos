package main

import "fmt"

func main() {
	var l List

	l = []interface{}{
		10,
		"Olá",
		1.9,
		false,
	}

	l.Show()

}

type List []interface{}

func (l List) Show() {
	fmt.Println(l...)
}
