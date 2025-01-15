package main

import "fmt"

func main() {
	
	// Konversi tipe data integer

	nilai32 := int32(32768)
	nilai64 := int64(nilai32)
	nilai16 := int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// Konversi tipe data string

	name := "Irsyam Okta Pratama Riyadi"
	e := name[0]
	eString := string(e)

	fmt.Println(e)
	fmt.Println(eString)

}