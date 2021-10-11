package main

import "fmt"

func main() {
	var adm Admin

	adm.ranking = 1
	adm.user.name = "Avatar"
	adm.user.online = true

	fmt.Println(adm)
}

type User struct {
	name   string
	online bool
}

type Admin struct {
	user    User
	ranking int
}
