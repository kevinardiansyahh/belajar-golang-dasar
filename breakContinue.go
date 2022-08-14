package main

import "fmt"

func main() {
	//Break
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke", i)
	}

	//Continue
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("perulangan ke", i)
	}
}
