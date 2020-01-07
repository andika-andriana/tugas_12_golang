package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Memulai Program...")

	var nama string
	fmt.Print("\nSilahkan Input Nama Lengkap Anda: ")
	fmt.Scanln(&nama)

	var regex, _ = regexp.Compile(`[[:alpha:]]+`)
	var idx = regex.FindStringIndex(nama)
	var str = regex.FindAllString(nama, 1)

	fmt.Println("\nNama Depan Anda Adalah: ", str)
	fmt.Println("\nDengan Index: ", idx)
	fmt.Println("\nSelesai...")
}
