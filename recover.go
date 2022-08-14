package main

import "fmt"

func endApp1() {
	message := recover()
	if message != nil {
		fmt.Println("Aplikasi Error Dengan Message :", message)
	} else {
		fmt.Println("Aplikasi selesai")
	}
}

func runApp1(error bool) {
	defer endApp1()
	if error {
		panic("APLIKASI ERROR!")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	// recover digunakan untuk mengambil message error dari panic function
	// & program tidak akan berhenti (recover jg bertugas menghentikan panic)
	runApp1(true)
	// print dibawah tetap akan berjalan karena ada recover
	fmt.Println("Jalan")
}
