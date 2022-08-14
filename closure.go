package main

import "fmt"

func main() {
	//jika ingin membuat function didalam function,
	//perhatikan variable yg digunakan dalam scope function tersebut
	//function yg ada di scope luarnya dapat diakses di didalam sub function
	//sedangkan variable dalam sub function tidak dapat diakses di scope luar

	//variable yg ada discope luar dibawah ini
	counter := 0
	name := "Kevin"

	increment := func() {
		fmt.Println("Increment ke", counter)
		//dapat diakses & dapat diubah di sub function dalamnya
		//dan isi variablenya pun dapat berubah
		name = "Ardi"
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
