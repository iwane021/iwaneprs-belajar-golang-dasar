package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Iwan",
		"address": "Tangerang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}
