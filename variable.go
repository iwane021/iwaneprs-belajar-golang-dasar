package main

import "fmt"

func main() {
	var name string

	name = "Olivier Giroud"
	fmt.Println(name)

	name = "Paolo Maldini"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	age := 30
	fmt.Println(age)

	var (
		firstName = "Iwan"
		lastName  = "Bule"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
