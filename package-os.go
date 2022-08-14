package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argumen :", args)
	//jika ingin menambahan argumen, tulis command go run package-os.go kevin ardiansyah hidayat
	//fmt.Println(args[1])
	//fmt.Println(args[2])
	//fmt.Println(args[3])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err.Error())
	}
}
