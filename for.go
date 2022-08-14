package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	//for menggunakan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	//for range
	sliceName := []string{"Kevin", "Ardiansyah", "Hidayat"}

	for i, name := range sliceName {
		fmt.Println("index = ", i, "value = ", name)
	}

	//map
	person := make(map[string]string)
	person["name"] = "Kevin"
	person["title"] = "Developer"

	for key, pp := range person {
		fmt.Println(key, "=", pp)
	}

}
