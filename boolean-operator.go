package main

import "fmt"

func main() {

	ujian := 75
	absensi := 88

	lulusUjian := ujian >= 80
	lulusAbsensi := absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)
	fmt.Println("----------------------")
	lulus := lulusAbsensi && lulusUjian
	fmt.Println("Lulus 1: ", lulus)
	fmt.Println("Lulus 2: ", lulusAbsensi && lulusUjian)
	fmt.Println("Lulus 3: ", ujian >= 80 && absensi >= 80)
}
