package main

import "fmt"

func main() {
	var nama1 = "iwan"
	var nama2 = "bule"

	var result bool = nama1 == nama2
	fmt.Println(result)
	fmt.Println("-----------------")

	var value1 = 100
	value2 := 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
