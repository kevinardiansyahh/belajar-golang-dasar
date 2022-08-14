package main

import "fmt"

func main() {
	name := "Kevin"

	switch name {
	case "Ardi":
		fmt.Println("Hai Ardi")
	case "Kevin":
		fmt.Println("Hai Kevin")
	default:
		fmt.Println("Haii")
	}
}
