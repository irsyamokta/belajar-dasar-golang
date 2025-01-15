package main

import "fmt"

func main() {
	
	// Operator Unary digunakan untuk menambahkan atau mengurangi nilai dari suatu variabel

	a := 10
	b := false

	a++ // increment
	a-- // decrement

	b = !b // negasi

	fmt.Println(a)
	fmt.Println(b)
}