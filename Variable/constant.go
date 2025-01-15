package main

import "fmt"

func main() {
	const firtstName = "Irsyam"
	const lastName = "Okta"
	const age = 20

	// Error
	// firtstName = "Okta"
	// lastName = "Irsyam"
	// age = 30

	fmt.Println(firtstName, lastName, age)
}