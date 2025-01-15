package main

import "fmt"

func main() {

	// Membuat tipe data baru dari tipe data yang sudah ada
	
	type NoKTP string
	var ktpAdi NoKTP = "123456789"

	contoh := "987654321"

	var ktpBaru NoKTP = NoKTP(contoh)

	fmt.Println(ktpAdi, ktpBaru)
}