package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR!")
	}
	fmt.Println("Aplikasi Berjalan")
}
func main() {
	// panic function merupakan function yang bisa digunakan untuk menghentikan program
	// panic function biasanya dipanggil ketika terjadi error saat program berjalan
	// saat panic function dipanggil, program berhenti, namun defer function tetap berjalan

	runApp(true)
}
