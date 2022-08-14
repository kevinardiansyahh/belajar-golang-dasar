package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Kevin"
	name[1] = "Ardiansyah"
	name[2] = "Hidayat"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	//membuat array langsung
	var values = [3]int{
		12, 13, 14,
	}
	fmt.Println(values)

	//mendapatkan panjang array
	fmt.Println(len(name))
	fmt.Println(len(values))
	var count [10]int
	fmt.Println(len(count))
}
