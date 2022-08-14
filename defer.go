package main

import "fmt"

func logging() {
	fmt.Println("Logging Berjalan")
}

func runApplication(angka int) {
	//defer disimpan diawal, agar tetep berjalan sekalipun function terjadi error
	defer logging()
	fmt.Println("Run Application")
	result := 10 / angka
	fmt.Println("Result", result)
}

func main() {
	// defer function merupakan function yg bisa dijadwalkan berjalan setelah suatu function
	// selesai dieksekusi
	// defer tetap akan berjalan sekalipun function tersebut error ketika dieksekusi
	runApplication(0)
}
