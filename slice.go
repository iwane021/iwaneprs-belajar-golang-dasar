package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"Sepetember",
		"Oktober",
		"November",
		"Desember",
	}

	//Slice
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0] = "Diubah"
	fmt.Println(slice1)
	fmt.Println(months)

	slice1[0] = "Mei"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)
	fmt.Println("-------------------------")

	//Append
	var slice3 = append(slice2, "Iwan")
	fmt.Println(slice3)

	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println("-------------------------")
	fmt.Println(slice2)
	fmt.Println(months)
	fmt.Println("-------------------------")

	// Make
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Iwan"
	newSlice[1] = "Setiawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println("-------------------------")

	// Copy
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("copy: ", copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("IniArray: ", iniArray)
	fmt.Println("IniSlice: ", iniSlice)
}
