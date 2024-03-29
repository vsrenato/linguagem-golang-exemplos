package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4}
	user := map[string]string{
		"name": "Renato",
		"nick": "avatar",
	}
	languages := []string{
		"Go",
		"Python",
		"Kubernetes",
		"Yaml",
	}

	for key, value := range arr {
		fmt.Println(key, value*value)
	}

	for key, value := range user {
		fmt.Printf("O campo \"%s\" tem o valor igual a \"%s\"\n", key, value)
	}

	for _, value := range languages {
		fmt.Println(value)
	}
}
