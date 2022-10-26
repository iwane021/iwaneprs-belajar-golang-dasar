package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Iwan"
	names[1] = "Setiawan"
	names[2] = "Pujianto"

	fmt.Println(names)
	fmt.Println(names[0])

	var values = [3]int{
		80,
		90,
		100,
	}

	fmt.Println(values)
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi))
}
