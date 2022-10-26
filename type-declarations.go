package main

import "fmt"

func main() {
	type NoKtp string
	type MerriedStatus bool

	var noKtpSaya NoKtp = "123456789456"
	var status MerriedStatus = true

	fmt.Println(noKtpSaya)
	fmt.Println(status)
}
