package main

import "fmt"

func main() {

	// Operator boolean

	fmt.Println(true && true) // true
	fmt.Println(true && false) // false
	fmt.Println(true || true) // true
	fmt.Println(true || false) // true

	fmt.Println(!true) // false
	fmt.Println(!false) // true

	// Studi kasus

	nilaiMataKuliahA := 80
	nilaiMataKuliahB := 60

	nilaiLulus := (nilaiMataKuliahA >= 80) && (nilaiMataKuliahB >= 80)

	fmt.Println("Anda lulus", nilaiLulus)
}