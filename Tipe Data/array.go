package main

import "fmt"

func main() {

	var number [3]int // Membuat deklarasi array

	number[0] = 1
	number[1] = 2
	number[2] = 3

	name := [...]string{"Irsyam", "Okta", "Pratama", "Riyadi"} // Membuat array secara langsung saat deklarasi

	fmt.Println(number[1])
	fmt.Println(name[0])

	// Function Array

	fmt.Println(len(number)) // Untuk menghitung panjang array
	fmt.Println(len(name))

	number[0] = 100 // Mengubah elemen array
	fmt.Println(number)
}