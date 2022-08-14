package main

import "fmt"

func random() interface{} {
	return false
}

func main() {
	result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)

	//assertion menggunakan switch
	//untuk menghindari panic jika type data yg di konversi salah
	switch value := result.(type) {
	case string:
		fmt.Println("Value String", value)
	case int:
		fmt.Println("Value Int", value)
	default:
		fmt.Println("Unknown Type")
	}
}
